package src

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"uptime/db"

	"text/tabwriter"

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

func GenerateAggregateQuery(startBlock int64, endBlock int64) []bson.M {

	// Read El Choco upgrade configs
	elChocoStartBlock = viper.Get("el_choco_startblock").(int64)
	elChocoEndBlock = viper.Get("el_choco_endblock").(int64)

	// Read Amazonas upgrade configs
	amazonasStartBlock = viper.Get("amazonas_startblock").(int64)
	amazonasEndBlock = viper.Get("amazonas_endblock").(int64)

	aggQuery := []bson.M{}

	//Query for filtering blocks in between given start block and end block
	matchQuery := bson.M{
		"$match": bson.M{
			"$and": []bson.M{
				bson.M{
					"height": bson.M{"$gte": startBlock},
				},
				bson.M{
					"height": bson.M{"$lte": endBlock},
				},
			},
		},
	}

	aggQuery = append(aggQuery, matchQuery)

	//Query for Unwind the Array of validators from each block
	unwindQuery := bson.M{
		"$unwind": "$validators",
	}

	aggQuery = append(aggQuery, unwindQuery)

	//Query for calculating uptime count, upgrade1 count and upgrade2 count
	groupQuery := bson.M{
		"$group": bson.M{
			"_id":          "$validators",
			"uptime_count": bson.M{"$sum": 1},
			"upgrade1_block": bson.M{
				"$min": bson.M{
					"$cond": []interface{}{
						bson.M{
							"$and": []bson.M{
								bson.M{"$gte": []interface{}{"$height", elChocoStartBlock}},
								bson.M{"$lte": []interface{}{"$height", elChocoEndBlock}},
							},
						},
						"$height",
						"null",
					},
				},
			},
			"upgrade2_block": bson.M{
				"$min": bson.M{
					"$cond": []interface{}{
						bson.M{
							"$and": []bson.M{
								bson.M{"$gte": []interface{}{"$height", amazonasStartBlock}},
								bson.M{"$lte": []interface{}{"$height", amazonasEndBlock}},
							},
						},
						"$height",
						"null",
					},
				},
			},
		},
	}

	aggQuery = append(aggQuery, groupQuery)

	//Query for getting moniker, operator address from validators
	lookUpQuery := bson.M{
		"$lookup": bson.M{
			"from": "validators",
			"let":  bson.M{"id": "$_id"},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{"$eq": []string{"$address", "$$id"}},
					},
				},
				bson.M{
					"$project": bson.M{
						"description.moniker": 1, "operator_address": 1, "address": 1, "_id": 0,
					},
				},
			},
			"as": "validator_details",
		},
	}

	aggQuery = append(aggQuery, lookUpQuery)

	return aggQuery
}

// CalculateUpgradePoints - Calculates upgrade score by using upgrade score per block,
// upgrade block and end block height
func CalculateUpgradePoints(upgradeScorePerBlock int64, upgradeBlock int64, endBlockHeight int64) int64 {
	if upgradeBlock == 0 {
		return 0
	}
	score := upgradeScorePerBlock * (endBlockHeight - upgradeBlock + 1)

	return score
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

	aggQuery := GenerateAggregateQuery(startBlock, endBlock)

	results, err := h.db.FetchAllBlocksByAgg(aggQuery)

	if err != nil {
		fmt.Printf("Error while fetching blocks by aggregation %v", err)
		db.HandleError(err)
	}

	for _, obj := range results {
		valInfo := ValidatorInfo{
			ValAddress: obj.Validator_details[0].Address,
			Info: Info{
				OperatorAddr:  obj.Validator_details[0].Operator_address,
				Moniker:       obj.Validator_details[0].Description.Moniker,
				UptimeCount:   obj.Uptime_count,
				Upgrade1Score: CalculateUpgradePoints(elChocoScorePerBlock, obj.Upgrade1_block, elChocoEndBlock),
				Upgrade2Score: CalculateUpgradePoints(amazonasScorePerBlock, obj.Upgrade2_block, amazonasEndBlock),
			},
		}

		validatorsList = append(validatorsList, valInfo)
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
		"\t Upgrade1 score \t Upgrade2 score \t Uptime score")

	for _, data := range validatorsList {
		var address string = data.Info.OperatorAddr

		//Assigning validator address if operator address is not found
		if address == "" {
			address = data.ValAddress + " (Hex Address)"
		}

		fmt.Fprintln(w, " "+address+"\t "+data.Info.Moniker+
			"\t  "+strconv.Itoa(int(data.Info.UptimeCount))+"\t "+strconv.Itoa(int(data.Info.Upgrade1Score))+
			" \t"+strconv.Itoa(int(data.Info.Upgrade2Score))+" \t"+fmt.Sprintf("%f", data.Info.UptimeScore))
	}

	w.Flush()

	//Export data to csv file
	ExportToCsv(validatorsList)
}

// ExportToCsv - Export data to CSV file
func ExportToCsv(data []ValidatorInfo) {
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
		var address string = record.Info.OperatorAddr

		//Assigning validator address if operator address is not found
		if address == "" {
			address = record.ValAddress + " (Hex Address)"
		}

		uptimeCount := strconv.Itoa(int(record.Info.UptimeCount))
		up1Score := strconv.Itoa(int(record.Info.Upgrade1Score))
		up2Score := strconv.Itoa(int(record.Info.Upgrade2Score))
		uptimeScore := fmt.Sprintf("%f", record.Info.UptimeScore)
		addrObj := []string{address, record.Info.Moniker, uptimeCount, up1Score, up2Score, uptimeScore}
		err := writer.Write(addrObj)

		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
