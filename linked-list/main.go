package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend(n *node){
	tail := l.head
	l.head = n
	l.head.next = tail
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Print("\n")
}

func (l *linkedList) deleteWithValue(value int){
	// Handling the empty list case
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func main(){
	list := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:99}
	node3 := &node{data:20}
	node4 := &node{data:9}
	node5 := &node{data:60}
	node6 := &node{data:73}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)
	list.prepend(node6)

	list.printListData()
	list.deleteWithValue(9)
	list.printListData()
}