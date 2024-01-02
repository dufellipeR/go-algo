package main

import (
	"algorithms/chapter_12"
	. "algorithms/chapter_14"
	"algorithms/chapter_8"
	"fmt"
)

func main() {
	array1 := []int{0, 2, 3, 45, 42, 3, 21, 256}
	array2 := []int{42, 21, 256}

	duplicated := []string{"a", "b", "c", "d", "c", "e", "f"}

	node1 := Node{Value: 1, Next: nil}
	node2 := Node{
		Value: 2,
		Next:  nil,
	}
	node3 := Node{
		Value: 3,
		Next:  nil,
	}

	node4 := Node{
		Value: 4,
		Next:  nil,
	}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	dNode1 := DoublyNode{Value: 1, Next: nil, Prev: nil}
	dNode2 := DoublyNode{
		Value: 2,
		Next:  nil,
		Prev:  nil,
	}
	dNode3 := DoublyNode{
		Value: 3,
		Next:  nil,
		Prev:  &dNode2,
	}

	dNode1.Next = &dNode2
	dNode2.Next = &dNode3
	dNode2.Prev = &dNode1

	fmt.Println(chapter_8.Intersection(array1, array2))

	fmt.Println(chapter_8.DuplicateValue(duplicated))

	fmt.Println(chapter_8.MissingLetter("the quick brown box jumps over a lazy dog"))

	fmt.Println(chapter_8.NonDuplicateChar("minimum"))

	fmt.Println(chapter_12.AddUntil100(array1))

	fmt.Println(chapter_12.Golomb(5, map[int]int{}))

	fmt.Println(chapter_12.UniquePaths(7, 3, map[string]int{}))

	list := LinkedList{GenesisNode: &node1}

	fmt.Println(list.Read(2))

	fmt.Println(list.IndexOf(3))

	list.InsertAtIndex(1, 21)

	fmt.Println(list.IndexOf(2))

	fmt.Println(list.Read(1))
	fmt.Println(list.Read(2))

	list.DeleteAtIndex(1)

	fmt.Println(list.Read(1))

	list.Show()

	doublyList := DoublyLinkedList{
		GenesisNode: &dNode1,
		LastNode:    &dNode3,
	}

	doublyList.ShowReverse()

	fmt.Println(list.ReadLast())

	list.Show()

}
