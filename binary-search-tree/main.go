package main

import "fmt"

var count int

// Node
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not be already
func (n *Node) Insert(k int){
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
			} else {
				n.Right.Insert(k)
			}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	count++
	
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

func main(){
	tree := &Node{Key: 100}
	fmt.Println(tree)

	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(200)
	tree.Insert(230)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(59)
	tree.Insert(90)

	fmt.Println(tree.Search(200))
	fmt.Println(count)
}