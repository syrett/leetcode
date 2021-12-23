package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}
	start, end := head, tmp
	i := 0
	for end.Next != nil {
		end = end.Next
		if i > n {
			start = start.Next
		}
		i++
	}
	start.Next = start.Next.Next
	return tmp.Next
}
