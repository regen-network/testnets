package db

import (
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//configuring db name and collections
var (
	DB_NAME, dbErr        = viper.Get("database").(string)
	BLOCKS_COLLECTION     = "blocks"
	VALIDATORS_COLLECTION = "validators"
)

type Blocks struct {
	ID         string   `json:"_id" bson:"_id"`
	Height     int64    `json:"height" bson:"height"`
	Validators []string `json:"validators" bson:"validators"`
}

type Validator struct {
	Address         string      `json:"address" bson:"address"`
	OperatorAddress string      `json:"operatorAddress" bson:"operator_address"`
	Description     Description `json:"description" bson:"description"`
}

type BlocksAggResult struct {
	Id 			string 		`json:"_id" bson.M:"_id"`
	Uptime_count  int64 	`json:"uptime_count" bson:"uptime_count"`
	Upgrade1_block	int64	`json:"upgrade1_block" bson:"upgrade1_block"`
	Upgrade2_block  int64 	`json:"upgrade2_block" bson:"upgrade2_block"`
	Validator_details []Validator_details `json:"validator_details" bson:"validator_details"`
}

type Validator_details struct {
	Delegator_address string      `json:"delegator_address" bson:"delegator_address"`
	Description       Description `json:"description" bson:"description"`
	Operator_address  string      `json:"operator_address" bson:"operator_address"`
}

type Description struct {
	Moniker string `json:"moniker" bson:"moniker"`
}

// Connect returns a pointer to a MongoDB instance,
// which is used for collecting the metrics required for uptime calculations
func Connect(info *mgo.DialInfo) (DB, error) {
	session, err := mgo.DialWithInfo(info)

	return Store{session: session}, err
}

// Terminate should be used to terminate a database session, generally in a defer statement inside main app file.
func (db Store) Terminate() {
	db.session.Close()
}

// FetchBlocks read the blocks data
func (db Store) FetchBlocks(startBlock int64, endBlock int64) ([]Blocks, error) {
	var blocks []Blocks

	andQuery := bson.M{"height": bson.M{"$gte": startBlock, "$lte": endBlock}}

	err := db.session.DB(DB_NAME).C(BLOCKS_COLLECTION).Find(andQuery).Sort("height").All(&blocks)

	return blocks, err
}

//Fetech all blocks by using aggregate
func (db Store) FetchAllBlocksByAgg(aggQuery []bson.M) (result []BlocksAggResult, err error)  {
	err = db.session.DB(DB_NAME).C(BLOCKS_COLLECTION).Pipe(aggQuery).All(&result)
	return result, err
}

//Get block by height
func (db Store) GetBlockByHeight(query bson.M) (Blocks, error) {
	var block Blocks
	err := db.session.DB(DB_NAME).C(BLOCKS_COLLECTION).Find(query).One(&block)
	return block, err
}

//GetValidator Read single validator info
func (db Store) GetValidator(query bson.M) (Validator, error) {
	var val Validator
	err := db.session.DB(DB_NAME).C(VALIDATORS_COLLECTION).Find(query).One(&val)

	return val, err
}

type (
	// DB interface defines all the methods accessible by the application
	DB interface {
		Terminate()
		FetchBlocks(startBlock int64, endBlock int64) ([]Blocks, error)
		GetValidator(query bson.M) (Validator, error)
		GetBlockByHeight(query bson.M) (Blocks, error)
		FetchAllBlocksByAgg(aggQuery []bson.M) ([]BlocksAggResult, error)
	}

	// Store will be used to satisfy the DB interface
	Store struct {
		session *mgo.Session
	}
)
