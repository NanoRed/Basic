package ds

import "fmt"

// 树 Tree

type Tree struct {
	root *TNode
	count uint
}

type TNode struct {
	Value interface{}
	Child []*TNode
}

func (t *Tree) Len() uint {
	return t.count
}

func (t *Tree) Append(index *[]int, val interface{}) *Tree {
	if index == nil {
		t.root = &TNode{
			Value: val,
		}
		t.count++
		return t
	}
	node := t.root
	for i := 0; i < len(*index); i++ {
		node = node.Child[(*index)[i]]
	}
	node.Child = append(node.Child, &TNode{
		Value: val,
	})
	t.count++
	return t
}

func (t *Tree) Remove(index *[]int) *Tree {
	if index == nil {
		t.root = nil
		t.count = 0
		return t
	}
	node := t.root
	count := len(*index)
	for i := 0; i < count; i++ {
		if i == count - 1 {
			node.Child = append(node.Child[:(*index)[i]], node.Child[(*index)[i]+1:]...)
			break
		}
		node = node.Child[(*index)[i]]
	}
	return t
}

func (t *Tree) DLR() {
	if t.root == nil {
		return
	}
	var f func(*TNode)
	f = func(node *TNode) {
		fmt.Println(node.Value)
		for i := 0; i < len(node.Child); i++ {
			f(node.Child[i])
		}
	}
	f(t.root)
}

func (t *Tree) LRD() {
	if t.root == nil {
		return
	}
	var f func(*TNode)
	f = func(node *TNode) {
		for i := 0; i < len(node.Child); i++ {
			f(node.Child[i])
		}
		fmt.Println(node.Value)
	}
	f(t.root)
}

func (t *Tree) DepthFirstSearch() {
	if t.root == nil {
		return
	}
	stack := []TNode{ *t.root }
	fmt.Println(stack[0].Value)
	for {
		if len(stack) == 0 {
			break
		}
		if len(stack[0].Child) == 0 {
			stack = stack[1:]
			if len(stack) != 0 {
				stack[0].Child = stack[0].Child[1:]
			}
		} else {
			stack = append([]TNode{ *stack[0].Child[0] }, stack...)
			fmt.Println(stack[0].Value)
		}
	}
}

func (t *Tree) BroadFirstSearch() {
	if t.root == nil {
		return
	}
	queue := []*TNode{ t.root }
	queueLen := 1
	for queueLen > 0 {
		current := queue[0]
		fmt.Println(current.Value)
		queue = queue[1:]
		queueLen--
		cLen := len(current.Child)
		if cLen > 0 {
			queue = append(queue, current.Child...)
			queueLen += cLen
		}
	}
}

func NewTree() *Tree {
	return &Tree{}
}
