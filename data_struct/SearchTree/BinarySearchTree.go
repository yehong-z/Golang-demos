package SearchTree

import "fmt"

type BinaryTree struct {
	root *Node
	SearchTree
}

type Node struct {
	key   int
	left  *Node
	right *Node
	p     *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{root: nil}
}

func (T *BinaryTree) Search(val int) bool {
	// T.InorderTreeWalk(T.root)
	// fmt.Println()
	if node := T.SearchNode(T.root, val); node != nil {
		return true
	} else {
		return false
	}
}

func (T *BinaryTree) Add(val int) {
	T.Insert(&Node{key: val})
}

func (T *BinaryTree) Erase(val int) bool {
	if node := T.SearchNode(T.root, val); node != nil {
		T.Delete(node)
		return true
	} else {
		return false
	}
}

// InorderTreeWalk 中序遍历
func (T *BinaryTree) InorderTreeWalk(node *Node) {
	if node == nil {
		return
	}
	T.InorderTreeWalk(node.left)
	fmt.Printf("%v ", node.key)
	T.InorderTreeWalk(node.right)
}

func (T *BinaryTree) SearchNode(node *Node, key int) *Node {
	for node != nil && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func (T *BinaryTree) Min(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
}

func (T *BinaryTree) Max(node *Node) *Node {
	for node.right != nil {
		node = node.right
	}
	return node
}

func (T *BinaryTree) Successor(node *Node) *Node {
	if node.right == nil {
		next := node.p
		for next.right == node {
			node = next
			next = next.p
		}
		return next
	} else {
		return T.Min(node.right)
	}
}

func (T *BinaryTree) Insert(node *Node) {
	var y *Node = nil
	x := T.root
	for x != nil {
		y = x
		if node.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.p = y
	if y == nil {
		T.root = node
	} else if node.key < y.key {
		y.left = node
	} else {
		y.right = node
	}
}

func (T *BinaryTree) Delete(z *Node) {
	if z.left == nil {
		T.transplant(z, z.right)
	} else if z.right == nil {
		T.transplant(z, z.left)
	} else {
		y := T.Min(z.right)
		if y.p != z {
			T.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		T.transplant(z, y)
		y.left = z.left
		y.left.p = y
	}
}

func (T *BinaryTree) transplant(u, v *Node) {
	if u.p == nil {
		T.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	if v != nil {
		v.p = u.p
	}
}
