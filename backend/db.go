package backend

import (
	"github.com/Xonline-Tech/XType/backend/common"
	"github.com/dgraph-io/badger/v4"
	"log"
)

var db *badger.DB

// InitDB init DB
func InitDB(opts *badger.Options) {
	if opts == nil {
		opts = new(badger.Options)
		*opts = badger.DefaultOptions("./tmp/badgerDB")
	}

	var err error

	if db, err = badger.Open(*opts); err != nil {
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

func SetStruct(key string, val any) {
	byteSlice := common.StructToByteSlice(&val)
	err := db.Update(func(txn *badger.Txn) error {
		if err := txn.Set([]byte(key), byteSlice); err != nil {
			log.Println(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func GetPrefix(keyPrefix string) {
	err := db.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iterator.Close()
		for iterator.Seek([]byte(keyPrefix)); iterator.ValidForPrefix([]byte(keyPrefix)); iterator.Next() {
			item := iterator.Item()
			key := item.Key()
			if err := item.Value(func(val []byte) error {
				log.Printf("获取testKey:的key:%s", key)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}
}
