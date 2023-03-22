package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head   *Node
	length int
}

func (l *Linkedlist) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l Linkedlist) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("\n")
}

func (l *Linkedlist) deleteWithValue(value int) {

	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {

	mylist := Linkedlist{}
	node1 := &Node{data: 48}
	mylist.prepend(node1)
	node2 := &Node{data: 18}
	mylist.prepend(node2)
	node3 := &Node{data: 16}
	mylist.prepend(node3)
	node4 := &Node{data: 11}
	mylist.prepend(node4)
	node5 := &Node{data: 7}
	mylist.prepend(node5)
	node6 := &Node{data: 2}
	mylist.prepend(node6)

	mylist.printListData()
	// fmt.Println(mylist)
	// USE CASE: Has a valid value inside LinkedList
	mylist.deleteWithValue(16)
	mylist.printListData()
	// USE CASE: Doesnt has a valid value inside LinkedList
	mylist.deleteWithValue(100)
	mylist.printListData()
	// USE CASE: Delete the first value of an LinkedList
	mylist.deleteWithValue(2)
	mylist.printListData()
	// USE CASE: Delete an empty LinkedList
	emptyList := Linkedlist{}
	emptyList.deleteWithValue(10)
}
