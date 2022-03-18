package tree

type BinaryTreeNode struct {
	left  *BinaryTreeNode
	right *BinaryTreeNode
	Value int // 使用 int 方便比较大小
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}
