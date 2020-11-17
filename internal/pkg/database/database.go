package database

import (
	"fmt"

	// use go-sqlite3 side effects
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeako/candidate-matcher/ent"
)

var db *ent.Client

// Init returns a new *ent.Client backed by sqlite3.
//
// If dbpath is empty, creates an in-memory db which isn't persisted. 
func Init(dbpath string) (*ent.Client, error) {
	var err error
	dataSourceName := "file:ent?mode=memory&cache=shared&_fk=1"
	if dbpath != "" {
		dataSourceName = fmt.Sprintf("file:%s?cache=shared&_fk=1", dbpath)
	}
	db, err = ent.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Get() *ent.Client {
	return db
}
