package src

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"text/tabwriter"

	"github.com/regen-network/testnets/kontraua/uptime/db"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

var (
	himalayaStartBlock     int64
	himalayaEndBlock       int64
	himalayaPointsPerBlock int64

	twilightStartBlock     int64
	twilightEndBlock       int64
	twilightPointsPerBlock int64

	nodeRewards int64
)

type handler struct {
	db db.DB
}

func New(db db.DB) handler {
	return handler{db}
}

func CalculateProposal1VoteScore(address string) int64 {
	proposal1Voters := viper.Get("himalaya_vote_validators").([]interface{})

	for _, obj := range proposal1Voters {
		if obj.(string) == address {
			return 100
		}
	}
	return 0
}

func CalculateProposal2VoteScore(address string) int64 {
	proposal2Voters := viper.Get("twilight_vote_validators").([]interface{})

	for _, obj := range proposal2Voters {
		if obj.(string) == address {
			return 100
		}
	}
	return 0
}

func GenerateAggregateQuery(startBlock int64, endBlock int64,
	himalayaStartBlock int64, himalayaEndBlock int64,
	twilightStartBlock int64, twilightEndBlock int64,
	phase6StartBlock int64, phase6EndBlock int64) []bson.M {

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
								bson.M{"$gte": []interface{}{"$height", himalayaStartBlock}},
								bson.M{"$lte": []interface{}{"$height", himalayaEndBlock}},
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
								bson.M{"$gte": []interface{}{"$height", twilightStartBlock}},
								bson.M{"$lte": []interface{}{"$height", twilightEndBlock}},
							},
						},
						"$height",
						"null",
					},
				},
			},
			"phase_6": bson.M{
				"$sum": bson.M{
					"$cond": []interface{}{
						bson.M{
							"$and": []bson.M{
								bson.M{"$gte": []interface{}{"$height", phase6StartBlock}},
								bson.M{"$lte": []interface{}{"$height", phase6EndBlock}},
							},
						},
						1,
						0,
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

// CalculateUpgradePoints - Calculates upgrade points by using upgrade points per block,
// upgrade block and end block height
func CalculateUpgradePoints(upgradePointsPerBlock int64, upgradeBlock int64, endBlockHeight int64) int64 {
	if upgradeBlock == 0 {
		return 0
	}
	points := upgradePointsPerBlock * (endBlockHeight - upgradeBlock + 1)

	return points
}

func GetCommonValidators(gentxVals, blockVals []string) (results []string) {
	m := make(map[string]bool)

	for _, item := range gentxVals {
		m[item] = true
	}

	for _, item := range blockVals {
		if _, ok := m[item]; ok {
			results = append(results, item)
		}
	}

	return results
}

func (h handler) CalculateGenesisPoints(address string) int64 {
	var aggQuery []bson.M

	matchQuery := bson.M{
		"$match": bson.M{
			"height": 2,
		},
	}

	aggQuery = append(aggQuery, matchQuery)

	lookUpQuery := bson.M{
		"$lookup": bson.M{
			"from": "validators",
			"let":  bson.M{"id": "$validators"},
			"pipeline": []bson.M{
				bson.M{
					"$match": bson.M{
						"$expr": bson.M{"$in": []string{"$address", "$$id"}},
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

	results, err := h.db.QueryValAggregateData(aggQuery)

	if err != nil {
		fmt.Printf("Error while fetching validator data at height 2 %v", err)
		db.HandleError(err)
	}

	var blockValidators []string

	if len(results) > 0 {
		for _, val := range results[0].Validator_details {
			blockValidators = append(blockValidators, val.Operator_address)
		}
	}

	// GentxValidators := viper.Get("gentx_validators").([]interface{})
	// var genxVal []string

	// for _, val := range GentxValidators {
	// 	genxVal = append(genxVal, val.(string))
	// }

	// commonValidators := GetCommonValidators(genxVal, blockValidators)

	// for _, val := range commonValidators {
	// 	if val == address {
	// 		return 100
	// 	}
	// }

	return 0

}

func (h handler) CalculateUptime(startBlock int64, endBlock int64) {
	//Read node rewards from config
	// nodeRewards = viper.Get("node_rewards").(int64)

	// Read El Choco upgrade configs
	himalayaStartBlock = viper.Get("el_choco_startblock").(int64) + 1 //Need to consider votes from next block after upgrade
	himalayaEndBlock = viper.Get("el_choco_endblock").(int64) + 1
	himalayaPointsPerBlock = viper.Get("el_choco_reward_points_per_block").(int64)

	// Read twilight upgrade configs
	twilightStartBlock = viper.Get("twilight_startblock").(int64) + 1 //Need to consider votes from next block after upgrade
	twilightEndBlock = viper.Get("twilight_endblock").(int64) + 1
	twilightPointsPerBlock = viper.Get("twilight_reward_points_per_block").(int64)

	// Read phase 6 config
	phase6StartBlock := viper.Get("phase6_startblock").(int64)
	phase6EndBlock := viper.Get("phase6_endblock").(int64) + 1

	var validatorsList []ValidatorInfo //Intializing validators uptime

	fmt.Println("Fetching blocks from:", startBlock, ", to:", endBlock)

	aggQuery := GenerateAggregateQuery(startBlock, endBlock, himalayaStartBlock,
		himalayaEndBlock, twilightStartBlock, twilightEndBlock, phase6StartBlock, phase6EndBlock)

	results, err := h.db.QueryValAggregateData(aggQuery)

	if err != nil {
		fmt.Printf("Error while fetching validator data %v", err)
		db.HandleError(err)
	}

	for _, obj := range results {
		valInfo := ValidatorInfo{
			ValAddress: obj.Validator_details[0].Address,
			Info: Info{
				OperatorAddr:   obj.Validator_details[0].Operator_address,
				Moniker:        obj.Validator_details[0].Description.Moniker,
				UptimeCount:    obj.Uptime_count,
				Upgrade1Points: CalculateUpgradePoints(himalayaPointsPerBlock, obj.Upgrade1_block, himalayaEndBlock),
				Upgrade2Points: CalculateUpgradePoints(twilightPointsPerBlock, obj.Upgrade2_block, twilightEndBlock),
				Phase6Points:   obj.Phase6_points,
			},
		}

		validatorsList = append(validatorsList, valInfo)
	}

	//calculating uptime points
	// for i, v := range validatorsList {
	// maxUptimeRewards := viper.Get("max_uptime_rewards").(int64)
	// uptimePoints := float64(v.Info.UptimeCount*maxUptimeRewards) / (float64(endBlock) - float64(startBlock))

	// validatorsList[i].Info.UptimePoints = uptimePoints

	//calculate proposal1 vote score
	// proposal1VoteScore := CalculateProposal1VoteScore(validatorsList[i].Info.OperatorAddr)

	//calculate proposal2 vote score
	// proposal2VoteScore := CalculateProposal2VoteScore(validatorsList[i].Info.OperatorAddr)

	// validatorsList[i].Info.Proposal1VoteScore = proposal1VoteScore
	// validatorsList[i].Info.Proposal2VoteScore = proposal2VoteScore

	// genesisPoints := h.CalculateGenesisPoints(validatorsList[i].Info.OperatorAddr)
	// validatorsList[i].Info.GenesisPoints = genesisPoints

	// validatorsList[i].Info.TotalPoints = float64(validatorsList[i].Info.Upgrade1Points) +
	// 	float64(validatorsList[i].Info.Upgrade2Points) + uptimePoints + float64(nodeRewards) +
	// 	float64(proposal1VoteScore) + float64(proposal2VoteScore) + float64(genesisPoints)

	// }

	//Printing Uptime results in tabular view
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, " Operator Addr \t Moniker\t Phase6 Count "+
		"\t Upgrade-1 Points \t Upgrade-2 Points")

	for _, data := range validatorsList {
		var address string = data.Info.OperatorAddr

		//Assigning validator address if operator address is not found
		if address == "" {
			address = data.ValAddress + " (Hex Address)"
		}

		fmt.Fprintln(w, " "+address+"\t "+data.Info.Moniker+
			"\t  "+strconv.Itoa(int(data.Info.Phase6Points))+" \t"+strconv.Itoa(int(data.Info.Upgrade1Points))+" \t"+strconv.Itoa(int(data.Info.Upgrade2Points)))
	}

	w.Flush()

	//Export data to csv file
	//ExportToCsv(validatorsList, nodeRewards)
}

// ExportToCsv - Export data to CSV file
func ExportToCsv(data []ValidatorInfo, nodeRewards int64) {
	Header := []string{
		"ValOper Address", "Moniker", "Uptime Count", "Upgrade1 Points",
		"Upgrade2 Points", "Uptime Points", "Node points",
		"Proposal1 Vote Points", "Proposal2 Vote Points", "Genesis Points", "Total Points",
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
		uptimePoints := fmt.Sprintf("%f", record.Info.UptimePoints)
		up1Points := strconv.Itoa(int(record.Info.Upgrade1Points))
		up2Points := strconv.Itoa(int(record.Info.Upgrade2Points))
		nodePoints := strconv.Itoa(int(nodeRewards))
		totalPoints := fmt.Sprintf("%f", record.Info.TotalPoints)
		p1VoteScore := strconv.Itoa(int(record.Info.Proposal1VoteScore))
		p2VoteScore := strconv.Itoa(int(record.Info.Proposal2VoteScore))
		genPoints := strconv.Itoa(int(record.Info.GenesisPoints))
		addrObj := []string{address, record.Info.Moniker, uptimeCount, up1Points,
			up2Points, uptimePoints, nodePoints, p1VoteScore, p2VoteScore, genPoints, totalPoints}
		err := writer.Write(addrObj)

		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
