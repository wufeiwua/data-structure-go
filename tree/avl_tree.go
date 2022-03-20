package tree

type AVLTreeNode struct {
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree() *AVLTree {

	tree := &AVLTree{
		root: nil,
	}
	return tree
}

func (tree *AVLTree) Insert(value int) *AVLTreeNode {
	node := tree.root
	tree.root = insertAVLNode(value, node)
	return tree.root
}

func insertAVLNode(value int, node *AVLTreeNode) *AVLTreeNode {
	return nil
}

func getBalance(node *AVLTreeNode) int {
	return getHeight(node.left) - getHeight(node.right)
}

func getHeight(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
