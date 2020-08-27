package binarysearchtree

// 二叉查找树 Binary Search Tree

import (
	"errors"
	"fmt"
	"strings"
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

// Append append a new node to the tree
func (t *Tree) Append(key int, val interface{}) *Tree {
	if update(&t.root, key, val) {
		t.count++
	}
	return t
}

func update(node **Node, key int, val interface{}) (countAdd bool) {
	if *node == nil {
		*node = &Node{
			Key:    key,
			Value:  val,
			Height: 1,
		}
		countAdd = true
		return
	} else if key == (*node).Key {
		(*node).Value = val
		return
	}
	if key < (*node).Key {
		countAdd = update(&(*node).Left, key, val)
		if (*node).Left.Height == (*node).Height {
			(*node).Height++
		}
	} else {
		countAdd = update(&(*node).Right, key, val)
		if (*node).Right.Height == (*node).Height {
			(*node).Height++
		}
	}
	return
}

// Search search node from the tree with key
func (t *Tree) Search(key int) (*Node, error) {
	current := t.root
	for current != nil {
		switch {
		case key < current.Key:
			current = current.Left
		case key > current.Key:
			current = current.Right
		default:
			return current, nil
		}
	}
	return nil, errors.New("not found")
}

type side int

const (
	locatedAtNoWhere side = iota
	locatedAtItself
	locatedAtLeft
	locatedAtRight
)

// Remove remove a pecific node from the tree
func (t *Tree) Remove(key int) *Tree {
	parentNode, side, path := t.root.findParent(key)
	switch side {
	case locatedAtItself:
		replacement := t.root.takeOutReplacement()
		if replacement == nil {
			t.root = t.root.Left
		} else {
			replacement.Left, replacement.Right, replacement.Height = t.root.Left, t.root.Right, t.root.Height
			t.root = replacement
		}
		t.count--
	case locatedAtLeft:
		replacement := parentNode.Left.takeOutReplacement()
		if replacement == nil {
			parentNode.Left = parentNode.Left.Left
		} else {
			replacement.Left, replacement.Right, replacement.Height = parentNode.Left.Left, parentNode.Left.Right, parentNode.Left.Height
			parentNode.Left = replacement
		}
		for _, val := range path {
			val.correctHeight()
		}
		t.count--
	case locatedAtRight:
		replacement := parentNode.Right.takeOutReplacement()
		if replacement == nil {
			parentNode.Right = parentNode.Right.Left
		} else {
			replacement.Left, replacement.Right, replacement.Height = parentNode.Right.Left, parentNode.Right.Right, parentNode.Right.Height
			parentNode.Right = replacement
		}
		for _, val := range path {
			val.correctHeight()
		}
		t.count--
	}
	return t
}

// locate the parent node of the node to be removed
// side means the node to be removed is in *side* side of the returned parent node
// path means the path from the returned parent node to the root node
func (n *Node) findParent(key int) (parentNode *Node, side side, path []*Node) {
	if key == n.Key {
		side = locatedAtItself
		return
	} else if n == nil {
		side = locatedAtNoWhere
		return
	}
	if key < n.Key {
		if parentNode, side, path = n.Left.findParent(key); side == locatedAtItself {
			parentNode = n
			side = locatedAtLeft
			path = []*Node{n}
		} else {
			path = append(path, n)
		}
	} else {
		if parentNode, side, path = n.Right.findParent(key); side == locatedAtItself {
			parentNode = n
			side = locatedAtRight
			path = []*Node{n}
		} else {
			path = append(path, n)
		}
	}
	return
}

// find and take out the replacement node
// which is the largest node of the right subtree of the node to be removed
func (n *Node) takeOutReplacement(params ...interface{}) (replacementNode *Node) {
	if len(params) == 0 {
		// non-recursive logic
		params = append(params, true)
		replacementNode = n.Right
		if replacementNode == nil {
			return
		} else if replacementNode.Left == nil {
			n.Right = replacementNode.Right
			n.correctHeight()
			return
		}
	} else {
		// the recursive part
		replacementNode = n.Left
		if replacementNode.Left == nil {
			n.Left = replacementNode.Right
			n.correctHeight()
			return
		}
	}
	replacementNode = replacementNode.takeOutReplacement(params...)
	n.correctHeight()
	return
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
	return
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

// Sprint print the tree data into a return variable
func (t *Tree) Sprint() (content string) {
	if t.root == nil {
		return
	}
	type nodeInfo struct {
		node       *Node
		layer      int
		count      int
		index      int
		len        int
		str        string
		leftNode   *nodeInfo
		rightNode  *nodeInfo
		parentNode *nodeInfo
	}
	layer := make([][]*nodeInfo, 0)
	first := &nodeInfo{
		node:  t.root,
		layer: 1,
		count: 1,
		str:   fmt.Sprintf("%d(h:%d)", t.root.Key, t.root.Height),
	}
	first.len = len(first.str)
	queue := []*nodeInfo{first}
	currentLayer := 2
	currentIndex := 0
	currentCount := 1
	for i := 0; i < int(t.count); i++ {
		current := queue[i]
		if current.node.Left != nil {
			left := &nodeInfo{
				node:       current.node.Left,
				layer:      current.layer + 1,
				str:        fmt.Sprintf("%d(h:%d)", current.node.Left.Key, current.node.Left.Height),
				parentNode: current,
			}
			left.len = len(left.str)
			if left.layer != currentLayer {
				currentLayer = left.layer
				currentIndex = 0
				currentCount = 1
			}
			left.count = currentCount
			currentCount++
			left.index = currentIndex
			currentIndex += left.len + 1
			queue = append(queue, left)
			current.leftNode = left
		}
		if current.node.Right != nil {
			right := &nodeInfo{
				node:       current.node.Right,
				layer:      current.layer + 1,
				str:        fmt.Sprintf("%d(h:%d)", current.node.Right.Key, current.node.Right.Height),
				parentNode: current,
			}
			right.len = len(right.str)
			if right.layer != currentLayer {
				currentLayer = right.layer
				currentIndex = 0
				currentCount = 1
			}
			right.count = currentCount
			currentCount++
			right.index = currentIndex
			currentIndex += right.len + 1
			queue = append(queue, right)
			current.rightNode = right
		}
		if current.layer > len(layer) {
			layer = append(layer, make([]*nodeInfo, 0))
		}
		layer[current.layer-1] = append(layer[current.layer-1], current)
	}
	var alignLeft func(*nodeInfo)
	var alignRight func(*nodeInfo)
	alignLeft = func(current *nodeInfo) {
		if current.leftNode == nil {
			return
		} else if val := current.index - current.leftNode.index; val > 0 { // 下一层移位
			for k := current.leftNode.count - 1; k < len(layer[current.layer]); k++ {
				layer[current.layer][k].index += val
			}
		} else if val < 0 { // 本层移位
			for k := current.count - 1; k < len(layer[current.layer-1]); k++ {
				layer[current.layer-1][k].index += -val
			}
		}
		for j := current.leftNode.count - 1; j < len(layer[current.layer]); j++ {
			alignLeft(layer[current.layer][j])
			alignRight(layer[current.layer][j])
		}
	}
	alignRight = func(current *nodeInfo) {
		if current.rightNode == nil {
			return
		} else if !strings.Contains(current.str, "-+") {
			if val := current.rightNode.index - current.index - current.len; val > 0 {
				tmp := ""
				for k := 0; k < val; k++ {
					tmp += "-"
				}
				tmp += "+"
				current.str += tmp
				newLen := len(current.str)
				offset := newLen - current.len
				current.len = newLen
				for k := current.count; k < len(layer[current.layer-1]); k++ {
					layer[current.layer-1][k].index += offset
				}
			} else {
				tmp := "-+"
				current.str += tmp
				newLen := len(current.str)
				offset := newLen - current.len
				current.len = newLen
				for k := current.count; k < len(layer[current.layer-1]); k++ {
					layer[current.layer-1][k].index += offset
				}
				offset2 := -val + 1
				for k := current.rightNode.count - 1; k < len(layer[current.layer]); k++ {
					layer[current.layer][k].index += offset2
				}
			}
		} else {
			for k := current.rightNode.count - 1; k < len(layer[current.layer]); k++ {
				layer[current.layer][k].index += current.index + current.len - 1 - current.rightNode.index
			}
		}
		for j := current.rightNode.count - 1; j < len(layer[current.layer]); j++ {
			alignLeft(layer[current.layer][j])
			alignRight(layer[current.layer][j])
		}
	}
	for i := len(layer) - 2; i >= 0; i-- {
		for j := 0; j < len(layer[i]); j++ {
			alignLeft(layer[i][j])
			alignRight(layer[i][j])
		}
	}
	for i := 0; i < len(layer); i++ {
		line1 := ""
		line2 := ""
		for j := 0; j < len(layer[i]); j++ {
			if val := layer[i][j].index - len(line1); val > 0 {
				for k := 0; k < val; k++ {
					line1 += " "
					line2 += " "
				}
			}
			line1 += layer[i][j].str
			if layer[i][j].leftNode != nil {
				line2 += "|"
			} else {
				line2 += " "
			}
			for k := 0; k < layer[i][j].len-2; k++ {
				line2 += " "
			}
			if layer[i][j].rightNode != nil {
				line2 += "|"
			} else {
				line2 += " "
			}
		}
		line1 = strings.ReplaceAll(line1, "-", "─")
		line1 = strings.ReplaceAll(line1, "+", "┐")
		content += line1 + "\n"
		content += line2 + "\n"
	}
	return
}

// NewTree create a new tree object
func NewTree() *Tree {
	return &Tree{}
}
