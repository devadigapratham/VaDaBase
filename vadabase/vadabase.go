package vadabase

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type Vadabase struct {
	db *bbolt.DB
}

func New() (*Vadabase, error) {
	dbname := fmt.Sprintf("%s.vadabase", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Vadabase{
		db: db,
	}, nil
}
