package backend

import (
	"github.com/dgraph-io/badger/v4"
	"log"
)

var db *badger.DB

// InitDB init DB
func InitDB() {
	options := badger.DefaultOptions("./tmp/badgerDB")

	var err error

	if db, err = badger.Open(options); err != nil {
		log.Println("Database open failed!", err)
	} else {
		log.Println("Database opened successfully!")
	}

}

// CloseDB close DB
func CloseDB() {
	if err := db.Close(); err != nil {
		log.Panicln("Database close failed!")
	}
}
