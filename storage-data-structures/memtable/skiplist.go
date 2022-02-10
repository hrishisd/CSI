package memtable

import (
	"bytes"
	"errors"
	"math/rand"
)

const numLevels = 16

type SkipList struct {
	header *node
	level  int
}

type node struct {
	key     []byte
	value   []byte
	forward [numLevels]*node
}

func emptySkipList() SkipList {
	forward := [numLevels]*node{}
	header := &node{
		key:     []byte{},
		value:   []byte{},
		forward: forward,
	}
	return SkipList{header: header, level: 0}
}

func (s SkipList) Search(key []byte) ([]byte, error) {
	x := s.header
	for i := s.level; i >= 0; i-- {
		for x.forward[i] != nil && bytes.Compare(x.forward[i].key, key) < 0 {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && bytes.Equal(x.key, key) {
		return x.value, nil
	} else {
		return []byte{}, errors.New("key not present")
	}
}

func (s *SkipList) Insert(key []byte, value []byte) {
	path := [numLevels]*node{}
	x := s.header
	for i := s.level; i >= 0; i-- {
		for x.forward[i] != nil && bytes.Compare(x.forward[i].key, key) < 0 {
			x = x.forward[i]
		}
		path[i] = x
	}
	// Now, x.key < key <= x.forward[0].key
	x = x.forward[0]
	if x != nil && bytes.Equal(x.key, key) {
		x.value = value
	} else {
		level := randomLevel()
		if level > s.level {
			for i := s.level + 1; i <= level; i++ {
				path[i] = s.header
			}
			s.level = level
		}
		newNode := node{key: key, value: value, forward: [16]*node{}}
		for i := 0; i < s.level; i++ {
			newNode.forward[i] = path[i].forward[i]
			path[i].forward[i] = &newNode
		}
	}
}

func (s *SkipList) Delete(key []byte) bool {
	path := [numLevels]*node{}
	x := s.header
	for level := s.level; level >= 0; level-- {
		for x.forward[level] != nil && bytes.Compare(x.forward[level].key, key) < 0 {
			x = x.forward[level]
		}
		path[level] = x
	}
	// x.key < key <= x.forward[0].key
	x = x.forward[0]

	if x == nil || !bytes.Equal(x.key, key) {
		return false
	}
	for level := 0; level <= s.level; level++ {
		if path[level].forward[level] != x {
			break
		}
		path[level].forward[level] = x.forward[level]
	}
	for s.level >= 0 && s.header.forward[s.level] == nil {
		s.level--
	}
	return true
}

func (s SkipList) SearchNode(key []byte) *node {
	x := s.header
	for i := s.level; i >= 0; i-- {
		for x.forward[i] != nil && bytes.Compare(x.forward[i].key, key) < 0 {
			x = x.forward[i]
		}
	}
	return x.forward[0]
}

func randomLevel() int {
	level := 1
	for rand.Intn(2) == 1 && level <= numLevels-1 {
		level += 1
	}
	return level
}
