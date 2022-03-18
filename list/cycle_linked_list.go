package list

import "errors"

type CycleLinkedList struct {
	head *DoublyNode
	Size int
}

func NewCycleLinkedList() *CycleLinkedList {
	return &CycleLinkedList{
		head: nil,
		Size: 0,
	}
}

func (list *CycleLinkedList) Insert(value interface{}) {
	node := &DoublyNode{
		prev:  nil,
		next:  nil,
		Value: value,
	}
	if list.head == nil {
		list.head = node
		list.head.prev = node
		list.head.next = node
	} else {
		// head prev is the last node
		list.head.prev.next = node
		node.prev = list.head.prev
		node.next = list.head
		// reset the last node
		list.head.prev = node
	}
	list.Size++
}

func (list *CycleLinkedList) Delete(value interface{}) {
	node, idx := list.head, 0
	for node != nil && node.Value != value && idx < list.Size {
		node = node.next
		idx++
	}
	if node == nil {
		// cannot find
		// do nothing
		return
	} else if list.head == node {
		// the node is first
		// reset head
		list.head = list.head.next
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	list.Size--
}

func (list *CycleLinkedList) Get(index int) (interface{}, error) {
	node, idx := list.head, 0
	if index < 0 || index > list.Size-1 {
		return nil, errors.New("out of index")
	}
	for node != nil && idx != index {
		idx++
		node = node.next
	}
	if node == nil {
		return nil, errors.New("node is nil")
	}
	return node.Value, nil
}

func (list *CycleLinkedList) Find(value interface{}) (int, error) {
	node, idx := list.head, 0
	for node != nil && node.Value != value {
		idx++
		node = node.next

		// on circle
		if idx >= list.Size {
			node = nil
			break
		}
	}
	if node == nil {
		return -1, errors.New("cannot find it")
	}
	return idx, nil
}
