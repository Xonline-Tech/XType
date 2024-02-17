package backend

import (
	"github.com/dgraph-io/badger/v4"
	"testing"
)

func TestSetStruct(t *testing.T) {
	opts := new(badger.Options)
	*opts = badger.DefaultOptions("./tmp/badgerDBTest")
	InitDB(opts)
	//runBadgerTest(t, nil, func(t *testing.T, db *badger.DB) {
	CreateObjectType("testObjectType5", "TestValue")
	//})
}

func TestQueryAll(t *testing.T) {
	opts := new(badger.Options)
	*opts = badger.DefaultOptions("./tmp/badgerDBTest")
	InitDB(opts)
	QueryAll()
}
