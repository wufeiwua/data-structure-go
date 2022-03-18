package list

import "errors"

type DoublyLinkedListNode struct {
	prev  *DoublyLinkedListNode
	next  *DoublyLinkedListNode
	Value interface{}
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		Size: 0,
	}
}

func (list *DoublyLinkedList) Insert(value interface{}) {
	node := &DoublyLinkedListNode{
		prev:  nil,
		next:  nil,
		Value: value,
	}
	if list.head == nil {
		list.head = node
	} else {
		last := list.last()
		last.next = node
		node.prev = last
	}
	list.Size++
}

func (list *DoublyLinkedList) Delete(value interface{}) {
	node := list.head
	for node != nil && node.Value != value {
		node = node.next
	}
	// cannot find it
	if node == nil {
		return
	} else if node.prev == nil {
		// the node is head, update head list
		list.head = node.next
		list.head.prev = nil
	} else if node.next == nil {
		// the node is last
		node.prev.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

func (list *DoublyLinkedList) Get(index int) (interface{}, error) {
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

func (list *DoublyLinkedList) Find(value interface{}) (int, error) {
	node, idx := list.head, 0
	for node != nil && node.Value != value {
		idx++
		node = node.next
	}
	if node == nil {
		return -1, errors.New("cannot find it")
	}
	return idx, nil
}

func (list *DoublyLinkedList) last() *DoublyLinkedListNode {
	node := list.head
	for node != nil && node.next != nil {
		node = node.next
	}
	return node
}
