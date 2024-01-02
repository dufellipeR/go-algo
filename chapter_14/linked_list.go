package chapter_14

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	GenesisNode *Node
}

func (l *LinkedList) Read(index int) interface{} {
	if index == 0 {
		return l.GenesisNode.Value
	}

	currentNode := l.GenesisNode
	currentIndex := 0

	for currentIndex < index {
		currentNode = currentNode.Next
		currentIndex++

		if currentNode == nil {
			return nil
		}
	}
	return currentNode.Value
}

func (l *LinkedList) IndexOf(value interface{}) interface{} {
	if value == l.GenesisNode.Value {
		return 0
	}

	currentNode := l.GenesisNode
	currentIndex := 0

	for currentNode.Value != value {
		if currentNode.Next == nil {
			return nil
		}

		currentNode = currentNode.Next
		currentIndex++

	}

	return currentIndex
}

func (l *LinkedList) InsertAtIndex(index int, value interface{}) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}

	if index == 0 {
		newNode.Next = l.GenesisNode
		l.GenesisNode = &newNode
		return
	}

	currentNode := l.GenesisNode
	currentIndex := 0

	for currentIndex < index-1 {
		currentNode = currentNode.Next
		currentIndex++
	}

	newNode.Next = currentNode.Next
	currentNode.Next = &newNode
}

func (l *LinkedList) DeleteAtIndex(index int) {

	if index == 0 {
		l.GenesisNode = l.GenesisNode.Next
	}

	currentNode := l.GenesisNode
	currentIndex := 0

	for currentIndex < index-1 {
		currentNode = currentNode.Next
		currentIndex++
	}

	currentNode.Next = currentNode.Next.Next
}

func (l *LinkedList) Show() {
	currentNode := l.GenesisNode

	for {
		if currentNode == nil {
			break
		}
		fmt.Println(currentNode)
		currentNode = currentNode.Next

	}
}

func (l *LinkedList) ReadLast() interface{} {
	currentNode := l.GenesisNode

	for {
		if currentNode.Next == nil {
			return currentNode
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) Reverse() {
	currentNode := l.GenesisNode
	var stack Stack
	var node *Node
	index := 0

	for {
		stack.Push(currentNode)
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}

	for {
		node = stack.Pop()
		if index == 0 {
			l.GenesisNode = node
		}
		node.Next = stack.Read()
		index++
		if stack.Read() == nil {
			break
		}
	}

}
