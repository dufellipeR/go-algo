package chapter_14

import "fmt"

type DoublyNode struct {
	Value interface{}
	Prev  *DoublyNode
	Next  *DoublyNode
}

type DoublyLinkedList struct {
	GenesisNode *DoublyNode
	LastNode    *DoublyNode
}

func (d *DoublyLinkedList) InsertAtEnd(value interface{}) {
	newNode := DoublyNode{
		Value: value,
		Prev:  d.LastNode,
		Next:  nil,
	}

	if d.GenesisNode == nil {
		d.GenesisNode = &newNode
	}

	d.LastNode.Next = &newNode
	d.LastNode = &newNode
}

func (d *DoublyLinkedList) RemoveFromFront() interface{} {
	removedNode := d.GenesisNode

	if d.GenesisNode.Next != nil {
		d.GenesisNode = d.GenesisNode.Next
	} else {
		d.GenesisNode = nil
	}

	return removedNode
}

func (d *DoublyLinkedList) ShowReverse() {
	currentNode := d.LastNode
	for {
		if currentNode == nil {
			break
		}
		fmt.Println(currentNode)
		currentNode = currentNode.Prev
	}
}
