package memtable

import (
	"bytes"
	"errors"
	"sort"
)

type Memtable struct {
	table map[string][]byte
}

func EmptyMemtable() Memtable {
	return Memtable{make(map[string][]byte)}
}

func (m *Memtable) Get(key []byte) (value []byte, err error) {
	value, contains := m.table[string(key)]
	if !contains {
		return nil, errors.New("key not present in memtable")
	}
	return value, nil
}

func (m *Memtable) Has(key []byte) bool {
	_, contains := m.table[string(key)]
	return contains
}

func (m *Memtable) Put(key, value []byte) {
	m.table[string(key)] = value
}

func (m *Memtable) Delete(key []byte) {
	delete(m.table, string(key))
}

// // RangeScan returns an Iterator (see below) for scanning through all
// // key-value pairs in the given range, ordered by key ascending.
func (m *Memtable) RangeScan(start, limit []byte) Iterator {
	orderedKeys := make([]string, 0)
	for key := range m.table {
		keyBytes := []byte(key)
		if bytes.Compare(keyBytes, start) >= 0 && bytes.Compare(keyBytes, limit) < 0 {
			orderedKeys = append(orderedKeys, key)
		}
	}
	sort.Strings(orderedKeys)
	return &MemtableIter{orderedKeys: orderedKeys, idx: 0, memtable: m}
}

type MemtableIter struct {
	orderedKeys []string
	idx         int
	memtable    *Memtable
}

// Next moves the iterator to the next key/value pair.
// It returns false if the iterator is exhausted.
func (m *MemtableIter) Next() bool {
	m.idx += 1
	return m.idx < len(m.orderedKeys)
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (m *MemtableIter) Error() error {
	// TODO: I don't understand this
	panic("not implemented") // TODO: Implement
}

// Key returns the key of the current key/value pair, or nil if done.
func (m *MemtableIter) Key() []byte {
	return []byte(m.orderedKeys[m.idx])
}

// Value returns the value of the current key/value pair, or nil if done.
func (m *MemtableIter) Value() []byte {
	key := m.orderedKeys[m.idx]
	return m.memtable.table[key]
}
