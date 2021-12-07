package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{}
	if l2 == nil {
		if l1.Val > 9 {
			resultNode.Val = 10 - l1.Val
			l1.Next.Val++
		} else {
			resultNode.Val = l1.Val
		}
		resultNode.Next = addTwoNumbers(l1.Next, l2)
		return resultNode
	}

	if l1 == nil && l2 == nil {
		return nil
	}

	isCarry := false
	if l1.Next != nil && l2.Next != nil {
		newVal := l1.Val + l2.Val
		if newVal > 9 {
			newVal = 10 - newVal
			isCarry = true
		}
		resultNode.Val = newVal
		if isCarry {
			l1.Next.Val++
		}
		resultNode.Next = addTwoNumbers(l1.Next, l2.Next)
	} else if l2.Next == nil {
		if isCarry {
			l1.Next.Val++
		}
		resultNode.Next = addTwoNumbers(l1.Next, l2)
	} else if l1.Next == nil {
		if isCarry {
			l2.Next.Val++
		}
		resultNode.Next = addTwoNumbers(l2.Next, l1)
	}

	return resultNode
}
