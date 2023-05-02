// https://leetcode.com/problems/search-in-rotated-sorted-array/
package leetcode


func binarySearchIndex(n []int, t int) int {
	l, r := 0, len(n)
	for l < r {
		m := (l + r) / 2
		if n[m] == t {
			return m
		}
		if t > n[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}


// v1
func Searchv1(nums []int, target int) int {
	ans, l, r := -1, 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[l] < nums[m] {
			ans = binarySearchIndex(nums[l:m+1], target)
			if ans != -1 {
				return l + ans
			} else {
				l = m + 1
			}
		} else {
			ans = binarySearchIndex(nums[m:r], target)
			if ans != -1 {
				return m + ans
			} else {
				r = m
			}
		}
	}
	return ans
}

// v2 better performance
func Searchv2(nums []int, target int) int {
    ans, l, r := -1, 0, len(nums)
    if nums[l] < nums[r-1] {
        return binarySearchIndex(nums, target)
    } else {
        for l < r {
            m := (l+r)/2
            if nums[l] < nums[m] {
                ans = binarySearchIndex(nums[l:m+1], target)
                if ans != -1 {
                    return l + ans
                } else {
                    l = m + 1
                }
            } else {
                ans = binarySearchIndex(nums[m:r], target)
                if ans != -1 {
                    return m + ans
                } else {
                    r = m
                }
            }
        }
    }
    return ans
}