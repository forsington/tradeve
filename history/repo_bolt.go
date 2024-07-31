package history

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
	bucketName = "history"
)

func (r *BoltRepository) Get(typeId, days, region int) (MarketHistories, error) {
	histories := MarketHistories{}

	err := r.db.View(func(tx *bolt.Tx) error {
		historyBucket := tx.Bucket([]byte(bucketName))

		v := historyBucket.Get([]byte(strconv.Itoa(typeId)))

		if len(v) == 0 {
			return fmt.Errorf("no history found for %d", typeId)
		}
		err := json.Unmarshal(v, &histories)
		return err
	})
	if err != nil {
		return nil, err
	}

	trimmed := histories.TrimDays(days)

	return trimmed, nil
}

func (r *BoltRepository) Upsert(typeId int, marketHistory MarketHistories) error {
	json, err := json.Marshal(marketHistory)
	if err != nil {
		return err
	}

	err = r.db.Update(func(tx *bolt.Tx) error {
		historyBucket := tx.Bucket([]byte(bucketName))

		err := historyBucket.Put([]byte(strconv.Itoa(typeId)), []byte(json))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *BoltRepository) initHistoryBucket() error {
	// make sure the history bucket exists
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

func NewBoltRepository(db *bolt.DB, logger hclog.Logger) (*BoltRepository, error) {
	r := &BoltRepository{
		db:     db,
		logger: logger,
	}

	// set up the history bucket if it doesn't exist
	err := r.initHistoryBucket()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *BoltRepository) Stop() {
	r.db.Close()
}
