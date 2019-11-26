package main

import (
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
}
