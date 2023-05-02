// https://leetcode.com/problems/isomorphic-strings/
package leetcode

func IsIsomorphic(s string, t string) bool {
	sM, tM, i := map[byte]byte{}, map[byte]byte{}, 0
	for i < len(s) {
		sV, sOk := sM[s[i]]
		tV, tOk := tM[t[i]]
		if (sOk && sV != t[i]) || (tOk && tV != s[i]) {
			return false
		} else {
			sM[s[i]], tM[t[i]] = t[i], s[i]
		}
		i++
	}
	return true
}
