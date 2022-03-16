package list

import (
	"errors"
	"fmt"
)

type LinkeListNode struct {
	Item int
	Next *LinkeListNode
}

type LinkeList struct {
	Head *LinkeListNode
	Tail *LinkeListNode
	Size int
}

func (list *LinkeList) New() {
	list.Head = nil
	list.Tail = nil
	list.Size = 0
}

func (list *LinkeList) Insert(element int) {
	node := new(LinkeListNode)
	node.Item = element
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		list.Size += 1
		return
	}
	list.Tail.Next = node
	list.Tail = node
	list.Size += 1
}

func (list *LinkeList) Delete(element int) {
	pre := list.Head
	node := list.Head
	for node != nil && node.Item != element {
		pre = node
		node = node.Next
	}
	if node == nil {
		fmt.Println("cannot delete, element is not existd.")
		return
	}
	if node == list.Head {
		list.Head = node.Next
		node = nil
		list.Size -= 1
		return
	}
	// 下移动指针
	pre.Next = node.Next
	list.Size -= 1
}

func (list *LinkeList) Get(index int) (element int, err error) {
	if index < 0 || index > list.Size-1 {
		return -1, errors.New("out of index")
	}
	node := list.Head
	for idx := 0; index != idx; idx++ {
		if node != nil {
			return node.Item, nil
		}

	}
	return -1, errors.New("not found")
}
