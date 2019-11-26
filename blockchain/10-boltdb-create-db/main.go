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

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create  bucket:%s", err)
		}
		if b != nil {
			err := b.Put([]byte("l"), []byte("send 100 BTC To Q"))
			if err != nil {
				log.Panic("数据存储失败")
			}
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
