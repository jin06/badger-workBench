package main

import (
	"time"

	"github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

func OpenDB(path string) error {
	// Open the database with the given path
	opts := badger.DefaultOptions(path).WithLogger(nil)
	var err error
	DB, err = badger.Open(opts)
	return err
}

func GetAllKeys() ([]string, error) {
	// Retrieve all keys from the database
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
	// Retrieve the value for a given key
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

func SetValueWithTTL(key, value string, ttl uint64) error {
	// Set a key-value pair with an optional TTL
	return DB.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), []byte(value))
		if ttl > 0 {
			entry.WithTTL(time.Duration(ttl) * time.Second) // Set TTL
		}
		return txn.SetEntry(entry)
	})
}

func GetKeyTTL(key string) (uint64, error) {
	// Retrieve the remaining TTL for a given key
	var ttl uint64
	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		if item.ExpiresAt() == 0 {
			ttl = 0
		} else {
			ttl = item.ExpiresAt() - uint64(time.Now().Unix())
		}
		return nil
	})
	return ttl, err
}
