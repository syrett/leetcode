package main

import "testing"

func TestRemoveElements(t *testing.T) {
	type Case struct {
		head           *ListNode
		val            int
		expectListNode *ListNode
	}

	var caseList = []Case{
		{&ListNode{1, &ListNode{3, &ListNode{7, &ListNode{4, nil}}}}, 7, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
		{&ListNode{1, &ListNode{3, &ListNode{3, &ListNode{4, nil}}}}, 3, &ListNode{1, &ListNode{4, nil}}},
	}

	for _, item := range caseList {
		result := removeElements(item.head, item.val)
		for result.Next != nil {
			if result.Val != item.expectListNode.Val {
				t.Errorf("result != expect: %d != %d", result.Val, item.expectListNode.Next.Val)
			} else {
				result = result.Next
				item.expectListNode = item.expectListNode.Next
			}
		}
	}
}
