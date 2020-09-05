package avltree

import (
	"fmt"

	"github.com/RedAFD/treeprint"
)

// AVL树（平衡二叉搜索树） AVL Tree

// Tree tree structure
type Tree struct {
	root  *Node
	count uint
}

// Node tree node structure
type Node struct {
	Key    int
	Value  interface{}
	Height uint
	Left   *Node
	Right  *Node
	Parent *Node
}

// Len total count of the tree nodes
func (t *Tree) Len() uint {
	return t.count
}

// Height tree height
func (t *Tree) Height() uint {
	return t.root.Height
}

// Entry get entry node
func (t *Tree) Entry() *Node {
	return t.root
}

// Search search node from the tree by key
func (t *Tree) Search(key int) *Node {
	current := t.root
	for current != nil {
		if key < current.Key {
			current = current.Left
		} else if key > current.Key {
			current = current.Right
		} else {
			break
		}
	}
	return current
}

// Append append a new node to the tree
func (t *Tree) Append(key int, val interface{}) {

	// search node
	current := &t.root
	var parent *Node
	for *current != nil {
		if key < (*current).Key {
			parent = (*current)
			current = &(*current).Left
		} else if key > (*current).Key {
			parent = (*current)
			current = &(*current).Right
		} else {
			break
		}
	}
	if *current == nil {
		node := &Node{
			Key:    key,
			Value:  val,
			Height: 1,
			Parent: parent,
		}
		*current = node
		t.count++
		for node.Parent != nil {
			node = node.Parent
			node.correctHeight()
			node.rebalance()
		}
	} else {
		(*current).Value = val
	}
}

func (n *Node) rebalance() {
	difference := n.leafHeightDifference()
	if difference > 1 {
		if n.Left.leafHeightDifference() < 0 {
			n.prepareRotateRight()
		}
		n.rotateRight()
	} else if difference < -1 {
		if n.Right.leafHeightDifference() > 0 {
			n.prepareRotateLeft()
		}
		n.rotateLeft()
	}
}

func (n *Node) leafHeightDifference() float64 {
	var difference float64 = 0
	if n.Left != nil {
		difference = float64(n.Left.Height)
	}
	if n.Right != nil {
		difference = difference - float64(n.Right.Height)
	}
	return difference
}

func (n *Node) prepareRotateLeft() {
	riseNode := n.Right.Left
	fallNode := n.Right
	n.Right, riseNode.Parent = riseNode, n
	fallNode.Left = riseNode.Right
	if riseNode.Right != nil {
		riseNode.Right.Parent = fallNode
	}
	riseNode.Right, fallNode.Parent = fallNode, riseNode
}

func (n *Node) rotateLeft() {
	riseNode := n.Right
	if n.Parent == nil {
		riseNode.Parent = nil
	} else if n.Parent.Left == n {
		n.Parent.Left, riseNode.Parent = riseNode, n.Parent
	} else {
		n.Parent.Right, riseNode.Parent = riseNode, n.Parent
	}
	n.Right = riseNode.Left
	if riseNode.Left != nil {
		riseNode.Left.Parent = n
	}
	riseNode.Left, n.Parent = n, riseNode
	n.correctHeight()
	riseNode.correctHeight()
}

func (n *Node) prepareRotateRight() {
	riseNode := n.Left.Right
	fallNode := n.Left
	n.Left, riseNode.Parent = riseNode, n
	fallNode.Right = riseNode.Left
	if riseNode.Left != nil {
		riseNode.Left.Parent = fallNode
	}
	riseNode.Left, fallNode.Parent = fallNode, riseNode
}

func (n *Node) rotateRight() {
	riseNode := n.Left
	if n.Parent == nil {
		riseNode.Parent = nil
	} else if n.Parent.Left == n {
		n.Parent.Left, riseNode.Parent = riseNode, n.Parent
	} else {
		n.Parent.Right, riseNode.Parent = riseNode, n.Parent
	}
	n.Left = riseNode.Right
	if riseNode.Right != nil {
		riseNode.Right.Parent = n
	}
	riseNode.Right, n.Parent = n, riseNode
	n.correctHeight()
	riseNode.correctHeight()
}

func (n *Node) correctHeight() {
	n.Height = 0
	if n.Left != nil {
		n.Height = n.Left.Height
	}
	if n.Right != nil && n.Right.Height > n.Height {
		n.Height = n.Right.Height
	}
	n.Height++
}

// NewTree create a new tree object
func NewTree() *Tree {
	return &Tree{}
}

// GetKey implement treeprint
func (n *Node) GetKey() interface{} {
	return n.Key
}

// GetValue implement treeprint
func (n *Node) GetValue() interface{} {
	return fmt.Sprintf("(h:%d)", n.Height) // n.Value
}

// RangeNode implement treeprint
func (n *Node) RangeNode() chan treeprint.TreeNode {
	c := make(chan treeprint.TreeNode, 2)
	c <- n.Left
	c <- n.Right
	close(c)
	return c
}
