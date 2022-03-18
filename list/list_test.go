package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	list.Delete(1)
	list.Delete(6)
	list.Delete(3)

	size := list.Size
	t.Log("the list size is ", size)
	for i := 0; i < size; i++ {
		value, err := list.Get(i)
		if err == nil {
			t.Log(value)
		}
	}

	if re, err := list.Find(1); err == nil {
		t.Log(re)
	} else {
		t.Log(err)
	}

	if re, err := list.Find(2); err == nil {
		t.Log(re)
	} else {
		t.Log(err)
	}

	if re, err := list.Find(5); err == nil {
		t.Log(re)
	} else {
		t.Log(err)
	}
}

func TestDoublyLinkedList(t *testing.T) {

}

func TestCycleLinkedList(t *testing.T) {

}
