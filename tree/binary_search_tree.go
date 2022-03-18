package tree

type BinarySearchTree struct {
	*BinaryTree
}

func NewBinarySearchTree() *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.root = nil
	return tree
}

func (tree *BinarySearchTree) Insert(value int) {
	tree.root = insertNode(value, tree.root)
}

func insertNode(value int, node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		node = &BinaryTreeNode{
			left:  nil,
			right: nil,
			Value: value,
		}
	} else if value < node.Value {
		node.left = insertNode(value, node.left)
	} else if value > node.Value {
		node.right = insertNode(value, node.right)
	}
	return node
}
