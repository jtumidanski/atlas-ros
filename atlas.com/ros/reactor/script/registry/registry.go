package registry

import (
	"atlas-ros/reactor/script"
	"errors"
	"sync"
)

type Registry struct {
	registry map[uint32]script.Script
}

var once sync.Once
var registry *Registry

func GetRegistry() *Registry {
	once.Do(func() {

		registry = initRegistry()
	})
	return registry
}

func initRegistry() *Registry {
	r := &Registry{make(map[uint32]script.Script)}
	return r
}

func (r *Registry) AddScripts(provider func() []script.Script) {
	for _, s := range provider() {
		r.registry[s.ReactorClassification()] = s
	}
}

func (r *Registry) GetScript(reactorClassification uint32) (*script.Script, error) {
	if val, ok := r.registry[reactorClassification]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}
