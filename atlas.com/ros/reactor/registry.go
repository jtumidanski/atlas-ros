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

var runningId = uint32(1000000001)

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

func (r *registry) Create(worldId byte, channelId byte, mapId uint32, classification uint32, name string, state int8, x int16, y int16, delay uint32, direction byte, statistics statistics.Model) Model {
	r.lock.Lock()
	id := r.getNextId()

	m := &Model{
		id:             id,
		worldId:        worldId,
		channelId:      channelId,
		mapId:          mapId,
		classification: classification,
		name:           name,
		statistics:     statistics,
		state:          state,
		delay:          delay,
		direction:      direction,
		x:              x,
		y:              y,
		alive:          true,
	}
	r.reactors[id] = m

	r.lock.Unlock()

	mk := MapKey{worldId, channelId, mapId}
	r.getMapLock(mk).Lock()
	if om, ok := r.mapReactors[mk]; ok {
		r.mapReactors[mk] = append(om, m.Id())
	} else {
		r.mapReactors[mk] = append([]uint32{}, m.Id())
	}
	r.getMapLock(mk).Unlock()
	return *m
}

func (r *registry) Update(id uint32, modifiers ...Modifier) (*Model, error) {
	r.lock.Lock()
	if val, ok := r.reactors[id]; ok {
		r.lock.Unlock()
		for _, modifier := range modifiers {
			modifier(val)
		}
		r.reactors[id] = val
		return val, nil
	} else {
		r.lock.Unlock()
		return nil, errors.New("unable to locate reactor")
	}
}

func (r *registry) getNextId() uint32 {
	ids := existingIds(r.reactors)

	var currentId = runningId
	for contains(ids, currentId) {
		currentId = currentId + 1
		if currentId > 2000000000 {
			currentId = 1000000001
		}
		runningId = currentId
	}
	return runningId
}

func (r *registry) Destroy(id uint32) {
	r.lock.Lock()
	val, ok := r.reactors[id]
	if !ok {
		return
	}
	delete(r.reactors, id)

	r.lock.Unlock()

	mk := MapKey{val.WorldId(), val.ChannelId(), val.MapId()}
	r.getMapLock(mk).Lock()
	if _, ok := r.mapReactors[mk]; ok {
		index := indexOf(id, r.mapReactors[mk])
		if index >= 0 && index < len(r.mapReactors[mk]) {
			r.mapReactors[mk] = remove(r.mapReactors[mk], index)
		}
	}
	r.getMapLock(mk).Unlock()
}

func existingIds(existing map[uint32]*Model) []uint32 {
	var ids []uint32
	for _, x := range existing {
		ids = append(ids, x.Id())
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

func indexOf(id uint32, data []uint32) int {
	for k, v := range data {
		if id == v {
			return k
		}
	}
	return -1 //not found.
}

func remove(s []uint32, i int) []uint32 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
