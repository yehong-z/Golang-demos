package SearchTree

import "fmt"

type BTree struct {
	root *BTreeNode
	t    int
}

type BTreeNode struct {
	keys []int
	c    []*BTreeNode
	leaf bool
}

func BTreeCreate(t int) *BTree {
	root := &BTreeNode{
		leaf: true,
		keys: make([]int, 1),
		c:    nil,
	}
	return &BTree{root: root, t: t}
}

func (T *BTree) DFS(node *BTreeNode) {
	for i := 1; i < len(node.keys); i++ {
		if !node.leaf {
			T.DFS(node.c[i])
		}
		fmt.Printf("%v ", node.keys[i])
	}
	if !node.leaf {
		T.DFS(node.c[len(node.keys)])
	}
}

func (T *BTree) SearchNode(x *BTreeNode, k int) (*BTreeNode, int) {
	i := 1
	for i < len(x.keys) && k > x.keys[i] {
		i++
	}
	if i < len(x.keys) && k == x.keys[i] {
		return x, i
	} else if x.leaf {
		return nil, -1
	} else {
		return T.SearchNode(x.c[i], k)
	}
}

func (T *BTree) SplitChild(x *BTreeNode, i int) {
	y := x.c[i]
	z := &BTreeNode{
		leaf: y.leaf,
		keys: make([]int, 1),
	}
	// 移动i右侧的children
	x.c = append(x.c, nil)
	for j := len(x.c) - 1; j > i+1; j-- {
		x.c[j] = x.c[j-1]
	}
	x.c[i+1] = z
	// 移动i右侧的key，包括i
	x.keys = append(x.keys, 0)
	for j := len(x.keys) - 1; j > i; j-- {
		x.keys[j] = x.keys[j-1]
	}
	x.keys[i] = y.keys[T.t]
	// 将y中一半的key结点移动到z
	for j := T.t + 1; j < len(y.keys); j++ {
		z.keys = append(z.keys, y.keys[j])
	}
	y.keys = y.keys[:T.t]
	if !y.leaf {
		z.c = make([]*BTreeNode, 1)
		for j := T.t + 1; j < len(y.c); j++ {
			z.c = append(z.c, y.c[j])
		}
		y.c = y.c[:T.t+1]
	}
}

func (T *BTree) Insert(k int) {
	r := T.root
	if len(r.keys) == 2*T.t {
		s := &BTreeNode{
			leaf: false,
			keys: make([]int, 1),
			c:    []*BTreeNode{nil, r},
		}
		T.root = s
		T.SplitChild(s, 1)
		T.InsertNonFull(s, k)
	} else {
		T.InsertNonFull(r, k)
	}
}

func (T *BTree) InsertNonFull(x *BTreeNode, k int) {
	i := len(x.keys) - 1
	if x.leaf {
		x.keys = append(x.keys, 0)
		for i >= 1 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
	} else {
		for i >= 1 && k < x.keys[i] {
			i--
		}
		i++
		if len(x.c[i].keys) == 2*T.t {
			T.SplitChild(x, i)
			if k > x.keys[i] {
				i++
			}
		}
		T.InsertNonFull(x.c[i], k)
	}
}

//func (T *BTree) Delete(x *BTreeNode, k int) {
//	i := 1
//	for i < len(x.keys) && x.keys[i] < k {
//		i++
//	}
//	if x.keys[i] == k {
//		if x.leaf {
//			for j := i; j < len(x.keys)-1; j++ {
//				x.keys[i] = x.keys[i+1]
//			}
//			x.keys = x.keys[:len(x.keys)-1]
//		} else {
//			y := x.c[i]
//			z := x.c[i+1]
//			if len(x.c[i].keys) > T.t {
//				x.keys[i] = y.keys[len(y.keys)-1]
//				T.Delete(y, y.keys[len(y.keys)-1])
//			} else if len(z.keys) > T.t {
//				x.keys[i] = z.keys[1]
//				T.Delete(z, z.keys[1])
//			} else {
//
//			}
//		}
//	}
//}
