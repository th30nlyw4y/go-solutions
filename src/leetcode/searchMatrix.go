// https://leetcode.com/problems/search-a-2d-matrix/
package leetcode

func binarySearch(n *[]int, t int) bool {
	l, r := 0, len(*n)
	for l < r {
		m := (l + r) / 2
		if (*n)[m] == t {
			return true
		}
		if t >= (*n)[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	var line []int
	for _, a := range matrix {
		if a[0] <= target && target <= a[len(a)-1] {
			line = a
			break
		}
	}
	return binarySearch(&line, target)
}
