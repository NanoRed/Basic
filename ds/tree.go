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
		return t
	}
	node := t.root
	for i := 0; i < len(*index); i++ {
		node = node.Child[(*index)[i]]
	}
	node.Child = append(node.Child, &TNode{
		Value: val,
	})
	return t
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

func NewTree() *Tree {
	return &Tree{}
}
