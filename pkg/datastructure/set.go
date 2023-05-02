package datastructure

func SetIntersectCount(x, y map[int]struct{}) int {
	ans := 0
	for k, _ := range x {
		if _, ok := y[k]; ok {
			ans++
		}
	}
	return ans
}
