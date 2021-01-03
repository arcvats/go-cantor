package doubly

// Node of doubly linked list
type Node struct {
	next *Node
	previous *Node
	value interface{}
}

// LinkedList doubly
type LinkedList struct {
	head *Node
	tail *Node
	length int
}

// NewNode constructor
func NewNode(value interface{}) *Node {
	return &Node{value: value, next: nil, previous: nil}
}

// NewLinkedList constructor
func NewLinkedList() LinkedList {
	return LinkedList{head: nil, tail: nil, length: 0}
}

// Append adds node at the end of linked list
func (list *LinkedList) Append(value interface{}) {
	var newNode *Node = NewNode(value)
	if (list.head == nil) && (list.tail == nil) {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.length++
}

// Prepend adds node at the beginning of the linked list
func (list *LinkedList) Prepend(value interface{}) {
	var newNode *Node = NewNode(value)
	if (list.head == nil) && (list.tail == nil) {
		list.head = newNode
		list.tail = newNode
	} else {
		list.head.previous = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.length++
}

// DeleteFirst deletes node at the beginning of the linked list
func (list *LinkedList) DeleteFirst() interface{} {
	if list.length <= 0 {
		return nil
	}
	var val interface{} = list.head.value
	if list.length == 1 {
		list.tail = nil
		list.head = nil
		list.length--
		return val
	}
	list.head = list.head.next
	list.head.previous = nil
	list.length--
	return val
}