package main

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func (l *LinkedList) AddToHead(val int) {
	newNode := &Node{Val: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Head.Prev = newNode
		newNode.Next = l.Head
		l.Head = newNode
	}
}

func (l *LinkedList) AddToTail(val int) {
	newNode := &Node{Val: val}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
}

func (l *LinkedList) Search(val int) *Node {
	for node := l.Head; node != nil; node = node.Next {
		if node.Val == val {
			return node
		}
	}
	return nil
}

func (l *LinkedList) AddAfter(insertVal int, searchVal int) {
	node := l.Search(searchVal)

	if node == nil {
		panic("Element not found")
	} else {
		newNode := &Node{Val: insertVal}
		newNode.Next = node.Next
		newNode.Prev = node
		if node.Next != nil {
			node.Next.Prev = newNode
		} else {
			l.Tail = newNode
		}
		node.Next = newNode
	}
}

func (l *LinkedList) IterateListForward() {
	fmt.Println("Iterating list forward")
	temp := l.Head
	for temp != nil {
		fmt.Printf("%d --> ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func (l *LinkedList) IterateListBackward() {
	fmt.Println("Iterating list backward")
	temp := l.Tail
	for temp != nil {
		fmt.Printf("%d --> ", temp.Val)
		temp = temp.Prev
	}
	fmt.Println()
}

func main() {
	fmt.Println("Implementing Doubly Linked List")

	list := LinkedList{}

	list.AddToHead(10)
	list.AddToHead(20)
	list.AddToHead(30)
	list.AddToHead(40)
	list.AddToHead(50)

	list.IterateListForward()
	list.IterateListBackward()

	list.AddToTail(60)
	list.AddToTail(70)
	list.AddToTail(80)
	list.AddToTail(90)

	list.IterateListForward()
	list.IterateListBackward()

	list.AddAfter(100, 10)
	list.AddAfter(110, 50)
	list.AddAfter(120, 90)

	list.IterateListForward()
	list.IterateListBackward()

}
