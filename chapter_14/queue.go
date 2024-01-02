package chapter_14

// Queue made with doubly linked list

type Queue struct {
	storage DoublyLinkedList
}

func (q *Queue) Enqueue(value interface{}) {
	q.storage.InsertAtEnd(value)
}

func (q *Queue) Dequeue() interface{} {
	return q.storage.RemoveFromFront()
}

func (q *Queue) Next() interface{} {
	if q.storage.GenesisNode == nil {
		return nil
	}
	return q.storage.GenesisNode.Value
}
