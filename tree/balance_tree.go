package tree

import "math"

type BalanceTreeNode struct {
	num   int                  // 元素个数
	Value []int                // 元素
	Node  [](*BalanceTreeNode) // 子节点
}

type BalanceTree struct {
	root *BalanceTreeNode // 根节点
	m    int              // 阶
}

func NewBalanceTree(m int) *BalanceTree {
	return &BalanceTree{
		root: nil,
		m:    m,
	}
}

func (tree *BalanceTree) Insert() *BalanceTreeNode {
	return nil
}

func (tree *BalanceTree) insertValue(value int, node *BalanceTreeNode) *BalanceTreeNode {
	if node == nil {
		node = tree.newNode()
	}
	// 元素个数大于 m -1
	if int(math.Ceil(float64(node.num/2))) > (tree.m - 1) {

	}
	return nil
}

func (tree *BalanceTree) newNode() *BalanceTreeNode {
	valueArr, nodeArr := [...]int{}, [...](*BalanceTreeNode){}
	return &BalanceTreeNode{
		num:   0,
		Value: valueArr[:],
		Node:  nodeArr[:],
	}
}
