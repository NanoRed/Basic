package redblacktree

import (
	"github.com/RedAFD/treeprint"
)

// 红黑树 Red-Black Tree

// Tree tree structure
type Tree struct {
	root  *Node
	count uint
}

// Node tree node structure
type Node struct {
	Key     int
	Value   interface{}
	IsBlack bool
	Left    *Node
	Right   *Node
	Parent  *Node
}

// Len total count of the tree nodes
func (t *Tree) Len() uint {
	return t.count
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
		*current = &Node{
			Key:    key,
			Value:  val,
			Parent: parent,
		}
		t.count++

		// recorrect tree
		{
			current := *current // new node is red by default
		START:
			parent := current.Parent
			if parent == nil {
				current.IsBlack = true // set the root node to black
				t.root = current       // reset root node
			} else if !parent.IsBlack && !current.IsBlack { // if parent node is red, grand node must be existed and be black
				grand := parent.Parent
				uncle := grand.Left
				if uncle == parent {
					uncle = grand.Right
				}
				if uncle == nil || uncle.IsBlack { // uncle node is black
					if grand.Left == parent {
						if parent.Right == current {
							grand.prepareRotateRight()
						}
						grand.rotateRight()
					} else {
						if parent.Left == current {
							grand.prepareRotateLeft()
						}
						grand.rotateLeft()
					}
					grand.IsBlack = false
					grand.Parent.IsBlack = true
					current = grand.Parent
					goto START
				} else { // uncle node is red
					uncle.IsBlack = true
					parent.IsBlack = true
					grand.IsBlack = false
					current = grand
					goto START
				}
			}
		}
	} else {
		(*current).Value = val
	}
}

// Remove remove a specific node from the tree
func (t *Tree) Remove(key int) {

	// find node to be removed
	var remNode *Node
	for remNode = t.root; remNode != nil; {
		if key < remNode.Key {
			remNode = remNode.Left
		} else if key > remNode.Key {
			remNode = remNode.Right
		} else {
			break
		}
	}
	if remNode == nil {
		return
	}

	// find replacement node
	if remNode.Right != nil {
		repNode := remNode.Right
		for repNode.Left != nil {
			repNode = repNode.Left
		}
		remNode.Key = repNode.Key
		remNode.Value = repNode.Value
		remNode = repNode
	} else if remNode.Left != nil {
		repNode := remNode.Left
		remNode.Key = repNode.Key
		remNode.Value = repNode.Value
		remNode = repNode
	}

	// recorrect when the node to be removed is black
	current := remNode
START:
	parent := current.Parent
	if parent != nil && current.IsBlack {
		sibling := parent.Left
		onTheLeft := true
		if sibling == current {
			sibling = parent.Right
			onTheLeft = false
		}
		if !sibling.IsBlack { // rotate for getting a black sibling
			if onTheLeft {
				parent.rotateRight()
			} else {
				parent.rotateLeft()
			}
			parent.IsBlack = false
			parent.Parent.IsBlack = true
			goto START
		}
		if onTheLeft {
		SIBLINGLEAF1:
			if sibling.Left != nil && !sibling.Left.IsBlack { // sibling's left node is red
				parent.rotateRight()
				parent.Parent.IsBlack = parent.IsBlack
				parent.IsBlack = true
				sibling.Left.IsBlack = true
			} else if sibling.Right != nil && !sibling.Right.IsBlack { // sibling's right node is red
				parent.prepareRotateRight()
				sibling.IsBlack = false
				sibling.Parent.IsBlack = true
				sibling = sibling.Parent
				goto SIBLINGLEAF1
			} else { // both sibling's leaf node is black
				if !parent.IsBlack { // parent node is red
					sibling.IsBlack = false
					parent.IsBlack = true
				} else { // parent node is black, need to ajust recursively
					sibling.IsBlack = false
					current = parent
					goto START
				}
			}
		} else {
		SIBLINGLEAF2:
			if sibling.Right != nil && !sibling.Right.IsBlack { // sibling's right node is red
				parent.rotateLeft()
				parent.Parent.IsBlack = parent.IsBlack
				parent.IsBlack = true
				sibling.Right.IsBlack = true
			} else if sibling.Left != nil && !sibling.Left.IsBlack { // sibling's left node is red
				parent.prepareRotateLeft()
				sibling.IsBlack = false
				sibling.Parent.IsBlack = true
				sibling = sibling.Parent
				goto SIBLINGLEAF2
			} else { // both sibling's leaf node is black
				if !parent.IsBlack { // parent node is red
					sibling.IsBlack = false
					parent.IsBlack = true
				} else { // parent node is black, need to ajust recursively
					sibling.IsBlack = false
					current = parent
					goto START
				}
			}
		}
	}

	// reset root node
	for t.root.Parent != nil {
		t.root = t.root.Parent
	}

	// remove node
	if remNode.Parent == nil {
		t.root = nil
	} else {
		if remNode.Parent.Left == remNode {
			if remNode.Left != nil {
				remNode.Parent.Left, remNode.Left.Parent = remNode.Left, remNode.Parent
			} else if remNode.Right != nil {
				remNode.Parent.Left, remNode.Right.Parent = remNode.Right, remNode.Parent
			} else {
				remNode.Parent.Left = nil
			}
		} else {
			if remNode.Left != nil {
				remNode.Parent.Right, remNode.Left.Parent = remNode.Left, remNode.Parent
			} else if remNode.Right != nil {
				remNode.Parent.Right, remNode.Right.Parent = remNode.Right, remNode.Parent
			} else {
				remNode.Parent.Right = nil
			}
		}
	}
	t.count--
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
	if n.IsBlack {
		return "(Black)"
	}
	return "(Red)"
}

// RangeNode implement treeprint
func (n *Node) RangeNode() chan treeprint.TreeNode {
	c := make(chan treeprint.TreeNode, 2)
	c <- n.Left
	c <- n.Right
	close(c)
	return c
}
