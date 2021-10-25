package prework

import (
	"unsafe"
)

func strlen(s string) int {
	len_ptr := uintptr(unsafe.Pointer(&s)) + 8
	return *(*int)(unsafe.Pointer(len_ptr))
}

type Point struct {
	x int
	y int
}

func getY(p Point) int {
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.y)))
}

func unsafeSum(ints []int) int {
	n := uint(len(ints))
	data_start := unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&ints)))
	sum := 0
	for i := uint(0); i < n; i++ {
		curr_ptr := unsafe.Add(data_start, uint(8)*i)
		curr := *(*int)(curr_ptr)
		sum += curr
	}
	return sum
}
