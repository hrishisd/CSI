package memtable

import (
	"bytes"
	"testing"
)

func TestSkipList_Search(t *testing.T) {
	tests := []struct {
		name       string
		keys       []string
		search_key []byte
		want       []byte
		wantErr    bool
	}{
		{
			name:       "search empty",
			keys:       []string{},
			search_key: []byte("missingkey"),
			want:       []byte{},
			wantErr:    true,
		},
		{
			name:       "search singleton list",
			keys:       []string{"a"},
			search_key: []byte("a"),
			want:       []byte("a"),
			wantErr:    false,
		},
		{
			name:       "present",
			keys:       []string{"a", "b", "c", "d", "e", "f", "g"},
			search_key: []byte("d"),
			want:       []byte("d"),
			wantErr:    false,
		},
		{
			name:       "missing, larger than all keys",
			keys:       []string{"a", "b", "c", "d", "e", "f", "g"},
			search_key: []byte("z"),
			want:       []byte{},
			wantErr:    true,
		},
		{
			name:       "missing, smaller than all keys",
			keys:       []string{"a", "b", "c", "d", "e", "f", "g"},
			search_key: []byte("0"),
			want:       []byte{},
			wantErr:    true,
		},
		{
			name:       "missing, in between keys",
			keys:       []string{"a", "b", "c", "e", "f", "g"},
			search_key: []byte("d"),
			want:       []byte{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := skipListWithKeys(tt.keys...)
			got, err := s.Search(tt.search_key)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkipList.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("SkipList.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkipList_Insert(t *testing.T) {
	t.Run("insert single element", func(t *testing.T) {
		s := emptySkipList()
		s.Insert([]byte("key"), []byte("value"))
		got, err := s.Search([]byte("key"))
		want := []byte("value")
		if err != nil {
			t.Errorf("SkipList.Search() error = %v, wantErr %v", err, nil)
			return
		}
		if !bytes.Equal(got, want) {
			t.Errorf("SkipList.Search() = %v, want %v", got, want)
		}
	})
	t.Run("insert duplicate element", func(t *testing.T) {
		s := emptySkipList()
		s.Insert([]byte("key"), []byte("value"))
		s.Insert([]byte("key"), []byte("value"))
		got, err := s.Search([]byte("key"))
		want := []byte("value")
		if err != nil {
			t.Errorf("SkipList.Search() error = %v, wantErr %v", err, nil)
			return
		}
		if !bytes.Equal(got, want) {
			t.Errorf("SkipList.Search() = %v, want %v", got, want)
		}
	})
}

func TestSkipList_Delete(t *testing.T) {
	t.Run("delete from empty list", func(t *testing.T) {
		s := emptySkipList()
		got := s.Delete([]byte("key"))
		want := false
		if got != want {
			t.Errorf("SkipList.Delete() = %v, want %v", got, "value")
		}
	})
	t.Run("delete from singleton list", func(t *testing.T) {
		s := skipListWithKeys("key")
		got := s.Delete([]byte("key"))
		want := true
		if got != want {
			t.Errorf("SkipList.Delete() = %v, want %v", got, "value")
		}
	})
	t.Run("delete all", func(t *testing.T) {
		keys := []string{"a", "b", "c", "d", "e", "f", "g"}
		s := skipListWithKeys(keys...)
		for _, key := range keys {
			key := []byte(key)
			got := s.Delete(key)
			want := true
			if got != want {
				t.Errorf("SkipList.Delete() = %v, want %v", got, "value")
			}
		}
		for _, key := range keys {
			key := []byte(key)
			got := s.Delete(key)
			want := false
			if got != want {
				t.Errorf("SkipList.Delete() = %v, want %v", got, "value")
			}
		}
	})
}

// Constructs a skip list with the provided keys.
// The value associated with each key is the the same value as the key.
func skipListWithKeys(keys ...string) SkipList {
	result := emptySkipList()
	for _, key := range keys {
		result.Insert([]byte(key), []byte(key))
	}
	return result
}
