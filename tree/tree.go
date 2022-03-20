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

func RRRoate(k2 *BinaryTreeNode) *BinaryTreeNode {
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	return k1
}

func LLRoate(k2 *BinaryTreeNode) *BinaryTreeNode {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	return k1

}

func RLRoate(k3 *BinaryTreeNode) *BinaryTreeNode {
	k3.right = LLRoate(k3.right)
	return RRRoate(k3)
}

func LRRoate(k3 *BinaryTreeNode) *BinaryTreeNode {
	k3.left = RRRoate(k3.left)
	return LLRoate(k3)
}
