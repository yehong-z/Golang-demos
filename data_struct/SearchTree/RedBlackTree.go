package SearchTree

import "fmt"

const (
	RED   = true
	BLACK = false
)

type RedBlackTree struct {
	root *RBNode
	nil  *RBNode
	SearchTree
}

type RBNode struct {
	key   int
	left  *RBNode
	right *RBNode
	p     *RBNode
	color bool
}

func NewRedBlackTree() RedBlackTree {
	RBTree := RedBlackTree{nil: &RBNode{color: BLACK}}
	RBTree.nil.p = RBTree.nil
	RBTree.root = RBTree.nil
	return RBTree
}

func (T *RedBlackTree) Search(val int) bool {
	if node := T.SearchNode(T.root, val); node != T.nil {
		return true
	} else {
		return false
	}
}

func (T *RedBlackTree) Add(val int) {
	T.RBInsert(&RBNode{key: val, color: RED, p: T.nil, left: T.nil, right: T.nil})
}

func (T *RedBlackTree) Erase(val int) bool {
	if node := T.SearchNode(T.root, val); node != T.nil {
		T.RBDelete(node)
		return true
	} else {
		return false
	}
}

func (T *RedBlackTree) travel(node *RBNode) {
	if node == T.nil {
		return
	}
	T.travel(node.left)
	fmt.Printf("%v ", node.key)
	T.travel(node.right)
}

func (T *RedBlackTree) SearchNode(node *RBNode, key int) *RBNode {
	for node != T.nil && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}

func (T *RedBlackTree) LeftRotate(x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != T.nil {
		y.left.p = x
	}
	y.p = x.p
	if x.p == T.nil {
		T.root = y
	} else if x.p.left == x {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func (T *RedBlackTree) RightRotate(y *RBNode) {
	x := y.left
	y.left = x.right
	if x.right != T.nil {
		x.right.p = y
	}
	x.p = y.p
	if y.p == T.nil {
		T.root = x
	} else if y.p.left == y {
		y.p.left = x
	} else {
		y.p.right = x
	}
	x.right = y
	y.p = x
}

func (T *RedBlackTree) RBInsert(z *RBNode) {
	y := T.nil
	x := T.root
	for x != T.nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == T.nil {
		T.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.left = T.nil
	z.right = T.nil
	z.color = RED
	T.RBInsertFixup(z)
}

func (T *RedBlackTree) RBInsertFixup(z *RBNode) {
	for z.p.color == RED {
		if z.p == z.p.p.left {
			y := z.p.p.right
			// 情况1：父亲结点和叔叔结点都是红色
			// 则都转为黑结点并将爷爷结点转为红色
			// 此时爷爷结点又不一定满足性质
			// 向上递归直到条件满足或者根节点结束
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				// 情况2：结点到爷爷结点的方向有变化
				// 通过旋转转为情况3
				if z == z.p.right {
					z = z.p
					T.LeftRotate(z)
				}
				// 情况3：到爷爷结点方向没有变化
				// 交换父亲结点和爷爷结点的颜色并旋转，保持红黑树性质
				z.p.color = BLACK
				z.p.p.color = RED
				T.RightRotate(z.p.p)
			}

		} else {
			// 对称操作
			y := z.p.p.left
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					T.RightRotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				T.LeftRotate(z.p.p)
			}
		}
	}
	T.root.color = BLACK
}

// transplant 替换子树
func (T *RedBlackTree) transplant(u, v *RBNode) {
	if u.p == T.nil {
		T.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (T *RedBlackTree) Min(node *RBNode) *RBNode {
	for node.left != T.nil {
		node = node.left
	}
	return node
}

func (T *RedBlackTree) RBDelete(z *RBNode) {
	// x表示可能需要Fixup颜色的结点
	// y表示从数中删除的或者要移入树内的结点
	y := z
	yOriginColor := y.color
	var x *RBNode
	// 情况1，2：要擅长的点的孩子小于两个，直接替换
	if z.left == T.nil {
		x = z.right
		T.transplant(z, z.right)
	} else if z.right == T.nil {
		x = z.left
		T.transplant(z, z.left)
	} else {
		//情况3：有两个孩子，找到后继，用后继的右孩子替换后继，并用后继替换y
		y = T.Min(z.right)
		yOriginColor = y.color
		x = y.right
		if y.p == z {
			x.p = y
		} else {
			T.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		T.transplant(z, y)
		y.left = z.left
		y.left.p = y
		y.color = z.color
	}
	if yOriginColor == BLACK {
		T.RBDeleteFixup(x)
	}
}

func (T *RedBlackTree) RBDeleteFixup(x *RBNode) {
	for x != T.root && x.color == BLACK {
		if x == x.p.left {
			w := x.p.right
			// 情况一，兄弟结点为红色
			// 交换兄弟结点和父亲结点颜色，并旋转
			// 使得根节点颜色不变，又能保持红黑树路径性质
			// 最后将x的兄弟结点转为黑色（情况2，3，4）
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				T.LeftRotate(x.p)
				w = x.p.right
			}
			// 情况二，兄弟结点和它的两个孩子都为黑色
			//	将兄弟结点变为红色，父亲结点上双重黑色，这里可能循环多次
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				// 情况三，兄弟结点的右孩子为黑色，且左孩子为红色
				// 通过右旋和上色，将兄弟结点变为情况四
				if w.right.color == BLACK {
					w.left.color = BLACK
					w.color = RED
					T.RightRotate(w)
					w = x.p.right
				}
				// 情况四，兄弟结点为黑色，兄弟结点的右孩子为红色
				// 通过左旋和上色消除x的双重黑色，成功保持了红黑树的性质
				w.color = x.p.color
				x.p.color = BLACK
				w.right.color = BLACK
				T.LeftRotate(x.p)
				x = T.root
			}
		} else {
			// 对称的操作
			w := x.p.left
			if w.color == RED {
				w.color = BLACK
				x.p.color = RED
				T.RightRotate(x.p)
				w = x.p.left
			}
			if w.left.color == BLACK && w.right.color == BLACK {
				w.color = RED
				x = x.p
			} else {
				if w.left.color == BLACK {
					w.right.color = BLACK
					w.color = RED
					T.LeftRotate(w)
					w = x.p.left
				}
				w.color = x.p.color
				x.p.color = BLACK
				w.left.color = BLACK
				T.RightRotate(x.p)
				x = T.root
			}
		}
	}
	x.color = BLACK
}
