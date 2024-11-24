package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key         int
	value       int
	left, right *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	treeNode := &Node{key, value, nil, nil}
	if tree.root == nil {
		tree.root = treeNode
	} else {
		insertTreeNode(tree.root, treeNode)
	}
}

func (tree *BinarySearchTree) InorderTraversal() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inorderTraverse(tree.root)
}

func inorderTraverse(node *Node) {
	if node != nil {
		inorderTraverse(node.left)
		fmt.Printf("%d %d\n", node.key, node.value)
		inorderTraverse(node.right)
	}
}

func (tree *BinarySearchTree) PostorderTraversal() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	postorderTraverse(tree.root)
}

func postorderTraverse(node *Node) {
	if node != nil {
		inorderTraverse(node.left)
		inorderTraverse(node.right)
		fmt.Printf("%d %d\n", node.key, node.value)
	}
}

func (tree *BinarySearchTree) PreorderTraversal() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preorderTraverse(tree.root)
}

func preorderTraverse(node *Node) {
	if node != nil {
		fmt.Printf("%d %d\n", node.key, node.value)
		inorderTraverse(node.left)
		inorderTraverse(node.right)
	}
}

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.root

	if treeNode == nil {
		return nil
	} else {
		for {
			if treeNode.left == nil {
				return &treeNode.key
			}
			treeNode = treeNode.left
		}
	}
}

func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.root

	if treeNode == nil {
		return nil
	} else {
		for {
			if treeNode.left == nil {
				return &treeNode.key
			}
			treeNode = treeNode.left
		}
	}
}

func (tree *BinarySearchTree) SearchNode(n int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.root, n)
}

func searchNode(node *Node, n int) bool {
	if node == nil {
		return false
	} else if node.key == n {
		return true
	} else if node.key > n {
		return searchNode(node.right, n)
	} else {
		return searchNode(node.left, n)
	}
}

func (tree *BinarySearchTree) RemoveNode(n int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.root, n)
}

func removeNode(node *Node, n int) *Node {
	if node == nil {
		return nil
	}

	if n < node.key {
		node.left = removeNode(node.left, n)
		return node
	}

	if n > node.key {
		node.right = removeNode(node.right, n)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	var leftMostRightNode *Node
	leftMostRightNode = node.right

	for {
		// find smallest value on the right side
		if leftMostRightNode != nil && leftMostRightNode.left != nil {
			leftMostRightNode = leftMostRightNode.left
		} else {
			break
		}
	}
	node.key, node.value = leftMostRightNode.key, leftMostRightNode.value
	node.right = removeNode(node.right, node.key)

	return node
}

func insertTreeNode(root *Node, node *Node) {
	if node.key < root.key {
		if root.left == nil {
			root.left = node
		} else {
			insertTreeNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insertTreeNode(root.right, node)
		}
	}
}

func main() {
	fmt.Println("Implementing Binary Search Trees")
}
