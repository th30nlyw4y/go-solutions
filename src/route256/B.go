package route256

import "fmt"

func B() {
	var n, f_len int
	long, short := map[int]bool{1: true, 3: true, 5: true, 7: true, 8: true, 10: true, 12: true}, map[int]bool{4: true, 6: true, 9: true, 11: true}
	_, _ = fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var d, m, y int
		_, _ = fmt.Scanln(&d, &m, &y)
		if y < 1950 || y > 2300 || m < 1 || m > 12 || d < 1 || d > 31 {
			fmt.Println("NO")
			continue
		}
		if _, ok := long[m]; ok {
			fmt.Println("YES")
		} else if _, ok := short[m]; ok {
			if d > 30 {
				fmt.Println("NO")
				continue
			}
			fmt.Println("YES")
		} else {
			if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
				f_len = 29
			} else {
				f_len = 28
			}
			if d > f_len {
				fmt.Println("NO")
				continue
			}
			fmt.Println("YES")
		}
	}
}
