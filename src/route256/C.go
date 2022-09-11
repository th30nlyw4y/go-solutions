package route256

import "fmt"

func C() {
	var n int
	var line, currCode, ans string
	codes := map[string]rune{
		"00":  'a',
		"100": 'b',
		"101": 'c',
		"11":  'd',
	}
	_, _ = fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		currCode, ans = "", ""
		_, _ = fmt.Scanln(&line)
		for _, r := range line {
			currCode += string(r)
			if c, ok := codes[currCode]; ok {
				currCode = ""
				ans += string(c)
			}
		}
		fmt.Println(ans)
	}
}
