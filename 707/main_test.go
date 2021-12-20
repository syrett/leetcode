package main

import (
	"fmt"
	"testing"
)

// test
func TestGenerateMatrix(t *testing.T) {
	mLL := Constructor()
	v := mLL

	// v := &MyLinkedList{}
	// mLL := *v
	// fmt.Printf("%+v, %+v\n", mLL, v)
	//	v.root = &Node{1, nil}
	v.AddAtHead(111)
	PrintLinkList("AddAtHead", v)
	v.AddAtTail(222)
	PrintLinkList("AddAtTail", v)
	v.AddAtIndex(0, 333)
	PrintLinkList("AddAtIndex", v)
	v.DeleteAtIndex(0)
	// fmt.Printf("%+v, %+v\n", mLL, v)
	PrintLinkList("DeleteAtIndex", v)
}

func PrintLinkList(funcName string, this MyLinkedList) {
	fmt.Printf("---------%s---------\n", funcName)
	for tmp := this.root; tmp != nil; {
		fmt.Printf("aaaaaaaaa:")
		fmt.Printf("%d\n", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Printf("\n")
}
