package tree

// 树 Tree

import (
	"errors"
	"fmt"
	"github.com/RedAFD/treeprint"
)

// Tree tree structure
type Tree struct {
	root  *Node
	count uint
}

// Node tree node structure
type Node struct {
	Value interface{}
	Leaf []*Node
}

// Len total count of the tree nodes
func (t *Tree) Len() uint {
	return t.count
}

// Search search node from the tree by index
func (t *Tree) Search(index []int) (node *Node, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("index out of range")
		}
	}()
	node = t.root
	layer := len(index)
	for i := 0; i < layer; i++ {
		node = node.Leaf[index[i]]
	}
	if node == nil {
		err = errors.New("index out of range")
	}
	return
}

// Append append a new node to the tree by index
// pass []int{} to index means root node
// pass []int{0} means the 1st leaf node of root
// pass []int{0, 1} means the 2nd leaf node of the 1st leaf node of root
// etc.
func (t *Tree) Append(index []int, val interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("index out of range")
		}
	}()
	entry := &t.root
	layer := len(index)
	for i := 0; i < layer; i++ {
		if i == layer-1 {
			if index[i] == len((*entry).Leaf) {
				(*entry).Leaf = append((*entry).Leaf, nil)
			}
		}
		entry = &(*entry).Leaf[index[i]]
	}
	if *entry == nil {
		*entry = &Node{
			Value: val,
		}
		t.count++
	} else {
		(*entry).Value = val
	}
	return
}

// Remove remove a specific node from the tree by index
// pass []int{} to index means root node
// pass []int{0} means the 1st leaf node of root
// pass []int{0, 1} means the 2nd leaf node of the 1st leaf node of root
// etc.
func (t *Tree) Remove(index []int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("index out of range")
		}
	}()
	entry := &t.root
	layer := len(index)
	for i := 0; i < layer; i++ {
		if i == layer-1 {
			(*entry).Leaf = append((*entry).Leaf[:index[i]], (*entry).Leaf[index[i]+1:]...)
			return
		}
		entry = &(*entry).Leaf[index[i]]
	}
	*entry = nil
	return
}

// DLR pre-order traversal
func (t *Tree) DLR() {
	if t.root == nil {
		return
	}
	var f func(*Node)
	f = func(node *Node) {
		fmt.Println(node.Value)
		for i := 0; i < len(node.Leaf); i++ {
			f(node.Leaf[i])
		}
	}
	f(t.root)
}

// LRD post-order traversal
func (t *Tree) LRD() {
	if t.root == nil {
		return
	}
	var f func(*Node)
	f = func(node *Node) {
		for i := 0; i < len(node.Leaf); i++ {
			f(node.Leaf[i])
		}
		fmt.Println(node.Value)
	}
	f(t.root)
}

// DepthFirstSearch depth first search
func (t *Tree) DepthFirstSearch() {
	if t.root == nil {
		return
	}
	stack := []Node{*t.root}
	fmt.Println(stack[0].Value)
	for {
		if len(stack) == 0 {
			break
		}
		if len(stack[0].Leaf) == 0 {
			stack = stack[1:]
			if len(stack) != 0 {
				stack[0].Leaf = stack[0].Leaf[1:]
			}
		} else {
			stack = append([]Node{*stack[0].Leaf[0]}, stack...)
			fmt.Println(stack[0].Value)
		}
	}
}

// BroadFirstSearch broad first search
func (t *Tree) BroadFirstSearch() {
	if t.root == nil {
		return
	}
	queue := []*Node{t.root}
	queueLen := 1
	for queueLen > 0 {
		current := queue[0]
		fmt.Println(current.Value)
		queue = queue[1:]
		queueLen--
		cLen := len(current.Leaf)
		if cLen > 0 {
			queue = append(queue, current.Leaf...)
			queueLen += cLen
		}
	}
}

// NewTree create a new tree object
func NewTree() *Tree {
	return &Tree{}
}

// GetKey implement treeprint
func (n *Node) GetKey() interface{} {
	return n.Value
}

// GetValue implement treeprint
func (n *Node) GetValue() interface{} {
	return ""
}

// RangeNode implement treeprint
func (n *Node) RangeNode() chan treeprint.TreeNode {
	count := len(n.Leaf)
	c := make(chan treeprint.TreeNode, count)
	for i := 0; i < count; i++ {
		c <- n.Leaf[i]
	}
	close(c)
	return c
}