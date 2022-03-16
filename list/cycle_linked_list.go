package list

/* 循环双向链表，无表头 */
import (
	"errors"
	"fmt"
)

type CycleLinkedListNode struct {
	Prev *CycleLinkedListNode
	Next *CycleLinkedListNode
	Item int
}

type CycleLinkedList struct {
	Head *CycleLinkedListNode
	Size int
}

func (list *CycleLinkedList) New() {
	list.Head = nil
	list.Size = 0
}

func (list *CycleLinkedList) Insert(element int) {
	node := &CycleLinkedListNode{Item: element, Next: list.Head, Prev: nil}
	if list.Head == nil {
		list.Head = node
		node.Prev = list.Head
		node.Next = list.Head
	}
	// set old tail next
	list.Head.Prev.Next = node
	// node prev pointer to old tail
	node.Prev = list.Head.Prev

	// head
	node.Next = list.Head

	// set new tail
	list.Head.Prev = node

	list.Size += 1
}

func (list *CycleLinkedList) Delete(element int) {
	node := list.Head
	for node != nil && node.Item != element {
		fmt.Printf("%p---%p----%p----%p\n", list.Head, list.Head.Prev, node, node.Next)
		node = node.Next
		if node == list.Head {
			fmt.Println("cannot delete, element is not existd.")
			return
		}
	}
	if node == list.Head {
		// ** 如果删除的是头节点 ** 重新设置头节点
		list.Head = list.Head.Next
	}
	if list.Size == 1 {
		list.Head = nil
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node = nil

}

func (list *CycleLinkedList) Get(index int) (element int, err error) {
	if index < 0 || index > list.Size-1 {
		return -1, errors.New("out of index")
	}
	node := list.Head
	for idx := 0; index != idx && idx < list.Size; idx++ {
		if node != nil {
			return node.Item, nil
		}

	}
	return -1, errors.New("not found")
}
