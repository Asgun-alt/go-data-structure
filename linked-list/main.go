package main


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

func main(){
	list := linkedList{}
	node1 := &node{data:48}
	node2 := &node{data:99}
	node3 := &node{data:12}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)

	fmt.Println(list)
}