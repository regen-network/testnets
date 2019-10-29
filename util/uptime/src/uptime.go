package src

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"uptime/db"

	"gopkg.in/mgo.v2/bson"
)

type handler struct {
	db db.DB
}

func New(db db.DB) handler {
	return handler{db}
}

func (h handler) CalculateUptime() {
	var finalValAddrs []ValidatorInfo //Intializing validators uptime

	//Read all blocks
	blocks, _ := h.db.ReadAllBlocks()

	for currentHeight := 0; currentHeight < len(blocks); currentHeight++ {
		for _, valAddr := range blocks[currentHeight].Validators {

			//Get validator address from existed validator uptime count
			existedValAddr, pos := GetExistedAddress(valAddr, finalValAddrs)

			if pos > 0 {
				//Removing existed validator from uptime count
				finalValAddrs = append(finalValAddrs[:pos], finalValAddrs[pos+1:]...)

				//Incrementing uptime count
				existedValAddr.Info.UptimeCount++

				//Inserting existed validator into uptime count
				finalValAddrs = append(finalValAddrs, *existedValAddr)
			} else {
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
						StartBlock:   int64(currentHeight),
					},
				}

				//Inserting new validator into uptime count
				finalValAddrs = append(finalValAddrs, valAddressInfo)
			}
		}
	}

	//Printing Uptime results in tabular view
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 0, ' ', tabwriter.Debug)
	fmt.Fprintln(w, " Address\t Moniker\t Uptime Count")

	for _, data := range finalValAddrs {
		fmt.Fprintln(w, " "+data.ValAddress+"\t "+data.Info.Moniker+"\t  "+strconv.Itoa(int(data.Info.UptimeCount)))
	}

	w.Flush()

	//Exporing into csv file
	ExportIntoCsv(finalValAddrs)
}

func GetExistedAddress(validatorAddr string, finalValAddrs []ValidatorInfo) (*ValidatorInfo, int) {
	var valAddrs ValidatorInfo
	var pos int

	for index, addr := range finalValAddrs {
		if addr.ValAddress == validatorAddr {
			valAddrs = addr
			pos = index
		}
	}

	return &valAddrs, pos
}

func ExportIntoCsv(data []ValidatorInfo) {
	Header := []string{
		"Address", "Moniker", "Uptime Count",
	}

	file, err := os.Create("result.csv")

	//Error handling
	checkError("Cannot create file", err)

	defer file.Close() //Close file

	writer := csv.NewWriter(file)

	defer writer.Flush()

	//Write header titles
	_ = writer.Write(Header)

	for _, record := range data {
		uptimeCount := strconv.Itoa(int(record.Info.UptimeCount))
		addrObj := []string{record.ValAddress, record.Info.Moniker, uptimeCount}
		err := writer.Write(addrObj)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
