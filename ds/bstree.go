package ds

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// 二叉查找树 Binary Search Tree

type BSTree struct {
	root *BSTNode
	count uint
	height uint
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

func (t *BSTree) Height() uint {
	return t.height
}

func (t *BSTree) Append(key int, val interface{}) *BSTree {
	var appendNode func(int, interface{}, *BSTNode) (*BSTNode, bool)
	appendNode = func(key int, val interface{}, node *BSTNode) (*BSTNode, bool) {
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
			return node, true
		} else if key == node.Key {
			if node.Value != nil {
				*node.Value = append(*node.Value, val)
			} else {
				tmp := make([]interface{}, 1)
				tmp[0] = val
				node.Value = &tmp
			}
			return node, false
		}

		if key < node.Key {
			var checkAddHeight bool
			node.Left, checkAddHeight = appendNode(key, val, node.Left)
			if checkAddHeight && node.Right == nil {
				t.height++
			}
		} else {
			var checkAddHeight bool
			node.Right, checkAddHeight = appendNode(key, val, node.Right)
			if checkAddHeight && node.Left == nil {
				t.height++
			}
		}

		return node, false
	}
	t.root, _ = appendNode(key, val, t.root)
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
	switch record {
	case -1:
		return t
	case 0:
		takeMinNode := takeMin(t.root, t.root.Right, true)
		if takeMinNode == nil {
			t.root = t.root.Left
		} else {
			takeMinNode.Left = t.root.Left
			takeMinNode.Right = t.root.Right
			t.root = takeMinNode
		}
		t.count--
	case 1:
		takeMinNode := takeMin(node.Left, node.Left.Right, true)
		if takeMinNode == nil {
			node.Left = node.Left.Left
		} else {
			takeMinNode.Left = node.Left.Left
			takeMinNode.Right = node.Left.Right
			node.Left = takeMinNode
		}
		t.count--
	case 2:
		takeMinNode := takeMin(node.Right, node.Right.Right, true)
		if takeMinNode == nil {
			node.Right = node.Right.Left
		} else {
			takeMinNode.Left = node.Right.Left
			takeMinNode.Right = node.Right.Right
			node.Right = takeMinNode
		}
		t.count--
	default:
		panic("unknown")
	}

	return t
}

