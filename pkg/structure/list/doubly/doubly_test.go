package doubly

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
	assert.Nil(t, newNode.previous, "previous should be nil")
}

func TestNewLinkedList(t *testing.T) {
	newList := NewLinkedList()
	assert.Nil(t, newList.head, "head should be nil")
	assert.Nil(t, newList.tail, "tail should be nil")
	assert.Zero(t, newList.length, "length should be 0")
}