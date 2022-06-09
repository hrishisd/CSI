package memtable

import (
	"bytes"
	"io"
	"testing"
)

func TestSSTable(t *testing.T) {
	memtable := memtableOf("a", "a_value", "b", "b_value", "c", "c_value")
	var b bytes.Buffer
	memtable.flushSSTable(io.Writer(&b))
	t.Run("get", func(t *testing.T) {
		sstable := OpenSSTable(bytes.NewReader(b.Bytes()))
		value, err := sstable.Get([]byte("a"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if string(value) != "a_value" {
			t.Fatalf("wanted 'a_value' but got %v", string(value))
		}
		value, err = sstable.Get([]byte("b"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if string(value) != "b_value" {
			t.Fatalf("wanted 'b_value' but got %v", string(value))
		}
		value, err = sstable.Get([]byte("c"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if string(value) != "c_value" {
			t.Fatalf("wanted 'c_value' but got %v", string(value))
		}
		value, err = sstable.Get([]byte("d"))
		if value != nil {
			t.Fatalf("got non-nil value: %v", value)
		}
		if err == nil {
			t.Fatalf("expected error")
		}
	})
	t.Run("has", func(t *testing.T) {
		sstable := OpenSSTable(bytes.NewReader(b.Bytes()))
		has, err := sstable.Has([]byte("a"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if !has {
			t.Fatalf("got false")
		}
		has, err = sstable.Has([]byte("b"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if !has {
			t.Fatalf("got false")
		}
		has, err = sstable.Has([]byte("c"))
		if err != nil {
			t.Fatalf("got non-nil error: %v", err)
		}
		if !has {
			t.Fatalf("got false")
		}
	})
}

/// usage: memtableOf("a_key", "a_value", "b_key", "b_value")
func memtableOf(entries ...string) Memtable {
	result := EmptyMemtable()
	for i := 0; i < len(entries)-1; i += 2 {
		result.Put([]byte(entries[i]), []byte(entries[i+1]))
	}
	return result
}
