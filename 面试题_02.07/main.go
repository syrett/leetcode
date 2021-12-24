package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 很浪漫，走完自己的路，然后跳到对方的路上走一遍， 如果有缘，总会遇上。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB

	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = headA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = headB.Next
		}
	}

	return curA
}