func (t *BSTree) DepthFirstSearch() {
	if t.root == nil {
		return
	}
	stack := []BSTNode{ *t.root }
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

func (t *BSTree) BroadFirstSearch() {
	if t.root == nil {
		return
	}
	queue := []BSTNode{ *t.root }
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

func (t *BSTree) Print() {
	toggle := 1
	GoOn := false
	sliceA := make([]*BSTNode, 0)
	sliceB := make([]*BSTNode, 0)
	strA := make([]string, 0)
	strB := make([]string, 0)
	lines := make([][][]string, 0)
	sliceA = append(sliceA, t.root)
	for {
		if toggle == 1 {
			current := sliceA[0]
			sliceA = sliceA[1:]
			if current != nil {
				v := strconv.Itoa(current.Key)
				space := ""
				for i := len(v); i > 1; i-- {
					space += " "
				}
				strA = append(strA, v)
				if current.Left != nil {
					strB = append(strB, "|" + space)
					sliceB = append(sliceB, current.Left)
					GoOn = true
				} else {
					strB = append(strB, " " + space)
					sliceB = append(sliceB, nil)
				}
				if current.Right != nil {
					strA = append(strA, "=")
					strB = append(strB, " ")
					strA = append(strA, "+")
					strA = append(strA, " ")
					strB = append(strB, "|")
					strB = append(strB, " ")
					sliceB = append(sliceB, current.Right)
					GoOn = true
				} else {
					strA = append(strA, " ")
					strB = append(strB, " ")
					sliceB = append(sliceB, nil)
				}
			} else {
				sliceB = append(sliceB, nil, nil)
			}
			if len(sliceA) == 0 {
				strG := make([][]string, 0)
				strG = append(strG, strA, strB)
				strA = make([]string, 0)
				strB = make([]string, 0)
				lines = append(lines, strG)
				c := len(lines)
				if c > 1 {
				PrintMethodGoBack:
					tmp := strings.Join(lines[c-2][1], "")
					n := -1
					cPos := 0
					for {
						var pos int
						tmpPos := strings.Index(tmp, "|")
						if tmpPos == -1 {
							break
						} else {
							tmp = tmp[tmpPos+1:]
							cPos += tmpPos + 1
							pos = cPos - 1
							n++
						}
						var index []int
						for i, v := range lines[c-1][0] {
							match, _ := regexp.MatchString(`^\d+$`, v)
							if match {
								index = append(index, i)
							}
						}
						front := strings.Join(lines[c-1][0][:index[n]], "")
						pos2 := len(front)
						if pos > pos2 {
							blank := ""
							for i := 0; i < pos - pos2; i++ {
								blank += " "
							}
							lines[c-1][0] = append(
								lines[c-1][0][:index[n]],
								append(
									[]string{blank},
									lines[c-1][0][index[n]:]...
								)...
							)
							lines[c-1][1] = append(
								lines[c-1][1][:index[n]],
								append(
									[]string{blank},
									lines[c-1][1][index[n]:]...
								)...
							)
						} else if pos < pos2 {
							blank := ""
							for i := 0; i < pos2 - pos; i++ {
								blank += " "
							}
							var index2 []int
							for i, v := range lines[c-2][1] {
								if v[0:1] == "|" {
									index2 = append(index2, i)
								}
							}
							lines[c-2][1] = append(
								lines[c-2][1][:index2[n]],
								append([]string{blank}, lines[c-2][1][index2[n]:]...)...
							)
							if lines[c-2][0][index2[n]][0:1] == "+" {
								lines[c-2][0] = append(
									lines[c-2][0][:index2[n]],
									append(
										[]string{strings.ReplaceAll(blank, " ", "=")},
										lines[c-2][0][index2[n]:]...
									)...
								)
							} else {
								lines[c-2][0] = append(
									lines[c-2][0][:index2[n]],
									append([]string{blank}, lines[c-2][0][index2[n]:]...)...
								)
							}
							goto PrintMethodGoBack
						}
					}
				}
				if GoOn {
					GoOn = false
					toggle *= -1
				} else {
					break
				}
			}
		} else {
			current := sliceB[0]
			sliceB = sliceB[1:]
			if current != nil {
				v := strconv.Itoa(current.Key)
				space := ""
				for i := len(v); i > 1; i-- {
					space += " "
				}
				strA = append(strA, v)
				if current.Left != nil {
					strB = append(strB, "|" + space)
					sliceA = append(sliceA, current.Left)
					GoOn = true
				} else {
					strB = append(strB, " " + space)
					sliceA = append(sliceA, nil)
				}
				if current.Right != nil {
					strA = append(strA, "=")
					strB = append(strB, " ")
					strA = append(strA, "+")
					strA = append(strA, " ")
					strB = append(strB, "|")
					strB = append(strB, " ")
					sliceA = append(sliceA, current.Right)
					GoOn = true
				} else {
					strA = append(strA, " ")
					strB = append(strB, " ")
					sliceA = append(sliceA, nil)
				}
			} else {
				sliceA = append(sliceA, nil, nil)
			}
			if len(sliceB) == 0 {
				strG := make([][]string, 0)
				strG = append(strG, strA, strB)
				strA = make([]string, 0)
				strB = make([]string, 0)
				lines = append(lines, strG)
				c := len(lines)
				if c > 1 {
				PrintMethodGoBack2:
					tmp := strings.Join(lines[c-2][1], "")
					n := -1
					cPos := 0
					for {
						var pos int
						tmpPos := strings.Index(tmp, "|")
						if tmpPos == -1 {
							break
						} else {
							tmp = tmp[tmpPos+1:]
							cPos += tmpPos + 1
							pos = cPos - 1
							n++
						}
						var index []int
						for i, v := range lines[c-1][0] {
							match, _ := regexp.MatchString(`^\d+$`, v)
							if match {
								index = append(index, i)
							}
						}
						front := strings.Join(lines[c-1][0][:index[n]], "")
						pos2 := len(front)
						if pos > pos2 {
							blank := ""
							for i := 0; i < pos - pos2; i++ {
								blank += " "
							}
							lines[c-1][0] = append(
								lines[c-1][0][:index[n]],
								append(
									[]string{blank},
									lines[c-1][0][index[n]:]...
								)...
							)
							lines[c-1][1] = append(
								lines[c-1][1][:index[n]],
								append(
									[]string{blank},
									lines[c-1][1][index[n]:]...
								)...
							)
						} else if pos < pos2 {
							blank := ""
							for i := 0; i < pos2 - pos; i++ {
								blank += " "
							}
							var index2 []int
							for i, v := range lines[c-2][1] {
								if v[0:1] == "|" {
									index2 = append(index2, i)
								}
							}
							lines[c-2][1] = append(
								lines[c-2][1][:index2[n]],
								append([]string{blank}, lines[c-2][1][index2[n]:]...)...
							)
							if lines[c-2][0][index2[n]][0:1] == "+" {
								lines[c-2][0] = append(
									lines[c-2][0][:index2[n]],
									append(
										[]string{strings.ReplaceAll(blank, " ", "=")},
										lines[c-2][0][index2[n]:]...
									)...
								)
							} else {
								lines[c-2][0] = append(
									lines[c-2][0][:index2[n]],
									append([]string{blank}, lines[c-2][0][index2[n]:]...)...
								)
							}
							goto PrintMethodGoBack2
						}
					}
				}
				if GoOn {
					GoOn = false
					toggle *= -1
				} else {
					break
				}
			}
		}
	}
	for i := 0; i < len(lines); i++ {
		rowA := strings.Join(lines[i][0], "")
		rowA = strings.ReplaceAll(rowA, "=", "─")
		rowA = strings.ReplaceAll(rowA, "+", "┐")
		rowB := strings.Join(lines[i][1], "")
		fmt.Println(rowA + "\n" + rowB)
	}
}

func NewBSTree() *BSTree {
	return &BSTree{}
}
