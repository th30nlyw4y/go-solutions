// https://leetcode.com/problems/detect-capital/

package leetcode

func DetectCapitalUse(word string) bool {
	cnt := 0
	for _, letter := range word {
		if letter < 91 { // Since we have only lowercase & uppercase letters (stated in the problem description), we may not check letter > 64
			cnt += 1
		}
	}
	switch cnt {
	case len(word):
		return true
	case 0:
		return true
	case 1:
		return word[0] < 'Z'
	default:
		return false
	}
}
