package main

import "fmt"

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewTerminalNode(value int) *Node {
	return &Node{value, nil, nil, nil}
}

func NewNode(value int, left, right *Node) *Node {
	node := &Node{value, left, right, nil}
	left.parent = node
	right.parent = node
	return node
}

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root, root, false}
	for i.Current.left != nil {
		i.Current = i.Current.left
	}
	return i
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func main() {
	//   1
	//  / \
	// 2   3

	// in-order: 2 1 3
	// pre-order: 1 2 3
	// post-order: 2 3 1

	root := NewNode(1, NewTerminalNode(2), NewTerminalNode(3))

	// it := NewInOrderIterator(root)
	// for it.MoveNext() {
	// 	fmt.Printf("%d,", it.Current.Value)
	// }

	tree := NewBinaryTree(root)
	it := tree.InOrder()
	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.Value)
	}

	fmt.Println("\b")
}
