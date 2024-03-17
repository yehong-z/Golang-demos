package main

import (
	"fmt"
	bolt "github.com/boltdb/bolt"
)

func main() {

	db, _ := bolt.Open("my.db", 0600, nil)
	defer db.Close()

	// Start a writable transaction.
	tx, err := db.Begin(true)

	defer tx.Rollback()

	// Use the transaction...
	bucket, err := tx.CreateBucket([]byte("MyBucket"))

	bucket.Put([]byte("foo"), []byte("bar"))

	// Commit the transaction and check for error.
	if err = tx.Commit(); err != nil {
		fmt.Println(err.Error())
	}
}
