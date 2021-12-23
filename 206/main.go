package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	for tmp := head; tmp != nil; {
		result = &ListNode{tmp.Val, result}
		tmp = tmp.Next
	}
	return result
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
