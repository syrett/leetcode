package main

type Node struct {
	Val  int
	Next *Node
}
type MyLinkedList struct {
	root *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{nil}
}

func (this *MyLinkedList) Get(index int) int {
	if this.root == nil {
		return -1
	}

	for tmp := this.root; tmp != nil && index >= 0; index-- {
		if index == 0 {
			return tmp.Val
		}

		tmp = tmp.Next
	}

	return -1

}

func (this *MyLinkedList) AddAtHead(val int) {
	this.root = &Node{val, this.root}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.root == nil {
		this.root = &Node{val, nil}
	} else {
		for tmp := this.root; tmp != nil; {
			if tmp.Next == nil {
				tmp.Next = &Node{val, nil}
				return
			}
			tmp = tmp.Next
		}
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.root = &Node{val, this.root}
		return
	}

	for tmp := this.root; tmp != nil && index >= 1; index-- {
		if index == 1 {
			tmp.Next = &Node{val, tmp.Next}
			break
		} else {
			tmp = tmp.Next
		}
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.root = this.root.Next
		return
	}

	for tmp := this.root; tmp.Next != nil && index >= 0; index-- {
		if index == 1 {
			tmp.Next = tmp.Next.Next
			break
		} else {
			tmp = tmp.Next
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
