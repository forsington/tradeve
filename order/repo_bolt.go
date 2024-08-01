package order

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/hashicorp/go-hclog"
)

type BoltRepository struct {
	db *bolt.DB

	logger hclog.Logger
}

const (
	bucketName = "order"
)

func (r *BoltRepository) Get(typeId, region int) (Orders, error) {
	orders := Orders{}

	err := r.db.View(func(tx *bolt.Tx) error {
		orderBucket := tx.Bucket([]byte(bucketName))

		v := orderBucket.Get([]byte(strconv.Itoa(typeId)))

		if len(v) == 0 {
			return fmt.Errorf("no order found for %d", typeId)
		}
		err := json.Unmarshal(v, &orders)
		return err
	})
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *BoltRepository) Upsert(typeId int, orders Orders) error {
	json, err := json.Marshal(orders)
	if err != nil {
		return err
	}

	err = r.db.Update(func(tx *bolt.Tx) error {
		orderBucket := tx.Bucket([]byte(bucketName))

		err := orderBucket.Put([]byte(strconv.Itoa(typeId)), []byte(json))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *BoltRepository) initOrderBucket() error {
	// make sure the order bucket exists
	err := r.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func NewBoltRepository(db *bolt.DB, logger hclog.Logger) (Repository, error) {
	r := &BoltRepository{
		db:     db,
		logger: logger,
	}

	// set up the order bucket if it doesn't exist
	err := r.initOrderBucket()
	if err != nil {
		return nil, err
	}

	return r, nil
}
