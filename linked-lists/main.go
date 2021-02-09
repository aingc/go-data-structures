package main

import "fmt"

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	length int
}

// Receiver is a pointer which means we want to make modifications to the receiver
// If its a value receiver then we'd just be working with a copy of it
func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// append at end of linked list
func (l *linkedList) append(n *node) {
	// if list is empty
	if l.length == 0 {
		l.head = n
		l.head.next = nil
		l.length++
		return
	}

	// If the list only has one node
	if l.head.next == nil {
		l.head.next = n
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
	l.length++
	return
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.value)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	// edge cases
	// if list is empty
	if l.length == 0 {
		return
	}

	// check if head has value
	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	// compare value in the next node so we don't lose our place when deleting
	for previousToDelete.next.value != value {
		// if value does not exist in the list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{value: 48}
	node2 := &node{value: 18}
	node3 := &node{value: 16}
	node4 := &node{value: 11}
	node5 := &node{value: 7}
	node6 := &node{value: 2}
	node7 := &node{value: 1337}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteWithValue(100)
	mylist.deleteWithValue(2)
	mylist.printListData()
	emptyList := linkedList{}
	emptyList.printListData()
	emptyList.deleteWithValue(10)
	emptyList.append(node7)
	emptyList.printListData()
	emptyList.append(node6)
	emptyList.printListData()
	mylist.append(node7)
	mylist.printListData()
}