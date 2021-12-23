package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	cur := head
	if cur == nil {
		return head
	}
	for cur != nil && cur.Next != nil {
		nextVal := cur.Next.Val
		cur.Next.Val = cur.Val
		cur.Val = nextVal

		cur = cur.Next.Next
	}
	return head
}
