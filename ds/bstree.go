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
	LeftChild *BSTNode
	RightChild *BSTNode
}

func (t *BSTree) Len() uint {
	return t.count
}

func (t *BSTree) Append(key int, val interface{}) *BSTree {
	curNode := t.root
	if curNode == nil {
		t.root = &BSTNode{
			Key: key,
		}
		if val != nil {
			tmp := make([]interface{}, 1)
			tmp[0] = val
			t.root.Value = &tmp
		}
		return t
	}
	for curNode != nil {
		if key < curNode.Key {
			if curNode.LeftChild == nil {
				curNode.LeftChild = &BSTNode{
					Key: key,
				}
				if val != nil {
					tmp := make([]interface{}, 1)
					tmp[0] = val
					curNode.LeftChild.Value = &tmp
				}
				break
			} else {
				curNode = curNode.LeftChild
			}
		} else if key > curNode.Key {
			if curNode.RightChild == nil {
				curNode.RightChild = &BSTNode{
					Key: key,
				}
				if val != nil {
					tmp := make([]interface{}, 1)
					tmp[0] = val
					curNode.RightChild.Value = &tmp
				}
				break
			} else {
				curNode = curNode.RightChild
			}
		} else {
			if val != nil {
				if curNode.Value != nil {
					*curNode.Value = append(*curNode.Value, val)
				} else {
					tmp := make([]interface{}, 1)
					tmp[0] = val
					curNode.Value = &tmp
				}
			}
			break
		}
	}
	return t
}

func (t *BSTree) Search(key int) (*BSTNode, error) {
	curNode := t.root
	for curNode != nil {
		switch {
		case key < curNode.Key:
			curNode = curNode.LeftChild
		case key > curNode.Key:
			curNode = curNode.RightChild
		default:
			return curNode, nil
		}
	}
	return nil, errors.New("not found")
}

//func (t *BSTree) Remove(key int) *BSTree {
//}
