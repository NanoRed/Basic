package binarysearchtree

// 二叉查找树 Binary Search Tree

import (
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
	Key    int
	Value  interface{}
	Height uint
	Left   *Node
	Right  *Node
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

	// nodes that need to be increased in height
	increaseHeight := make(map[*Node]struct{})
	increaseTracking := func(parent *Node, selected *Node, another *Node) {
		// if selected leaf node height equal or higher than the another one, than it means has to increase
		if another == nil || (selected != nil && selected.Height >= another.Height) {
			increaseHeight[parent] = struct{}{}
		} else {
			increaseHeight = make(map[*Node]struct{})
		}
	}

	// search node
	current := &t.root
	for *current != nil {
		if key < (*current).Key {
			increaseTracking((*current), (*current).Left, (*current).Right)
			current = &(*current).Left
		} else if key > (*current).Key {
			increaseTracking((*current), (*current).Right, (*current).Left)
			current = &(*current).Right
		} else {
			break
		}
	}
	if *current == nil {
		*current = &Node{
			Key:    key,
			Value:  val,
			Height: 1,
		}
		t.count++
		for node := range increaseHeight {
			node.Height++
		}
	} else {
		(*current).Value = val
	}
}

// Remove remove a specific node from the tree
func (t *Tree) Remove(key int) {

	// nodes that need to be reduced in height
	reduceHeight := make(map[*Node]struct{})
	reduceTracking := func(parent *Node, selected *Node, another *Node) {
		// if another leaf node height equal or higher than the selected one, than it means no need to reduce
		if selected == nil || (another != nil && another.Height >= selected.Height) {
			reduceHeight = make(map[*Node]struct{})
		} else {
			reduceHeight[parent] = struct{}{}
		}
	}

	// find node to be removed and its parent node
	var remNode *Node
	var remNodeParent *Node
	for remNode = t.root; remNode != nil; {
		if key < remNode.Key {
			reduceTracking(remNode, remNode.Left, remNode.Right)
			remNodeParent = remNode
			remNode = remNode.Left
		} else if key > remNode.Key {
			reduceTracking(remNode, remNode.Right, remNode.Left)
			remNodeParent = remNode
			remNode = remNode.Right
		} else {
			break
		}
	}
	if remNode == nil {
		return
	}

	// find replacement node and take out
	var repNode *Node
	if remNode.Right != nil {
		reduceTracking(remNode, remNode.Right, remNode.Left)
		repNode = remNode.Right
		repNodeParent := remNode
		for repNode.Left != nil {
			reduceTracking(repNode, repNode.Left, repNode.Right)
			repNodeParent = repNode
			repNode = repNode.Left
		}
		if repNodeParent == remNode {
			repNodeParent.Right = repNode.Right
		} else {
			repNodeParent.Left = repNode.Right
		}
	} else if remNode.Left != nil {
		reduceHeight[remNode] = struct{}{} // not need to call reduceTracking(), 100% need to be reduced
		repNode = remNode.Left
		remNode.Left = repNode.Left
		remNode.Right = repNode.Right
	}

	// reduce node height
	for node := range reduceHeight {
		node.Height--
	}

	// replace node
	if repNode != nil {
		repNode.Left = remNode.Left
		repNode.Right = remNode.Right
		repNode.Height = remNode.Height
	}
	if remNodeParent != nil {
		if remNodeParent.Left == remNode {
			remNodeParent.Left = repNode
		} else {
			remNodeParent.Right = repNode
		}
	} else {
		t.root = repNode
	}

	t.count--
}

// DepthFirstSearch depth first search
func (t *Tree) DepthFirstSearch() {
	if t.root == nil {
		return
	}
	stack := []Node{*t.root}
	var stackLen uint = 1
	entrance := stack[0]
	for {
		for _, value := range stack {
			fmt.Printf("%#v ", value.Key)
		}
		fmt.Println()
		if entrance.Left != nil {
			entrance = *entrance.Left
			stack[stackLen-1].Left = nil
			stack = append(stack, entrance)
			stackLen++
		} else if entrance.Right != nil {
			entrance = *entrance.Right
			stack[stackLen-1].Right = nil
			stack = append(stack, entrance)
			stackLen++
		} else {
			stack = stack[:stackLen-1]
			stackLen--
			if stackLen <= 0 {
				break
			}
			entrance = stack[stackLen-1]
		}
	}
}

// BroadFirstSearch broad first search
func (t *Tree) BroadFirstSearch() {
	if t.root == nil {
		return
	}
	queue := []Node{*t.root}
	var queueLen uint = 1
	for {
		current := queue[0]
		queue = queue[1:]
		queueLen--
		if current.Left != nil {
			queue = append(queue, *current.Left)
			queueLen++
		}
		if current.Right != nil {
			queue = append(queue, *current.Right)
			queueLen++
		}
		fmt.Println(current.Key)
		if queueLen == 0 {
			break
		}
	}
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
