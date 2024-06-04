package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

func main() {
	// Open the database file
	db, err := bbolt.Open(".db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Define user data
	user := map[string]string{
		"name": "Prathamesh",
		"age":  "19",
	}

	// Store data in the database
	err = db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}

		id := uuid.New().String()
		for k, v := range user {
			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}

		if err := bucket.Put([]byte("id"), []byte(id)); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve data from the database
	userData := make(map[string]string)
	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("users"))
		if bucket == nil {
			return fmt.Errorf("Bucket (%s) not found!", "users")
		}

		err := bucket.ForEach(func(k, v []byte) error {
			userData[string(k)] = string(v)
			return nil
		})
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
