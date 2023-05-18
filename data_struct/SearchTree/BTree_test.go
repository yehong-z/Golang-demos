package SearchTree

import "testing"

func TestBTreeInsert(t *testing.T) {
	BTree := BTreeCreate(2)
	for i := 0; i < 100; i++ {
		BTree.Insert(i)
	}
	BTree.DFS(BTree.root)
}
