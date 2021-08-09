package reactor

import (
	"sync"
	"time"
)

type timeoutRegistry struct {
	timeouts map[uint32]chan bool
	lock     sync.Mutex
}

var once2 sync.Once
var tr *timeoutRegistry

func TimeoutRegistry() *timeoutRegistry {
	once2.Do(func() {
		tr = &timeoutRegistry{
			timeouts: make(map[uint32]chan bool, 0),
			lock:     sync.Mutex{},
		}
	})
	return tr
}

func (t *timeoutRegistry) Schedule(reactorId uint32, task func(), delay time.Duration) {
	go func() {
		timeoutChannel := make(chan bool)
		t.lock.Lock()
		t.timeouts[reactorId] = timeoutChannel
		t.lock.Unlock()

		select {
		case <-timeoutChannel:
			break
		case <-time.After(delay):
			task()
			break
		}
		t.lock.Lock()
		delete(t.timeouts, reactorId)
		t.lock.Unlock()
	}()
}

func (t *timeoutRegistry) Cancel(reactorId uint32) {
	if val, ok := t.timeouts[reactorId]; ok {
		val <- true
	}
}
