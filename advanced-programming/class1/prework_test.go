package prework

import "testing"

func TestStrlen(t *testing.T) {
	for _, s := range [...]string{"hello", "world", "lalalalal"} {
		got := strlen(s)
		want := len(s)
		if got != want {
			t.Errorf("got %d want %d given %s", got, want, s)
		}
	}
}

func TestGetY(t *testing.T) {
	for _, p := range [...]Point{{}, {x: 5}, {y: 6}, {x: 4, y: 7}} {
		got := getY(p)
		want := p.y
		if got != want {
			t.Errorf("got %d want %d given %q", got, want, p)
		}
	}
}

func TestUnsafeSum(t *testing.T) {
	referenceSum := func(nums []int) int {
		res := 0
		for _, num := range nums {
			res += num
		}
		return res
	}

	referenceSum([]int{1, 2, 3})
	ints := []int{1, 2, 3, 4}
	got := unsafeSum(ints)
	want := 10
	if got != want {
		t.Errorf("got %d want %d given %d", got, want, ints)
	}
}
