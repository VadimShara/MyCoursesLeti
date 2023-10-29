package LinkedList

import (
	"fmt"
	"strconv"
)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

func NewNode(value int) *Node {
	var node Node
	node.value = value
	return &node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func New(size int) *LinkedList {
	var l LinkedList
	for i := 0; i < size; i++ {
		l.Add(0)
	}
	return &l
}

func NewFromSlice(s []int) *LinkedList {
	var l LinkedList
	for i := 0; i < len(s); i++ {
		node := NewNode(s[i])
		if i == 0 {
			l.head = node
			l.tail = node
			continue
		}
		l.Add(s[i])
	}
	return &l
}

func (list *LinkedList) Add(value int) {
	node := NewNode(value)
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}
	if list.head == list.tail {
		list.tail = node
		list.head.next = list.tail
		return
	}
	list.tail.next = node
	node.prev = list.tail
	list.tail = node
}

func (list *LinkedList) Pop() {
	list.tail.prev.next = nil
	list.tail = list.tail.prev
}

func (list *LinkedList) Size() int {
	var size int
	if list.head == nil {
		return 0
	}
	node := list.head
	for {
		if node == list.tail {
			size++
			break
		}
		node = node.next
		size++
	}
	return size
}

func (list *LinkedList) DeleteAtHead() {
	if list.head == nil {
		return
	}
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}
	node := list.head
	list.head = node.next
	list.head.prev = nil
}

func (list *LinkedList) DeleteAtTail() {
	node := list.tail
	list.tail = node.prev
	list.tail.next = nil
}

func (list *LinkedList) DeleteFrom(pos int) {
	if pos > list.Size() {
		fmt.Println("Error! The specified position is larger than the list size.")
		return
	}
	if pos == 0 {
		list.DeleteAtHead()
		return
	}
	if pos == list.Size()-1 {
		list.DeleteAtTail()
		return
	}
	node := list.head
	for i := 0; i < pos; i++ {
		node = node.next
	}
	node.prev = node.next
	node.next.prev = node.prev
}

func (l *LinkedList) UpdateAt(pos, value int) {
	if pos > l.Size() {
		fmt.Println("Error! The specified position is larger than the list size.")
		return
	}
	node := l.head
	for i := 0; i < pos; i++ {
		node = node.next
	}
	node.value = value
}

func (l *LinkedList) Print() {
	if l.Size() == 0 {
		fmt.Print("[]")
		return
	}
	node := l.head
	fmt.Print("[")
	for i := 0; i < l.Size(); i++ {
		fmt.Print(strconv.Itoa(node.value) + " -> ")
		node = node.next
	}
	fmt.Println("nil]")
}
