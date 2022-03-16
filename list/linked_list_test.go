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
	t.SkipNow()
}
