package global

import (
	"sync"
)

type SafeDict struct {
	data map[string]interface{}
	*sync.RWMutex
}

func NewSafeDict(d map[string]interface{}) *SafeDict {
	return &SafeDict{d, &sync.RWMutex{}}
}

func (d *SafeDict) Get(key string) (interface{}, bool) {
	d.RLock()
	defer d.RUnlock()
	value, ok := d.data[key]
	return value, ok
}

func (d *SafeDict) Put(key string, value interface{}) {
	d.Lock()
	defer d.Unlock()
	d.data[key] = value
}
