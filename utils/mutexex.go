package utils

import (
	"sync"
)

type cache struct {
	mu sync.RWMutex
	M  map[string]int
}

var InMemCache = &cache{M: make(map[string]int)}

func Save(k string, v int, wg *sync.WaitGroup) {

	InMemCache.mu.Lock()
	InMemCache.M[k] = v
	InMemCache.mu.Unlock()
	wg.Done()
}

func saveSlow(k string, v int) {
	InMemCache.M[k] = v
}
