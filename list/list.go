package list

type SingleNode struct {
	next  *SingleNode
	Value interface{}
}

type DoublyNode struct {
	prev  *DoublyNode
	next  *DoublyNode
	Value interface{}
}

// type List interface {
// 	Insert(value interface{})
// 	Delete(value interface{})
// 	Get(index int) (interface{}, error)
// 	Find(value interface{}) (int, error)
// }
