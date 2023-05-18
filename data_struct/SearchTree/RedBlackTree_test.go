package SearchTree

import (
	"math/rand"
	"testing"
)

func TestRedBlackTree_Insert(t *testing.T) {
	RBTree := NewRedBlackTree()
	//values := []int{10, 5, 15, 3, 7, 13, 18, 1, 0, 20}
	//
	//for _, val := range values {
	//	RBTree.Add(val)
	//}
	for i := 0; i < 1e6; i++ {
		RBTree.Add(rand.Int())
	}
	// 检查插入后根节点是否为黑色
	if RBTree.root.color != BLACK {
		t.Errorf("root node's color is not black")
	}
	// 验证红黑树的结构是否满足要求
	if !RBTree.validateRBTree(RBTree.root) {
		t.Errorf("red-black tree validation failed")
	}
}

// 验证红黑树是否符合标准：
// 1. 根节点为黑色
// 2. 每个节点要么是红色，要么是黑色
// 3. 红色节点的子节点都是黑色
// 4. 从任意节点到其子树叶节点的所有路径都包含相同数目的黑色节点
func (T *RedBlackTree) validateRBTree(node *RBNode) bool {
	if node == T.nil { // 空节点为黑色
		return true
	}
	if node.color == RED {
		if node.left.color == RED || node.right.color == RED {
			return false
		}
	}
	leftBlack := T.getBlackDepth(node.left)
	rightBlack := T.getBlackDepth(node.right)
	if leftBlack != rightBlack {
		return false
	}
	return T.validateRBTree(node.left) && T.validateRBTree(node.right)
}

// 计算子树中黑色节点的深度
func (T *RedBlackTree) getBlackDepth(node *RBNode) int {
	if node == T.nil {
		return 0
	}
	leftBlack := T.getBlackDepth(node.left)
	rightBlack := T.getBlackDepth(node.right)
	res := 0
	if leftBlack > rightBlack {
		res = leftBlack
	} else {
		res = rightBlack
	}
	if node.color == BLACK {
		res++
	}
	return res
}
