package storage

import (
    "sync"
)

type ReceiptStore struct {
    data map[string]int
    mu   sync.Mutex
}

var Store = &ReceiptStore{
    data: make(map[string]int),
}

func (s *ReceiptStore) Save(id string, points int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[id] = points
}

func (s *ReceiptStore) Get(id string) (int, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    points, exists := s.data[id]
    return points, exists
}

