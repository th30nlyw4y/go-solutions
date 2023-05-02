// https://leetcode.com/problems/word-pattern/
package leetcode

import "strings"

func WordPattern(pattern string, s string) bool {
	letterMappings, stringMappings := map[rune]string{}, map[string]rune{}
	sSplit := strings.Split(s, " ")
	if len(pattern) != len(sSplit) {
		return false
	}
	for index, letter := range pattern {
		if letterMapping, ok := letterMappings[letter]; ok && letterMapping != sSplit[index] {
			return false
		} else if stringMapping, ok := stringMappings[sSplit[index]]; ok && stringMapping != letter {
			return false
		} else {
			letterMappings[letter] = sSplit[index]
			stringMappings[sSplit[index]] = letter
		}
	}
	return true
}
