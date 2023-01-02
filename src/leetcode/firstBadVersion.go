// https://leetcode.com/problems/first-bad-version/

package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 *
 * Dummy
 */
func isBadVersion(v int) bool {
	return true
}

func FirstBadVersion(n int) int {
	l, h := 0, n
	var m int
	for l < h {
		m = (l + h) / 2
		if isBadVersion(m) {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
