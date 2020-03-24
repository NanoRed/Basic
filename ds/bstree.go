package ds

import "errors"

// 二叉查找树 Binary Search Tree

type BSTree struct {
	root *BSTNode
	count uint
}

type BSTNode struct {
	Key int
	Value *[]interface{}
	Left *BSTNode
	Right *BSTNode
}

func (t *BSTree) Len() uint {
	return t.count
}

func (t *BSTree) Append(key int, val interface{}) *BSTree {
	var appendNode func(int, interface{}, *BSTNode) *BSTNode
	appendNode = func(key int, val interface{}, node *BSTNode) *BSTNode {
		if node == nil {
			node = &BSTNode{
				Key: key,
			}
			t.count++
			if val != nil {
				tmp := make([]interface{}, 1)
				tmp[0] = val
				node.Value = &tmp
			}
			return node
		} else if key == node.Key {
			if node.Value != nil {
				*node.Value = append(*node.Value, val)
			} else {
				tmp := make([]interface{}, 1)
				tmp[0] = val
				node.Value = &tmp
			}
			return node
		}

		if key < node.Key {
			node.Left = appendNode(key, val, node.Left)
		} else {
			node.Right = appendNode(key, val, node.Right)
		}

		return node
	}
	t.root = appendNode(key, val, t.root)
	return t
}

func (t *BSTree) Search(key int) (*BSTNode, error) {
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

func (t *BSTree) Remove(key int) *BSTree {
	// 移除节点定位
	var findPNode func(int, *BSTNode) (*BSTNode, int)
	findPNode = func(key int, node *BSTNode) (*BSTNode, int) {
		if key == node.Key {
			return nil, 0
		} else if node == nil {
			return nil, -1
		}
		var record int // 1代表要移除的节点为返回节点的左子节点，2反之，-1代表没找到
		if key < node.Key {
			tmp, tmp2 := findPNode(key, node.Left)
			if tmp != nil {
				node = tmp
			}
			if tmp2 == 0 {
				record = 1
			} else {
				record = tmp2
			}
		} else if key > node.Key {
			tmp, tmp2 := findPNode(key, node.Right)
			if tmp != nil {
				node = tmp
			}
			if tmp2 == 0 {
				record = 2
			} else {
				record = tmp2
			}
		}
		return node, record
	}

	// 取出替代节点，即取出移除节点右子树中的最小值
	var takeMin func(pNode *BSTNode, node *BSTNode, flag bool) *BSTNode
	takeMin = func(pNode *BSTNode, node *BSTNode, flag bool) *BSTNode {
		if node == nil {
			return node
		} else if node.Left == nil {
			if flag { // 函数调用时，flag恒传true，即初次为true
				pNode.Right = node.Right
			} else {
				pNode.Left = node.Right
			}
			return node
		}
		node = takeMin(node, node.Left, false)
		return node
	}

	// 移除节点
	node, record := findPNode(key, t.root)
	if record == -1 {
		return t
	} else if record == 0 {
		takeMinNode := takeMin(t.root, t.root.Right, true)
		if takeMinNode == nil {
			t.root = t.root.Left
		} else {
			takeMinNode.Left = t.root.Left
			takeMinNode.Right = t.root.Right
			t.root = takeMinNode
		}
	} else if record == 1 {
		takeMinNode := takeMin(node.Left, node.Left.Right, true)
		if takeMinNode == nil {
			node.Left = node.Left.Left
		} else {
			takeMinNode.Left = node.Left.Left
			takeMinNode.Right = node.Left.Right
			node.Left = takeMinNode
		}
	} else if record == 2 {
		takeMinNode := takeMin(node.Right, node.Right.Right, true)
		if takeMinNode == nil {
			node.Right = node.Right.Left
		} else {
			takeMinNode.Left = node.Right.Left
			takeMinNode.Right = node.Right.Right
			node.Right = takeMinNode
		}
	} else {
		panic("unknown")
	}

	return t
}
