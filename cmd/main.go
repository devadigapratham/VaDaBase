package main

import (
	"fmt"
	"log"

	"github.com/devadigapratham/VaDaBase/vadabase"
)

func main() {
	user := map[string]string{
		"name": "Prathamesh",
		"age":  "20",
	}

	_ = user
	db, err := vadabase.New()
	if err != nil {
		log.Fatal(err)
	}

	db.Insert("users", "name", "")
	// coll, err := db.CreateCollection("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%+v\n", coll)
}
