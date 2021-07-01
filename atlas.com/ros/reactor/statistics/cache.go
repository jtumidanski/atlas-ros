package statistics

import (
	"sync"
)

type statCache struct {
	stats map[uint32]Model
	lock  sync.RWMutex
}

var cache *statCache
var once sync.Once

func GetCache() *statCache {
	once.Do(func() {
		cache = &statCache{
			stats: make(map[uint32]Model, 0),
			lock:  sync.RWMutex{},
		}
	})
	return cache
}

func (e *statCache) GetFile(id uint32) (*Model, error) {
	e.lock.RLock()
	if val, ok := e.stats[id]; ok {
		e.lock.RUnlock()
		return &val, nil
	} else {
		e.lock.RUnlock()
		e.lock.Lock()
		s, err := readStatistics(id)
		if err != nil {
			e.lock.Unlock()
			return nil, err
		}
		e.stats[id] = *s
		e.lock.Unlock()
		return s, nil
	}
}
