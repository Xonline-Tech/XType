package backend

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func removeDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		panic(err)
	}
}

func getTestOptions(dir string) badger.Options {
	opt := badger.DefaultOptions(dir).
		WithSyncWrites(false).
		WithLoggingLevel(badger.WARNING)
	return opt
}

func runBadgerTest(t *testing.T, opts *badger.Options, test func(t *testing.T, db *badger.DB)) {
	dir, err := os.MkdirTemp("", "badger-test")
	require.NoError(t, err)
	//defer removeDir(dir)
	if opts == nil {
		opts = new(badger.Options)
		*opts = getTestOptions(dir)
	} else {
		opts.Dir = dir
		opts.ValueDir = dir
	}

	if opts.InMemory {
		opts.Dir = ""
		opts.ValueDir = ""
	}
	db, err := badger.Open(*opts)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, db.Close())
	}()
	test(t, db)
}
