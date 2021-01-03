// package main
package main

import (
	"github.com/go-cantor/pkg/structure/list/doubly"
	"github.com/go-cantor/pkg/structure/list/singly"
)

func main() {
	list := singly.NewLinkedList()
	listD := doubly.NewLinkedList()

	listD.Append(1)
	listD.Append(2)
	listD.Append(3)
	listD.Append(4)
	//list2 := singly.NewLinkedList()
	//list3 := singly.NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	//list.deleteAtPosition(4)
	// list.prepend(0)
	// list.prepend(-1)
	// list.insertAtPosition(3, 4.5)
	// list.insertAtPosition(0, 0)
	// list.insertAtPosition(1, 0.5)
	//arr := list.toArray()
	//arr2 := list2.toArray()
	//list3.fromArray([]interface{}{11, 12, 13, 14})
	// list.fromArray([]interface{}{5, 6, 7})
	//println(arr, arr2)
	// list.destroy()
	// list.deleteFirst()
	// list.deleteLast()
	// list3.deleteFirst()
	// list3.deleteLast()
	// list3.deleteLast()
	//list.prepend(-2)
}
