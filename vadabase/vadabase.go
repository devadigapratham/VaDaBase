package vadabase

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type M map[string]string

type Collection struct {
	*bbolt.Bucket
}

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

func (h *Vadabase) CreateCollection(name string) (*Collection, error) {
	coll := Collection{}
	err := h.db.Update(func(tx *bbolt.Tx) error {
		var (
			err    error
			bucket *bbolt.Bucket
		)
		bucket = tx.Bucket([]byte(name))
		if bucket == nil {
			bucket, err := tx.CreateBucket([]byte(name))
			if err != nil {
				return err
			}
		}
		coll.Bucket = bucket
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &coll, nil
}

func (h *Vadabase) Insert(collName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	coll, err := h.CreateCollection(collName)
	if err != nil {
		return id, err
	}
	for k, v := range data {
		if err := coll.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := coll.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}
}

func (h *Vadabase) String(coll string, k string, query any) {}

// db.Update(func(tx *bbolt.Tx) error {
// 	id := uuid.New().String()
// 	for k, v := range user {
// 		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
// 			return err
// 		}
// 	}

// 	if err := bucket.Put([]byte("id"), []byte(id)); err != nil {
// 		return err
// 	}

// 	return nil
// })
