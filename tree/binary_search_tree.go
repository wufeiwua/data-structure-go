package tree

type BinarySearchTree struct {
	*BinaryTree
}

func NewBinarySearchTree() *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.BinaryTree = NewBinaryTree()
	tree.root = nil
	return tree
}

func (tree *BinarySearchTree) Insert(value int) {
	tree.root = insertNode(value, tree.root)
}

func (tree *BinarySearchTree) Delete(value int) {
	tree.root = deleteNode(value, tree.root)
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

func deleteNode(value int, node *BinaryTreeNode) *BinaryTreeNode {
	if value < node.Value {
		node.left = deleteNode(value, node.left)
	} else if value > node.Value {
		node.right = deleteNode(value, node.right)
	} else {
		// 有两个子节点
		if node.left != nil && node.right != nil {
			min := findMin(node.right)
			node.Value = min.Value
			node.right = deleteNode(min.Value, node.right)
		} else if node.left == nil && node.right != nil {
			return node.right
		} else if node.left != nil && node.right == nil {
			return node.left
		} else {
			return nil
		}
	}
	return node
}

func findMin(node *BinaryTreeNode) *BinaryTreeNode {
	min := node
	for min.left != nil {
		min = node.left
	}
	return min
}
