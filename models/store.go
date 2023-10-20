package models

import "sync"

type KeyValueStore struct {
	Store map[string]string
	mutex sync.RWMutex
}

func (kvs *KeyValueStore) Get(key string) string {
	kvs.mutex.RLock()
	defer kvs.mutex.RUnlock()
	return kvs.Store[key]
}
func (kvs *KeyValueStore) Put(key, value string) {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	kvs.Store[key] = value
}
func (kvs *KeyValueStore) Delete(key string) {
	kvs.mutex.Lock()
	defer kvs.mutex.Unlock()
	delete(kvs.Store, key)
}
