package src

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"uptime/db"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

var (
	elChocoStartBlock    int64
	elChocoEndBlock      int64
	elChocoScorePerBlock int64

	amazonasStartBlock    int64
	amazonasEndBlock      int64
	amazonasScorePerBlock int64
)

type handler struct {
	db db.DB
}

func New(db db.DB) handler {
	return handler{db}
}

func (h handler) CalculateUptime(startBlock int64, endBlock int64) {
	// Read El Choco upgrade configs
	elChocoStartBlock = viper.Get("el_choco_startblock").(int64)
	elChocoEndBlock = viper.Get("el_choco_endblock").(int64)
	elChocoScorePerBlock = viper.Get("el_choco_reward_score_per_block").(int64)

	// Read Amazonas upgrade configs
	amazonasStartBlock = viper.Get("amazonas_startblock").(int64)
	amazonasEndBlock = viper.Get("amazonas_endblock").(int64)
	amazonasScorePerBlock = viper.Get("amazonas_reward_score_per_block").(int64)

	var validatorsList []ValidatorInfo //Intializing validators uptime

	fmt.Println("Fetching blocks from:", startBlock, ", to:", endBlock)

	//Read all blocks
	blocks, err := h.db.FetchBlocks(startBlock, endBlock)

	if err != nil {
		db.HandleError(err)
	}

	blocksLen := len(blocks)

	for i := 0; i < blocksLen; i++ {
		currentBlockHeight := blocks[i].Height
		var genesisScore int64 = 0

		//Assigning genesis score to validators at block height 2
		if currentBlockHeight == 2 {
			genesisScore = 100
		}

		for _, valAddr := range blocks[i].Validators {
			//Get the validator index from validatorsList
			index := GetValidatorIndex(valAddr, validatorsList)

			if index > 0 {
				// If validator is present in the list already (i.e., joined the network in previous block heights)
				// Update uptime details
				validatorsList[index].Info.UptimeCount++

				//Block height must be in between El Choco upgrade startblock height and endblock height
				if currentBlockHeight >= elChocoStartBlock && currentBlockHeight <= elChocoEndBlock {
					if validatorsList[index].Info.Upgrade1Score == 0 {
						validatorsList[index].Info.Upgrade1Score = elChocoScorePerBlock * (elChocoEndBlock - currentBlockHeight + 1)
					}
				}

				//Block height must be in between Amazonas upgrade startblock height and endblock height
				if (currentBlockHeight >= amazonasStartBlock) && currentBlockHeight <= amazonasEndBlock {
					if validatorsList[index].Info.Upgrade2Score == 0 {
						validatorsList[index].Info.Upgrade2Score = amazonasScorePerBlock * (amazonasEndBlock - currentBlockHeight + 1)
					}
				}
			} else {
				// If the validator is not present in the list i.e., newly joined in the current block
				// Fetch Validator information and Push to validators list
				// Initialize the validator uptime info with default info (i.e., 1)

				query := bson.M{
					"address": valAddr,
				}

				//Get validator by using validator address
				validator, err := h.db.GetValidator(query)

				if err != nil {
					db.HandleError(err)
				}

				valAddressInfo := ValidatorInfo{
					ValAddress: valAddr,
					Info: Info{
						UptimeCount:  1,
						Moniker:      validator.Description.Moniker,
						OperatorAddr: validator.OperatorAddress,
						StartBlock:   int64(currentBlockHeight),
						GenesisScore: genesisScore,
					},
				}

				//Block height must be in between El Choco upgrade startblock height and endblock height
				if currentBlockHeight >= elChocoStartBlock && currentBlockHeight <= elChocoEndBlock {
					if valAddressInfo.Info.Upgrade1Score == 0 {
						valAddressInfo.Info.Upgrade1Score = elChocoScorePerBlock * (elChocoEndBlock - currentBlockHeight + 1)
					}
				}

				//Block height must be in between Amazonas upgrade startblock and endblock
				if (currentBlockHeight >= amazonasStartBlock) && currentBlockHeight <= amazonasEndBlock {
					if valAddressInfo.Info.Upgrade2Score == 0 {
						valAddressInfo.Info.Upgrade2Score = amazonasScorePerBlock * (amazonasEndBlock - currentBlockHeight + 1)
					}
				}

				//Inserting new validator into uptime count
				validatorsList = append(validatorsList, valAddressInfo)
			}
		}
	}

	//calculating uptime score
	for i, v := range validatorsList {
		uptime := float64(v.Info.UptimeCount) / (float64(endBlock) - float64(startBlock))
		uptimeScore := uptime * 300
		validatorsList[i].Info.UptimeScore = uptimeScore

		//Assigning every validator a node score 100
		validatorsList[i].Info.NodeScore = 100
	}

	//Printing Uptime results in tabular view
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, " Operator Addr \t Moniker\t Uptime Count "+
		"\t Upgrade1 score \t Upgrade2 score \t Uptime score \t Node Score \t Genesis Score")

	for _, data := range validatorsList {
		fmt.Fprintln(w, " "+data.Info.OperatorAddr+"\t "+data.Info.Moniker+
			"\t  "+strconv.Itoa(int(data.Info.UptimeCount))+"\t "+strconv.Itoa(int(data.Info.Upgrade1Score))+
			" \t"+strconv.Itoa(int(data.Info.Upgrade2Score))+" \t"+fmt.Sprintf("%f", data.Info.UptimeScore)+
			"\t"+strconv.Itoa(int(data.Info.NodeScore))+"\t"+strconv.Itoa(int(data.Info.GenesisScore)))
	}

	w.Flush()

	//Exporing into csv file
	ExportIntoCsv(validatorsList)
}

// GetValidatorIndex returns the index of the validator from the list
func GetValidatorIndex(validatorAddr string, validatorsList []ValidatorInfo) int {
	var pos int

	for index, addr := range validatorsList {
		if addr.ValAddress == validatorAddr {
			pos = index
		}
	}

	return pos
}

func ExportIntoCsv(data []ValidatorInfo) {
	Header := []string{
		"ValOper Address", "Moniker", "Uptime Count", "elChoco Score", "Upgrade2 Score", "Uptime Score",
	}

	file, err := os.Create("result.csv")

	if err != nil {
		log.Fatal("Cannot write to file", err)
	}

	defer file.Close() //Close file

	writer := csv.NewWriter(file)

	defer writer.Flush()

	//Write header titles
	_ = writer.Write(Header)

	for _, record := range data {
		uptimeCount := strconv.Itoa(int(record.Info.UptimeCount))
		up1Score := strconv.Itoa(int(record.Info.Upgrade1Score))
		up2Score := strconv.Itoa(int(record.Info.Upgrade2Score))
		uptimeScore := fmt.Sprintf("%f", record.Info.UptimeScore)
		addrObj := []string{record.Info.OperatorAddr, record.Info.Moniker, uptimeCount, up1Score, up2Score, uptimeScore}
		err := writer.Write(addrObj)

		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
