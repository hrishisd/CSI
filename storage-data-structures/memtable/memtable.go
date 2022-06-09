package memtable

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Memtable struct {
	list SkipList
	size int
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

func (m *Memtable) Put(key, value []byte) error {
	prevVal := m.list.Insert(key, value)
	if prevVal == nil {
		m.size += len(key) + len(value)
	} else {
		m.size += len(value) - len(prevVal)
	}
	return nil
}

func (m *Memtable) Delete(key []byte) {
	m.list.Delete(key)
}

// // RangeScan returns an Iterator (see below) for scanning through all
// // key-value pairs in the given range, ordered by key ascending.
func (m *Memtable) RangeScan(start, limit []byte) Iterator {
	return &MemtableIter{curr: m.list.SearchNode(start), limit: limit}
}

// flat sequence of (key length, value length, key, value)
// key length and value length are uint16 each
func (m *Memtable) flushSSTable(w io.Writer) error {
	node := m.list.header.forward[0]
	sizeBuf := make([]byte, 2)
	for node != nil {
		keyLen := uint16(len(node.key))
		binary.BigEndian.PutUint16(sizeBuf, keyLen)
		_, err := w.Write(sizeBuf)
		if err != nil {
			return err
		}
		valLen := uint16(len(node.value))
		binary.BigEndian.PutUint16(sizeBuf, valLen)
		_, err = w.Write(sizeBuf)
		if err != nil {
			return err
		}
		_, err = w.Write(node.key)
		if err != nil {
			return err
		}
		_, err = w.Write(node.value)
		if err != nil {
			return err
		}
		node = node.forward[0]
	}
	return nil
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
