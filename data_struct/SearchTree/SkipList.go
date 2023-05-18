package SearchTree

import "math/rand"

const MaxLevel = 32

type SkipListNode struct {
	val  int
	next []*SkipListNode
}

type SkipList struct {
	head  *SkipListNode
	level int
	SearchTree
}

func NewSkipList() *SkipList {
	return &SkipList{
		head: &SkipListNode{
			val:  -1,
			next: make([]*SkipListNode, MaxLevel),
		},
		level: 0,
	}
}

func (s *SkipList) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < target {
			curr = curr.next[i]
		}
	}
	curr = curr.next[0]
	return curr != nil && curr.val == target
}

func (s *SkipList) Add(num int) {
	update := make([]*SkipListNode, MaxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	lv := randLevel()
	s.level = max(s.level, lv)
	newNode := &SkipListNode{num, make([]*SkipListNode, lv)}
	for i, node := range update[:lv] {
		newNode.next[i] = node.next[i]
		node.next[i] = newNode
	}
}

func (s *SkipList) Erase(num int) bool {
	update := make([]*SkipListNode, MaxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		for curr.next[i] != nil && curr.next[i].val < num {
			curr = curr.next[i]
		}
		update[i] = curr
	}
	curr = curr.next[0]
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].next[i] == curr; i++ {
		update[i].next[i] = curr.next[i]
	}
	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func randLevel() int {
	i := 1
	for i < MaxLevel && rand.Intn(2) == 0 {
		i++
	}
	return i
}
