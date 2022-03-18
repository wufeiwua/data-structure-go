package list

import "errors"

type linkedListNode struct {
	next  *linkedListNode
	Value interface{}
}

type LinkedList struct {
	head *linkedListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		Size: 0,
	}
}

func (list *LinkedList) findPrevious(value interface{}) *linkedListNode {
	node := list.head
	var prev *linkedListNode = nil
	for node != nil && node.Value != value {
		prev = node
		node = node.next
	}
	return prev
}

func (list *LinkedList) getLast() *linkedListNode {
	node := list.head
	for node != nil && node.next != nil {
		node = node.next
	}
	return node
}

func (list *LinkedList) Insert(value interface{}) {
	node := &linkedListNode{next: nil, Value: value}
	if list.head == nil {
		list.head = node
	} else {
		last := list.getLast()
		last.next = node
	}
	list.Size++
}

func (list *LinkedList) Delete(value interface{}) {
	prev := list.findPrevious(value)
	// the node is the head
	if prev == nil {
		head := list.head
		list.head = head.next
		head = nil
	} else {
		node := prev.next
		if node == nil {
			// the node is the empty, prev node is the last
			// do nothing
			return
		}
		prev.next = node.next
		node = nil
	}
	list.Size--
}

func (list *LinkedList) Get(index int) (interface{}, error) {
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

func (list *LinkedList) Find(value interface{}) (int, error) {
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
