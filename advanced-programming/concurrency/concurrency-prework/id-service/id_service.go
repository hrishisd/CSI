package idservice

import (
	"sync"
	"sync/atomic"
)

type idService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently without any additional synchronization.
	getNext() uint64
}

type noSyncService struct {
	currId uint64
}

func (s *noSyncService) getNext() uint64 {
	result := s.currId
	s.currId += 1
	return result
}

type atomicService struct {
	currId uint64
}

func (s *atomicService) getNext() uint64 {
	for {
		currVal := s.currId
		swapped := atomic.CompareAndSwapUint64(&s.currId, currVal, currVal+1)
		if swapped {
			return currVal + 1
		}
	}
}

type mutexIdService struct {
	currId uint64
	m      sync.Mutex
}

func (s *mutexIdService) getNext() uint64 {
	s.m.Lock()
	s.currId += 1
	result := s.currId
	s.m.Unlock()
	return result
}

type threadedIdService struct {
	ids <-chan uint64
}

func makeThreadedIdService() *threadedIdService {
	ids := make(chan uint64)
	go func() {
		var id uint64 = 0
		for {
			id += 1
			ids <- id
		}
	}()
	return &threadedIdService{ids: ids}
}

func (s *threadedIdService) getNext() uint64 {
	return <-s.ids
}
