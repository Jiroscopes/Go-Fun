package main

import "fmt"

// Node only contains strings as val
type Node struct {
	val  string
	next *Node
}

// List holds Nodes
type List struct {
	name    string
	head    *Node
	current *Node
}

func createList(listName string) *List {
	return &List{
		name: listName,
	}
}

func (list *List) addNode(name string) error {
	// Create the new Node
	newNode := &Node{
		val: name,
	}

	// If the list is empty then assign to head
	if list.head == nil {
		list.head = newNode
	} else {

		// Otherwise loop until you find the end and add it there
		currentNode := list.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
	}
	return nil
}

func (list *List) printList() {
	currentNode := list.head

	// If list only has one node
	if currentNode.next == nil {
		fmt.Println(currentNode.val)
	} else {

		// print all nodes until end
		for currentNode.next != nil {
			fmt.Println(currentNode.val)
			currentNode = currentNode.next
		}

		fmt.Println(currentNode.val)
	}
}

func (list *List) removeNode(val string) {
	currentNode := list.head

	if currentNode.next == nil {
		list.head = nil
	} else {
		for currentNode.next.val != val {
			currentNode = currentNode.next
		}

		currentNode.next = currentNode.next.next
	}
}

func main() {

	// Create list
	list := createList("Playlist")
	// Add songs
	list.addNode("South of Heaven")
	list.addNode("Ultralight beam")
	list.addNode("Saint Pablo")
	list.addNode("MIDDLE CHILD")
	list.addNode("RICKY")

	list.addNode("ELEMENT.")
	// Print
	list.printList()

	list.removeNode("RICKY")

	fmt.Println("======NEW PLAYLIST======")

	list.printList()

}
