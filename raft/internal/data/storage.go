package data

import (
	"sync"

	"github.com/hashicorp/raft"
)

type Storage struct {
	Data map[string]int
	mu   sync.Mutex
}

func NewStorage() Storage {
	return Storage{
		Data: make(map[string]int),
	}
}

func (s *Storage) Get(key string) int {

}

func (s *Storage) Set(key string, value int) error {

}

func (s *Storage) Persist(sink raft.SnapshotSink) error {

}
