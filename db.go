package main

import (
	"github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func OpenDB(path string) error {
	opts := badger.DefaultOptions(path).WithLogger(nil)
	var err error
	DB, err = badger.Open(opts)
	return err
}

func GetAllKeys() ([]string, error) {
	var keys []string
	err := DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			keys = append(keys, string(item.Key()))
		}
		return nil
	})
	return keys, err
}

func GetValue(key string) (string, error) {
	var val string
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		v, err := item.ValueCopy(nil)
		val = string(v)
		return err
	})
	return val, err
}
