package storage

import (
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func StartDB() {
	var err error
	db, err = bolt.Open("my.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
}

func WriteAnalysis(name string, mb MetricBank) {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		encoded, err := json.Marshal(mb)
		if err != nil {
			return fmt.Errorf("marshal metricbank: %s", err)
		}
		err1 := b.Put([]byte(name), []byte(encoded))
		if err1 != nil {
			log.Fatalf(err1.Error())
		}
		return nil
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func GetAnalysis(name string) MetricBank {
	var mb MetricBank
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(name))
		err1 := json.Unmarshal(v, &mb)
		if err1 != nil {
			return fmt.Errorf("Unmarshal byte: %s", err1)
		}
		return nil
	})
	if err != nil {
		log.Println("Key not present in DB for retrieval")
	}
	return mb
}

func GetProjectList() string {
	var mylist map[string]string = make(map[string]string)
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		if b != nil {
			c := b.Cursor()
			var mb MetricBank
			for k, v := c.First(); k != nil; k, v = c.Next() {
				err1 := json.Unmarshal(v, &mb)
				if err1 != nil {
					log.Println(err1)
				}
				mylist[string(k)] = mb.Date
			}
		} else {
			log.Println("Database is empty")
		}
		return nil
	})
	if err != nil {
		log.Println("Database is empty")
	}
	result, _ := json.Marshal(mylist)
	return string(result)
}

func DeleteAnalysis(name string) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		b.Delete([]byte(name))
		return nil
	})
	if err != nil {
		log.Println("Key not present in DB for deletion")
	}
}
