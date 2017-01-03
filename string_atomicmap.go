// Generated by: main
// TypeWriter: atomicmap
// Directive: +gen on String

package atomicstring

import (
	"sync"
	"sync/atomic"
)

// StringAtomicMap is a copy-on-write thread-safe map of string
type StringAtomicMap struct {
	mu  sync.Mutex
	val atomic.Value
}

type _StringMap map[string]string

// NewStringAtomicMap returns a new initialized StringAtomicMap
func NewStringAtomicMap() *StringAtomicMap {
	am := &StringAtomicMap{}
	am.val.Store(make(_StringMap, 0))
	return am
}

// Get returns a String for a given key
func (am *StringAtomicMap) Get(key string) (value string, ok bool) {
	value, ok = am.val.Load().(_StringMap)[key]
	return value, ok
}

// GetAll returns the underlying map of String
// this map must NOT be modified, to change the map safely use the Set and Delete
// functions and Get the value again
func (am *StringAtomicMap) GetAll() map[string]string {
	return am.val.Load().(_StringMap)
}

// Len returns the number of elements in the map
func (am *StringAtomicMap) Len() int {
	return len(am.val.Load().(_StringMap))
}

// Set inserts in the map a String under a given key
func (am *StringAtomicMap) Set(key string, value string) {
	am.mu.Lock()
	defer am.mu.Unlock()

	m1 := am.val.Load().(_StringMap)
	m2 := make(_StringMap, len(m1)+1)
	for k, v := range m1 {
		m2[k] = v
	}

	m2[key] = value
	am.val.Store(m2)
	return
}

// Delete removes the String under key from the map
func (am *StringAtomicMap) Delete(key string) {
	am.mu.Lock()
	defer am.mu.Unlock()

	m1 := am.val.Load().(_StringMap)
	_, ok := m1[key]
	if !ok {
		return
	}

	m2 := make(_StringMap, len(m1)-1)
	for k, v := range m1 {
		if k != key {
			m2[k] = v
		}
	}

	am.val.Store(m2)
	return
}
