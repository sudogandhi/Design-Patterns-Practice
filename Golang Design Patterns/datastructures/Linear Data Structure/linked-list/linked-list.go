package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) AddToHead(value int) {
	node := &Node{value, nil}
	if l.Head != nil {
		node.Next = l.Head
	}
	l.Head = node
}

func (l *LinkedList) AddToTail(value int) {
	node := &Node{value, nil}
	temp := l.Head

	if temp == nil {
		l.Head = node
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
	}
}

func (l *LinkedList) IterateList() {
	temp := l.Head
	for temp != nil {
		fmt.Printf("%d --> ", temp.Value)
		temp = temp.Next
	}
	fmt.Println()
}

func (l *LinkedList) LastNode() *Node {
	temp := l.Head
	// Traverse the list to find the last node
	for temp.Next != nil {
		temp = temp.Next
	}
	return temp
}

func (l *LinkedList) Search(value int) *Node {
	temp := l.Head
	// Traverse the list to find the node with the given value
	for temp != nil {
		if temp.Value == value {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (l *LinkedList) InsertAtPosition(value int, position int) {
	if position < 0 {
		panic("position cannot be negative")
	}
	// Check if the position is valid
	if position == 0 {
		l.AddToHead(value)
	} else {
		temp := l.Head
		i := 1

		// Traverse the list to find the node at the given position
		for temp != nil && i < position {
			temp = temp.Next
			i++
		}

		if temp == nil {
			panic("position of elements does not exist")
		} else {
			// Create a new node with the given value
			newNode := &Node{Value: value}
			newNode.Next = temp.Next
			temp.Next = newNode
		}
	}
}

func (l *LinkedList) InsertAfter(searchVal int, insertVal int) {
	node := l.Search(searchVal)

	if node == nil {
		fmt.Println("Node with value ", searchVal, " not found in the list")
	} else {
		newNode := &Node{Value: insertVal}
		newNode.Next = node.Next
		node.Next = newNode
	}
}

func main() {
	linked_list := &LinkedList{nil}
	linked_list.AddToTail(1)
	linked_list.AddToTail(2)
	linked_list.AddToTail(3)
	linked_list.AddToTail(4)
	linked_list.AddToTail(5)
	linked_list.AddToTail(6)
	linked_list.AddToTail(7)
	linked_list.AddToTail(8)
	linked_list.AddToTail(9)

	linked_list.IterateList()

	linked_list.AddToHead(10)
	linked_list.AddToHead(11)
	linked_list.AddToHead(12)
	linked_list.AddToHead(13)
	linked_list.AddToHead(14)
	linked_list.AddToHead(15)

	linked_list.IterateList()

	fmt.Println("Last Node = ", linked_list.LastNode().Value)

	linked_list.InsertAfter(14, 16)
	linked_list.InsertAtPosition(17, 4)

	linked_list.IterateList()

}
