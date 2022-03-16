package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := new(LinkeList)
	list.New()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)
	list.Delete(1)
	list.Delete(7)
	list.Delete(3)
	list.Delete(8)

	node := list.Head
	for node != nil {
		t.Log(node.Item)
		node = node.Next
	}
	t.Log("size is ", list.Size)
	num, err := list.Get(-1)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}
	num, err = list.Get(10)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}

	num, err = list.Get(3)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}

	t.SkipNow()
}

func TestDoublyLinkedList(t *testing.T) {
	list := new(CycleLinkedList)
	list.New()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)
	list.Insert(7)
	list.Delete(1)
	list.Delete(7)
	list.Delete(3)
	list.Delete(8)
	t.Log("size is ", list.Size)

	for node := list.Head; node != nil; {
		t.Log(node.Item)
		node = node.Next
		if node == list.Head {
			break
		}
	}

	num, err := list.Get(-1)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}
	num, err = list.Get(10)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}

	num, err = list.Get(3)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Log(num)
	}

	t.Log("print reverse")
	for node := list.Head.Prev; node != nil; {
		t.Log(node.Item)
		node = node.Prev
		if node == list.Head.Prev {
			break
		}
	}

	t.SkipNow()
}