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
	upgrade1Height int64
	upgrade2Height int64
)

type handler struct {
	db db.DB
}

func New(db db.DB) handler {
	return handler{db}
}

func (h handler) CalculateUptime(startBlock int64, endBlock int64) {

	upgrade1Height = viper.Get("upgrade1height").(int64)
	upgrade2Height = viper.Get("upgrade2height").(int64)

	var validatorsList []ValidatorInfo //Intializing validators uptime

	fmt.Println("Fetching blocks from:", startBlock, ", to:", endBlock)

	//Read all blocks
	blocks, err := h.db.FetchBlocks(startBlock, endBlock+1)

	if err != nil {
		fmt.Printf("Error while fetching all blocks %v", err)
		os.Exit(1)
	}

	for i := startBlock; i < endBlock; i++ {

		for _, valAddr := range blocks[i].Validators {
			//Get the validator index from validatorsList
			index := GetValidatorIndex(valAddr, validatorsList)

			if index > 0 {
				// If validator is present in the list already (i.e., joined the network in previous block heights)
				// Update uptime details
				validatorsList[index].Info.UptimeCount++

				//Updating upgrade1 score if score is not present and
				//validator already exists in validator list
				if validatorsList[index].Info.Upgrade1Score == 0 {
					validatorsList[index].Info.Upgrade1Score = GetUpgradeScore(i, upgrade1Height)
				}

				//Updating upgrade2 score if score is not present and
				//validator already exists in validator list
				if validatorsList[index].Info.Upgrade2Score == 0 {
					validatorsList[index].Info.Upgrade2Score = GetUpgradeScore(i, upgrade2Height)
				}
			} else {
				// If the validator is not present in the list i.e., newly joined in the current block
				// Fetch Validator information and Push to validators list
				// Initialize the validator uptime info with default info (i.e., 1)

				query := bson.M{
					"address": valAddr,
				}

				//Get validator by using validator address
				validator, _ := h.db.GetValidator(query)

				valAddressInfo := ValidatorInfo{
					ValAddress: valAddr,
					Info: Info{
						UptimeCount:  1,
						Moniker:      validator.Description.Moniker,
						OperatorAddr: validator.OperatorAddress,
						StartBlock:   int64(i),
					},
				}

				if valAddressInfo.Info.Upgrade1Score == 0 {
					valAddressInfo.Info.Upgrade1Score = GetUpgradeScore(i, upgrade1Height)
				}

				if valAddressInfo.Info.Upgrade2Score == 0 {
					valAddressInfo.Info.Upgrade2Score = GetUpgradeScore(i, upgrade2Height)
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
	}

	//Printing Uptime results in tabular view
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, " Address\t Moniker\t Uptime Count \t Upgrade1 score \t Upgrade2 score \t Uptime score")

	for _, data := range validatorsList {
		fmt.Fprintln(w, " "+data.Info.OperatorAddr+"\t "+data.Info.Moniker+
			"\t  "+strconv.Itoa(int(data.Info.UptimeCount))+"\t "+strconv.Itoa(int(data.Info.Upgrade1Score))+
			" \t"+strconv.Itoa(int(data.Info.Upgrade2Score))+" \t"+fmt.Sprintf("%f", data.Info.UptimeScore))
	}

	w.Flush()

	//Exporing into csv file
	ExportIntoCsv(validatorsList)
}

func GetUpgradeScore(blockHeight int64, upgradeHeight int64) int64 {
	var result int64
	value := upgradeHeight + 200

	//Block height must be in between upgrade1 height and sum of 200 and upgrade1 height
	if (blockHeight >= upgradeHeight) && blockHeight < value {
		result = 200 - (blockHeight - upgradeHeight)
	}

	return result
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
		"Address", "Moniker", "Uptime Count", "Upgrade1 Score", "Upgrade2 Score", "Uptime Score",
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
		addrObj := []string{record.ValAddress, record.Info.Moniker, uptimeCount, up1Score, up2Score, uptimeScore}
		err := writer.Write(addrObj)

		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
