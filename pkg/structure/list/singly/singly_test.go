package singly

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewNode(t *testing.T) {
	num := rand.Int()
	newNode := NewNode(num)
	assert.Equal(t, newNode.value, num, "should be equal")
	assert.Nil(t, newNode.next, "next should be nil")
}

func TestNewLinkedList(t *testing.T) {
	newList := NewLinkedList()
	assert.Nil(t, newList.head, "head should be nil")
	assert.Nil(t, newList.tail, "tail should be nil")
	assert.Zero(t, newList.length, "length should be 0")
}

func TestAppend(t *testing.T) {
	newList := NewLinkedList()
	t.Run("empty list", func(t *testing.T){
		newList.Append(1)
		assert.Equal(t, newList.head.value, 1, "head.value should be equal")
		assert.Equal(t, newList.tail.value, 1, "tail.value should be equal")
		assert.Equal(t, newList.length, 1, "length should be equal")
	})
	t.Run("with one node", func(t *testing.T){
		newList.Append(2)
		assert.Equal(t, newList.head.value, 1, "head.value should be equal")
		assert.Equal(t, newList.head.next.value, 2, "tail.next.value should be equal")
		assert.Equal(t, newList.tail.value, 2, "tail.value should be equal")
		assert.Equal(t, newList.length, 2, "length should be equal")
	})
}
