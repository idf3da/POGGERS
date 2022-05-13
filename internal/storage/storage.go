package storage

import (
	"crypto/sha256"

	"github.com/pkg/errors"
)

type Store struct {
	Lib map[[32]byte][]byte
	T   string
}

func (s *Store) Add(data []byte) ([32]byte, error) {
	hash := sha256.Sum256(data)

	if _, ok := s.Lib[hash]; ok {
		return hash, nil
	}

	s.Lib[hash] = data

	return hash, nil
}

func (s *Store) Get(hash [32]byte) ([]byte, error) {
	if data, ok := s.Lib[hash]; ok {
		return data, nil
	}

	return nil, errors.WithMessagef(nil, "[%x] is not present in storage.", hash)
}

// TODO: Test fixed map size speed?
