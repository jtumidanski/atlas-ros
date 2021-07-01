package reactor

import (
	"atlas-ros/reactor/statistics"
	"errors"
	"sync"
)

type registry struct {
	reactors    map[uint32]*Model
	mapReactors map[MapKey][]uint32
	mapLocks    map[MapKey]*sync.Mutex
	lock        sync.RWMutex
}

var once sync.Once
var reg *registry

var uniqueId = uint32(1000000001)

type MapKey struct {
	worldId   byte
	channelId byte
	mapId     uint32
}

func GetRegistry() *registry {
	once.Do(func() {
		reg = &registry{
			reactors:    make(map[uint32]*Model, 0),
			mapReactors: make(map[MapKey][]uint32),
			mapLocks:    make(map[MapKey]*sync.Mutex, 0),
			lock:        sync.RWMutex{},
		}
	})
	return reg
}

func (r *registry) Get(id uint32) (*Model, error) {
	r.lock.RLock()
	if val, ok := r.reactors[id]; ok {
		r.lock.RUnlock()
		return val, nil
	} else {
		r.lock.RUnlock()
		return nil, errors.New("unable to locate reactor")
	}
}

func (r *registry) GetInMap(worldId byte, channelId byte, mapId uint32) []Model {
	mk := MapKey{worldId, channelId, mapId}
	r.getMapLock(mk).Lock()
	r.lock.RLock()
	var result []Model
	for _, x := range r.mapReactors[mk] {
		result = append(result, *r.reactors[x])
	}
	r.lock.RUnlock()
	r.getMapLock(mk).Unlock()
	return result
}

func (r *registry) getMapLock(key MapKey) *sync.Mutex {
	r.lock.RLock()
	if val, ok := r.mapLocks[key]; ok {
		r.lock.RUnlock()
		return val
	} else {
		r.lock.RUnlock()
		var cm = &sync.Mutex{}
		r.lock.Lock()
		r.mapLocks[key] = cm
		r.lock.Unlock()
		return cm
	}
}

func (r *registry) Create(worldId byte, channelId byte, mapId uint32, reactorId uint32, name string, state byte, x int16, y int16, delay uint32, direction byte, statistics statistics.Model) Model {
	r.lock.Lock()
	uid := r.getNextUniqueId()

	m := &Model{
		uniqueId:   uid,
		worldId:    worldId,
		channelId:  channelId,
		mapId:      mapId,
		id:         reactorId,
		name:       name,
		statistics: statistics,
		state:      state,
		delay:      delay,
		direction:  direction,
		x:          x,
		y:          y,
		alive:      true,
	}
	r.reactors[uid] = m

	r.lock.Unlock()

	mk := MapKey{worldId, channelId, mapId}
	r.getMapLock(mk).Lock()
	if om, ok := r.mapReactors[mk]; ok {
		r.mapReactors[mk] = append(om, m.UniqueId())
	} else {
		r.mapReactors[mk] = append([]uint32{}, m.UniqueId())
	}
	r.getMapLock(mk).Unlock()
	return *m
}

func (r *registry) getNextUniqueId() uint32 {
	ids := existingIds(r.reactors)

	var currentUniqueId = uniqueId
	for contains(ids, currentUniqueId) {
		currentUniqueId = currentUniqueId + 1
		if currentUniqueId > 2000000000 {
			currentUniqueId = 1000000001
		}
		uniqueId = currentUniqueId
	}
	return uniqueId
}

func existingIds(existing map[uint32]*Model) []uint32 {
	var ids []uint32
	for _, x := range existing {
		ids = append(ids, x.UniqueId())
	}
	return ids
}

func contains(ids []uint32, id uint32) bool {
	for _, element := range ids {
		if element == id {
			return true
		}
	}
	return false
}
