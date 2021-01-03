package singly

import (
	"math/rand"
	"testing"
)

func TestNewNode(t *testing.T) {
	var num int = rand.Int()
	var newNode *Node = NewNode(num)
	if newNode.value != num {
		t.Errorf("Expected NewNode(%v) value to equal %v", num, newNode.value)
	}
	if newNode.next != nil {
		t.Errorf("Expected NewNode(%v) next to equal %v", num, newNode.next)
	}
}

func TestNewLinkedList(t *testing.T) {
	var newList LinkedList = NewLinkedList()
	if newList.head != nil {
		t.Errorf("Expected NewLinkedList() head nil to equal %v", newList.head)
	}
	if newList.tail != nil {
		t.Errorf("Expected NewLinkedList() tail nil to equal %v", newList.tail)
	}
	if newList.length != 0 {
		t.Errorf("Expected NewLinkedList() length 0 to equal %v", newList.length)
	}
}
