package backend

import (
	"github.com/asdine/storm"
)

type StormStorageConfig struct {
	DB       *storm.DB
	Disabled bool
	IsReady  bool
}

func (s *StormStorageConfig) Open(filePath string) (*storm.DB, error) {
	var err error
	s.DB, err = storm.Open(filePath, storm.AutoIncrement())
	if err == nil {
		s.IsReady = true
	}
	return s.DB, err
}

func (s *StormStorageConfig) Close() error {
	s.IsReady = false
	return s.DB.Close()
}
