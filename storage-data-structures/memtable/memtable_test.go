package memtable

import (
	"reflect"
	"testing"
)

func TestPutAndHas(t *testing.T) {
	memtable := EmptyMemtable()
	memtable.Put([]byte("key"), []byte("value"))
	has := memtable.Has([]byte("key"))
	if has != true {
		t.Errorf("Got false but expected true")
	}
}

func TestHasNotPresent(t *testing.T) {
	memtable := EmptyMemtable()
	has := memtable.Has([]byte("key"))
	if has != false {
		t.Errorf("Got true but expected false")
	}
}

func TestGetNotPresent(t *testing.T) {
	memtable := EmptyMemtable()
	_, err := memtable.Get([]byte("key"))
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestPutAndGet(t *testing.T) {
	memtable := EmptyMemtable()
	memtable.Put([]byte("key"), []byte("value"))
	value, err := memtable.Get([]byte("key"))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(value) != "value" {
		t.Errorf("Expected %s, got %s", "value", string(value))
	}
}

func TestPutTwice(t *testing.T) {
	memtable := EmptyMemtable()
	memtable.Put([]byte("key"), []byte("value"))
	memtable.Put([]byte("key"), []byte("value"))
}

func TestDelete(t *testing.T) {
	memtable := EmptyMemtable()
	memtable.Put([]byte("key"), []byte("value"))
	memtable.Delete([]byte("key"))
	if memtable.Has([]byte("key")) {
		t.Errorf("key should have been removed from memtable")
	}
}

func TestRangeScan(t *testing.T) {
	memtable := EmptyMemtable()
	memtable.Put([]byte("a"), []byte("a_val"))
	memtable.Put([]byte("b"), []byte("b_val"))
	memtable.Put([]byte("c"), []byte("c_val"))
	memtable.Put([]byte("d"), []byte("d_val"))
	memtable.Put([]byte("e"), []byte("e_val"))

	t.Run("empty range", func(t *testing.T) {
		iter := memtable.RangeScan([]byte("a"), []byte("a"))
		if iter.Next() {
			t.Error("scan should be empty")
		}
	})
	t.Run("full range", func(t *testing.T) {
		iter := memtable.RangeScan([]byte("a"), []byte("z"))
		got := collect(iter)
		want := [][]byte{[]byte("a_val"), []byte("b_val"), []byte("c_val"), []byte("d_val"), []byte("e_val")}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v but got %v", want, got)
		}
	})
}

func collect(i Iterator) [][]byte {
	result := make([][]byte, 0)
	for {
		val := i.Value()
		if val != nil {
			result = append(result, val)
		}
		if !i.Next() {
			break
		}
	}
	return result
}
