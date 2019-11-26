package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	//没有就创建
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))

		if b != nil {
			data := b.Get([]byte("l"))
			fmt.Printf("%s\n", data)
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
