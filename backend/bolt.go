package backend

import (
	"github.com/boltdb/bolt"
)

type BoltStorageConfig struct {
	DB       *bolt.DB
	Disabled bool
	IsReady  bool
}
