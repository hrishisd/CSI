package memtable

import (
	"fmt"
	"os"
)

const FLUSH_THRESHOLD int = 1 << 11

type LSMTree struct {
	memtable Memtable
	sstables []SSTable
	dir      string
}

// Get gets the value for the given key. It returns an error if the
// DB does not contain the key.
func (l *LSMTree) Get(key []byte) (value []byte, err error) {
	panic("not implemented") // TODO: Implement
}

// Has returns true if the DB contains the given key.
func (l *LSMTree) Has(key []byte) bool {
	panic("not implemented") // TODO: Implement
}

// Put sets the value for the given key. It overwrites any previous value
// for that key; a DB is not a multi-map.
func (l *LSMTree) Put(key []byte, value []byte) error {
	l.memtable.Put(key, value)
	if l.memtable.size > FLUSH_THRESHOLD {
		err := os.MkdirAll(l.dir, 0755)
		if err != nil {
			return fmt.Errorf("creating directory %q: %w", l.dir, err)
		}
		file, err := os.CreateTemp(l.dir, "sstable-*")
		if err != nil {
			return fmt.Errorf("error while creating sstable file: %w", err)
		}
		l.memtable.flushSSTable(file)
		newSSTable := OpenSSTable(file)
		l.sstables = append(l.sstables, *newSSTable)
	}
	return nil
}

// Delete deletes the value for the given key.
func (l *LSMTree) Delete(key []byte) {
	panic("not implemented") // TODO: Implement
}

// RangeScan returns an Iterator (see below) for scanning through all
// key-value pairs in the given range, ordered by key ascending.
func (l *LSMTree) RangeScan(start []byte, limit []byte) (Iterator, error) {
	panic("not implemented") // TODO: Implement
}
