package memtable

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
)

// SSTable on-disk format:
// flat sequence of (key length, value length, key, value)
// key length and value length are uint16 and are written as big endian

type SSTable struct {
	reader io.ReadSeeker
}

func OpenSSTable(reader io.ReadSeeker) *SSTable {
	return &SSTable{reader}
}

// Get gets the value for the given key. It returns an error if the
// DB does not contain the key.
func (s *SSTable) Get(key []byte) (value []byte, err error) {
	if _, err := s.reader.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	// what if key is not in table?
	for {
		var keyLength uint16
		var valueLength uint16
		// todo handle errors
		err := binary.Read(s.reader, binary.BigEndian, &keyLength)
		if err == io.EOF {
			return nil, errors.New("Key not found in sstable")
		}
		binary.Read(s.reader, binary.BigEndian, &valueLength)
		currKey := make([]byte, keyLength)
		// handle error
		s.reader.Read(currKey)
		if bytes.Equal(key, currKey) {
			value := make([]byte, valueLength)
			// handle error
			s.reader.Read(value)
			return value, nil
		}
		s.reader.Seek(int64(valueLength), io.SeekCurrent)
	}
}

// Has returns true if the DB contains the given key.
func (s *SSTable) Has(key []byte) (ret bool, err error) {
	if _, err := s.reader.Seek(0, io.SeekStart); err != nil {
		return false, err
	}
	// what if key is not in table?
	for {
		var keyLength uint16
		var valueLength uint16
		// todo handle errors
		err := binary.Read(s.reader, binary.BigEndian, &keyLength)
		if err == io.EOF {
			return false, nil
		}
		if err = binary.Read(s.reader, binary.BigEndian, &valueLength); err != nil {
			return false, err
		}
		currKey := make([]byte, keyLength)
		// handle error
		s.reader.Read(currKey)
		if bytes.Equal(key, currKey) {
			return true, nil
		}
		s.reader.Seek(int64(valueLength), io.SeekCurrent)
	}
}

// RangeScan returns an Iterator (see below) for scanning through all
// key-value pairs in the given range, ordered by key ascending.
func (s *SSTable) RangeScan(start, limit []byte) (Iterator, error) {
	_, err := s.reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return &SSTableIterator{
		reader: s.reader,
		start:  start,
		limit:  limit,
	}, nil
}

type SSTableIterator struct {
	reader    io.ReadSeeker
	start     []byte
	limit     []byte
	current   *Entry
	err       error
	exhausted bool
}

type Entry struct {
	key   []byte
	value []byte
}

// Next moves the iterator to the next key/value pair.
// It returns false if the iterator is exhausted.
func (s *SSTableIterator) Next() bool {
	if s.exhausted {
		return false
	}
	var keyLength uint16
	var valueLength uint16
	err := binary.Read(s.reader, binary.BigEndian, &keyLength)
	if err == io.EOF {
		s.exhausted = true
		return false
	}
	binary.Read(s.reader, binary.BigEndian, &valueLength)
	s.reader.Read(s.current.key)
	if bytes.Compare(s.current.key, s.limit) >= 0 {
		s.exhausted = true
		return false
	}
	s.reader.Read(s.current.value)
	return true
}

// Error returns any accumulated error. Exhausting all the key/value pairs
// is not considered to be an error.
func (s *SSTableIterator) Error() error {
	return s.err
}

// Key returns the key of the current key/value pair, or nil if done.
func (s *SSTableIterator) Key() []byte {
	return s.current.key
}

// Value returns the value of the current key/value pair, or nil if done.
func (s *SSTableIterator) Value() []byte {
	return s.current.value
}
