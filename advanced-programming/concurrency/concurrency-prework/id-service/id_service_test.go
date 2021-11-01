package idservice

import (
	"fmt"
	"sync"
	"testing"
)

func TestServices(t *testing.T) {
	services := []idService{
		&atomicService{}, &mutexIdService{}, makeThreadedIdService(),
	}
	for _, service := range services {
		max := uint64(0)
		mutex := sync.Mutex{}
		concurrency := 1000
		wg := sync.WaitGroup{}
		wg.Add(concurrency)
		getId := func() {
			id := service.getNext()
			mutex.Lock()
			if id > max {
				max = id
			}
			mutex.Unlock()
			wg.Done()
		}
		for i := 0; i < concurrency; i++ {
			go getId()
		}
		wg.Wait()
		got := max
		want := concurrency
		if got != uint64(want) {
			t.Errorf("Want max id of %d but got %d from service %T", want, got, service)
		}
	}
}

func BenchmarkService(b *testing.B) {
	services := []idService{
		&noSyncService{}, &atomicService{}, &mutexIdService{}, makeThreadedIdService(),
	}
	concurrency := 500_000
	for _, s := range services {
		b.Run(fmt.Sprintf("%T", s), func(b *testing.B) {
			wg := sync.WaitGroup{}
			wg.Add(concurrency)

			callService := func() {
				defer wg.Done()
				s.getNext()
			}
			for i := 0; i < concurrency; i++ {
				go callService()
			}
			wg.Wait()
		})
	}
}
