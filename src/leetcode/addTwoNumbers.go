package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// v1
func add(l1 *ListNode, l2 *ListNode, pRem int) *ListNode {
	s := l1.Val + l2.Val + pRem
	cRem, ans := 0, &ListNode{}
	if s >= 10 {
		s = s % 10
		cRem = 1
	}
	if l1.Next != nil && l2.Next != nil {
		ans = &ListNode{Val: s, Next: add(l1.Next, l2.Next, cRem)}
	} else if l1.Next == nil && l2.Next == nil {
		if cRem != 0 {
			ans = &ListNode{Val: s, Next: &ListNode{Val: 1, Next: nil}}
		} else {
			ans = &ListNode{Val: s, Next: nil}
		}
	} else if l1.Next == nil {
		ans = &ListNode{Val: s, Next: add(&ListNode{Val: 0, Next: nil}, l2.Next, cRem)}
	} else if l2.Next == nil {
		ans = &ListNode{Val: s, Next: add(l1.Next, &ListNode{Val: 0, Next: nil}, cRem)}
	}
	return ans
}

func AddTwoNumbersv2(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

// v2: leetcode solution (performance is the same, a bit less memory (~400 KB))
func AddTwoNumbersv1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: nil}
	curr, carry := dummyHead, 0
	for l1 != nil || l2 != nil || carry != 0 {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
		} else {
			l1Val = 0
		}
		if l2 != nil {
			l2Val = l2.Val
		} else {
			l2Val = 0
		}
		s := l1Val + l2Val + carry
		carry = s / 10
		newNode := &ListNode{Val: s % 10, Next: nil}
		curr.Next = newNode
		curr = newNode
		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}
	return dummyHead.Next
}
