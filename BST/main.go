package main

import "fmt"

// Our nodes of the tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

// Tree holding nodes
// Less than goes left
// greater than goes right
type Tree struct {
	name string
	root *Node
}

func createTree(treeName string) *Tree {
	return &Tree{
		name: treeName,
	}
}

func createNode(value int) *Node {
	return &Node{
		val: value,
	}
}

func (tree *Tree) addNode(newNode *Node, currentNode *Node) {
	if tree.root == nil {
		tree.root = newNode
		return
	}

	if currentNode == nil {
		currentNode = newNode
		return
	}

	// Don't enter duplicates
	if currentNode.val == newNode.val {
		return
	}

	if newNode.val > currentNode.val {
		// go right
		if currentNode.right == nil {
			currentNode.right = newNode
		} else {
			tree.addNode(newNode, currentNode.right)
		}
	} else {
		// go left
		if currentNode.left == nil {
			currentNode.left = nil
		} else {
			tree.addNode(newNode, currentNode.left)
		}
	}
}

func (tree *Tree) searchTree(value int, currentNode *Node) bool {

	if currentNode == nil {
		return false
	}

	fmt.Println("=== Checking Node: ", currentNode.val, "===")

	if currentNode.val == value {
		return true
	}

	if value > currentNode.val {
		// go right
		fmt.Println("Going right")
		tree.searchTree(value, currentNode.right)
	} else {
		// go left
		fmt.Println("Going left")
		tree.searchTree(value, currentNode.left)
	}

	return false
}

func main() {

	tree := createTree("Numbers")

	tree.addNode(createNode(80), tree.root)
	tree.addNode(createNode(87), tree.root)
	tree.addNode(createNode(60), tree.root)

	searchFor := 90

	fmt.Println("===== Looking for", searchFor, "=====")

	if tree.searchTree(searchFor, tree.root) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
