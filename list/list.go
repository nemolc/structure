package list

type List struct {
	Head *ListNode
}

type ListNode struct {
	Data interface{}
	last *ListNode
	next *ListNode
}

