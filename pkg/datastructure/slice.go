package datastructure

import "sort"

func SliceBinaryInsert(s []int, x int) []int {
	idx := sort.SearchInts(s, x)
	s = append(s, 0)
	copy(s[idx+1:], s[idx:])
	s[idx] = x
	return s
}
