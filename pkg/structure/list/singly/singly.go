package singly

import "errors"

// Node of singly linked list
type Node struct {
	next  *Node
	value interface{}
}

// LinkedList singly implementation
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// NewNode constructor
func NewNode(value interface{}) *Node {
	return &Node{value: value, next: nil}
}

//NewLinkedList constructor
func NewLinkedList() LinkedList {
	return LinkedList{head: nil, tail: nil, length: 0}
}

//Append adds node at the end of linked list
func (list *LinkedList) Append(value interface{}) {
	var newNode *Node = NewNode(value)
	if list.head == nil && list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
	list.length++
}

//Prepend adds node at the beginning of the linked list
func (list *LinkedList) Prepend(value interface{}) {
	var newNode *Node = NewNode(value)
	if list.head == nil && list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
	list.length++
}

//InsertAtPosition node is inserted at given position
func (list *LinkedList) InsertAtPosition(position int, value interface{}) error {
	if (position > list.length) || (position < 0) {
		return errors.New("error: position is either greater than length or less than 0")
	}
	if position == 0 {
		list.Prepend(value)
	} else if position == list.length {
		list.Append(value)
	} else {
		var newNode *Node = NewNode(value)
		var current *Node = list.head
		for idx := 0; idx < position-1; idx++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
		list.length++
	}
	return nil
}

//DeleteFirst node of the linked list
func (list *LinkedList) DeleteFirst() interface{} {
	if list.length <= 0 {
		return nil
	}
	if list.length == 1 {
		list.tail = nil
	}
	var val interface{} = list.head.value
	list.head = list.head.next
	list.length--
	return val
}

//DeleteLast node of the linked list
func (list *LinkedList) DeleteLast() interface{} {
	if list.length <= 0 {
		return nil
	}
	if list.length == 1 {
		return list.DeleteFirst()
	}
	var val interface{} = list.tail.value
	var current *Node = list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	list.tail = current
	list.length--
	return val
}

//DeleteAtPosition a node at a position in the linked list
func (list *LinkedList) DeleteAtPosition(position int) (interface{}, error) {
	if (position > list.length) || (position < 0) {
		return nil, errors.New("error: position is either greater than length or less than 0")
	}
	if position == 0 {
		return list.DeleteFirst(), nil
	} else if position == list.length {
		return list.DeleteLast(), nil
	} else {
		var current *Node = list.head
		for idx := 0; idx < position-1; idx++ {
			current = current.next
		}
		var val interface{} = current.next.value
		current.next = current.next.next
		list.length--
		return val, nil
	}
}

//ToArray converts linked list to array
func (list *LinkedList) ToArray() []interface{} {
	var array []interface{}
	var current *Node = list.head
	for current != nil {
		array = append(array, current.value)
		current = current.next
	}
	return array
}

//FromArray creates a linked list from array
func (list *LinkedList) FromArray(array []interface{}) {
	for _, value := range array {
		list.Append(value)
	}
}

//Destroy the linked list
func (list *LinkedList) Destroy() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

/*
1. Append -- DONE
2. Prepend -- DONE
3. insert at position -- DONE
4. delete first -- DONE
5. delete last -- DONE
6. delete at position -- DONE
7. from array -- DONE
8. to array -- DONE
9. Destroy -- DONE
10. search
11. delete by value with flag to support multiple
12. sort
13. reverse
14. print
*/
