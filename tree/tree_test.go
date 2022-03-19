package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(8)

	tree.Delete(3)
	tree.Delete(2)
	tree.Delete(6)

	preorder(tree.root)
	t.Log("====preorder====")

	inorder(tree.root)
	t.Log("====inorder====")

	postorder(tree.root)
	t.Log("====postorder====")
}

func preorder(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%v,", node.Value)
	preorder(node.left)
	preorder(node.right)
}

func inorder(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Printf("%v,", node.Value)
	inorder(node.right)
}

func postorder(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Printf("%v,", node.Value)
}
