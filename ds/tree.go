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
	lastIndex := 0
	fmt.Println(stack[lastIndex].Value)
	for lastIndex >= 0 {
		if len(stack[lastIndex].Child) == 0 {
			stack = stack[:lastIndex]
			lastIndex--
			if lastIndex >= 0 {
				stack[lastIndex].Child = stack[lastIndex].Child[1:]
			}
		} else {
			stack = append(stack, *stack[lastIndex].Child[0])
			lastIndex++
			fmt.Println(stack[lastIndex].Value)
		}
	}
}

func NewTree() *Tree {
	return &Tree{}
}
