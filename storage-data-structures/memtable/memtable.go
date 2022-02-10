package memtable

import "bytes"

type Memtable struct {
	list SkipList
}

func EmptyMemtable() Memtable {
	return Memtable{list: emptySkipList()}
}

func (m *Memtable) Get(key []byte) (value []byte, err error) {
	return m.list.Search(key)
}

func (m *Memtable) Has(key []byte) bool {
	_, err := m.list.Search(key)
	return err == nil
}

func (m *Memtable) Put(key, value []byte) {
	m.list.Insert(key, value)
}

func (m *Memtable) Delete(key []byte) {
	m.list.Delete(key)
}

// // RangeScan returns an Iterator (see below) for scanning through all
// // key-value pairs in the given range, ordered by key ascending.
func (m *Memtable) RangeScan(start, limit []byte) Iterator {
	return &MemtableIter{curr: m.list.SearchNode(start), limit: limit}
}

type MemtableIter struct {
	curr  *node
	limit []byte
}

// Next moves the iterator to the next key/value pair.
// It returns false if the iterator is exhausted.
func (m *MemtableIter) Next() bool {
	if m.exhausted() {
		return false
	}
	m.curr = m.curr.forward[0]
	return !m.exhausted()
}

func (m *MemtableIter) exhausted() bool {
	return m.curr == nil || bytes.Compare(m.curr.key, m.limit) > 0
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (m *MemtableIter) Error() error {
	// TODO: I don't understand this
	panic("not implemented")
}

// Key returns the key of the current key/value pair, or nil if done.
func (m *MemtableIter) Key() []byte {
	if m.curr == nil {
		return nil
	}
	return m.curr.key
}

// Value returns the value of the current key/value pair, or nil if done.
func (m *MemtableIter) Value() []byte {
	if m.curr == nil {
		return nil
	}
	return m.curr.value
}
