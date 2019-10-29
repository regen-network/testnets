package db

import (
	mgo "gopkg.in/mgo.v2"
)

// Connect returns a pointer to a MongoDB instance,
// which is used for collecting the metrics required for uptime calculations
func Connect(uri string) (DB, error) {
	session, err := mgo.Dial(uri)

	return Store{session: session}, err
}

// Terminate should be used to terminate a database session, generally in a defer statement inside main app file.
func (db Store) Terminate() {
	db.session.Close()
}

type (
	// DB interface defines all the methods accessible by the application
	DB interface {
		Terminate()
	}

	// Store will be used to satisfy the DB interface
	Store struct {
		session *mgo.Session
	}
)
