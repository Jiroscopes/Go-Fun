package main

import "fmt"

// Node only contains strings as val
type Node struct {
	val  string
	prev *Node
	next *Node
}

// List holds Nodes
type List struct {
	name    string
	head    *Node
	tail    *Node
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
		list.tail = newNode
		// fmt.Println("CURRENT HEAD:", list.head.val)
	} else {

		// Otherwise loop until you find the end and add it there
		currentNode := list.tail
		currentNode.next = newNode
		newNode.prev = currentNode
		list.tail = newNode
	}

	return nil
}

func (list *List) printListForward() {
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

func (list *List) printListBackwards() {
	currentNode := list.tail

	// fmt.Println(currentNode.next.val)
	if currentNode.prev == nil {
		fmt.Println(currentNode.val)
	} else {
		for currentNode.prev != nil {
			fmt.Println(currentNode.val)
			currentNode = currentNode.prev
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

		afterNode := currentNode.next.next

		if afterNode != nil {
			afterNode.prev = currentNode
			currentNode.next = afterNode
		} else {
			list.tail = currentNode
			currentNode.next = nil
		}

	}
}

func main() {

	// Create list
	list := createList("Playlist")
	// Add songs
	list.addNode("South of Heaven")
	list.addNode("Ultralight beam")
	list.addNode("Saint Pablo")
	// list.addNode("MIDDLE CHILD")
	list.addNode("RICKY")
	// list.addNode("ELEMENT.")

	// Print
	list.printListBackwards()
	// list.printListForward()
	list.removeNode("RICKY")

	fmt.Println("======NEW PLAYLIST======")

	list.printListBackwards()

}
