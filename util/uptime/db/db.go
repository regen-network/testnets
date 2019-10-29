package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//configuring db name and collections
var (
	DB_NAME               = "bigdipper_db"
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
func (db Store) FetchBlocks(startBlock int, endBlock int) ([]Blocks, error) {
	var blocks []Blocks
	err := db.session.DB(DB_NAME).C(BLOCKS_COLLECTION).Find(nil).Skip(startBlock).Limit(endBlock).All(&blocks)

	return blocks, err
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
		FetchBlocks(startBlock int, endBlock int) ([]Blocks, error)
		GetValidator(query bson.M) (Validator, error)
	}

	// Store will be used to satisfy the DB interface
	Store struct {
		session *mgo.Session
	}
)
