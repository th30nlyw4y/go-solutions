package route256

import (
	"fmt"
)

type User struct {
	id    int
	lastR int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func D() {
	var n, l, r, ans, nReq, rId int
	var currUsers []User
	_, _ = fmt.Scanln(&n)
	for nCase := 0; nCase < n; nCase++ {
		l, r, ans, currUsers = 0, 0, 0, []User{}
		_, _ = fmt.Scanln(&nReq)
		for r < nReq {
			_, _ = fmt.Scan(&rId)
			switch len(currUsers) {
			case 0:
				currUsers = append(currUsers, User{id: rId, lastR: r})
			case 1:
				if currUsers[0].id == rId {
					currUsers[0].lastR = r
				} else {
					currUsers = append(currUsers, User{id: rId, lastR: r})
				}
			case 2:
				if currUsers[0].id == rId {
					currUsers[0].lastR = r
				} else if currUsers[1].id == rId {
					currUsers[1].lastR = r
				} else {
					ans = max(r-l, ans)
					if currUsers[0].lastR > currUsers[1].lastR {
						l = max(l, currUsers[1].lastR+1)
						currUsers[1], currUsers[0] = currUsers[0], User{id: rId, lastR: r}
					} else {
						l = max(l, currUsers[0].lastR+1)
						currUsers[0], currUsers[0] = currUsers[1], User{id: rId, lastR: r}
					}
				}
			}
			r++
		}
		if ans == 0 {
			ans = nReq
		}
		fmt.Println(max(ans, r-l))
		fmt.Scanln()
	}
}
