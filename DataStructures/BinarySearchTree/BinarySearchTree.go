package BinarySearchTree

// 二叉查找树 Binary Search Tree

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Tree struct {
	root *Node
	count uint
}

type Node struct {
	Key int
	Value interface{}
	Height uint
	Left *Node
	Right *Node
}

func (t *Tree) Len() uint {
	return t.count
}

func (t *Tree) Height() uint {
	return t.root.Height
}

func (t *Tree) Append(key int, val interface{}) *Tree {
	if t.root.update(key, val) {
		t.count++
	}
	return t
}

func (n *Node) update(key int, val interface{}) (countAdd bool) {
	if n == nil {
		n = &Node{
			Key: key,
			Value: val,
			Height: 1,
		}
		countAdd = true
		return
	} else if key == n.Key {
		n.Value = val
		return
	}
	if key < n.Key {
		countAdd = n.Left.update(key, val)
		if n.Left.Height == n.Height {
			n.Height++
		}
	} else {
		countAdd = n.Right.update(key, val)
		if n.Right.Height == n.Height {
			n.Height++
		}
	}
	return
}

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

type SIDE int

const (
	locatedAtNoWhere SIDE = iota
	locatedAtItself
	locatedAtLeft
	locatedAtRight
)

func (t *Tree) Remove(key int) *Tree {
	parentNode, side := t.root.findParent(key)
	switch side {
	case locatedAtItself:
		replacement := t.root.takeOutReplacement()
		if replacement == nil {
			t.root = t.root.Left
		} else {
			replacement.Left, replacement.Right = t.root.Left, t.root.Right
			t.root = replacement
		}
		t.count--
	case locatedAtLeft:
		replacement := parentNode.Left.takeOutReplacement()
		if replacement == nil {
			parentNode.Left = parentNode.Left.Left
		} else {
			replacement.Left, replacement.Right = parentNode.Left.Left, parentNode.Left.Right
			parentNode.Left = replacement
		}
		t.count--
	case locatedAtRight:
		replacement := parentNode.Right.takeOutReplacement()
		if replacement == nil {
			parentNode.Right = parentNode.Right.Left
		} else {
			replacement.Left, replacement.Right = parentNode.Right.Left, parentNode.Right.Right
			parentNode.Right = replacement
		}
		t.count--
	}
	return t
}

// 移除节点定位，找出节点是父节点的那一边子节点
func (n *Node) findParent(key int) (parentNode *Node, side SIDE) {
	if key == n.Key {
		side = locatedAtItself
		return
	} else if n == nil {
		side = locatedAtNoWhere
		return
	}
	if key < n.Key {
		if parentNode, side = n.Left.findParent(key); side == locatedAtItself {
			parentNode = n
			side = locatedAtLeft
		}
	} else {
		if parentNode, side = n.Right.findParent(key); side == locatedAtItself {
			parentNode = n
			side = locatedAtRight
		}
	}
	return
}

// 取出替代节点，即取出待替换节点右树最大值
func (n *Node) takeOutReplacement(params... interface{}) (replacementNode *Node) {
	if len(params) == 0 {
		// 非递归逻辑（最表层调用）
		params = append(params, true)
		replacementNode = n.Right
		if replacementNode == nil {
			return
		} else if replacementNode.Left == nil {
			n.Right = replacementNode.Right
			return
		}
	} else {
		// 递归逻辑
		replacementNode = n.Left
		if replacementNode == nil {
			return
		} else if replacementNode.Left == nil {
			n.Left = replacementNode.Right
			return
		}
	}
	replacementNode = replacementNode.takeOutReplacement(params...)
	return
}

func (t *Tree) DepthFirstTraverse() {
	if t.root == nil {
		return
	}
	stack := []Node{ *t.root }
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

func (t *Tree) BroadFirstTraverse() {
	if t.root == nil {
		return
	}
	queue := []Node{ *t.root }
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

func (t *Tree) Print() {
	toggle := 1
	ctn := false
	sliceA := make([]*Node, 0)
	sliceB := make([]*Node, 0)
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
					ctn = true
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
					ctn = true
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
				PrintMethodGoBackA:
				if c > 1 {
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
							c--
							goto PrintMethodGoBackA
						}
					}
				}
				if ctn {
					ctn = false
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
					ctn = true
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
					ctn = true
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
				PrintMethodGoBackB:
				if c > 1 {
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
							c--
							goto PrintMethodGoBackB
						}
					}
				}
				if ctn {
					ctn = false
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

func NewTree() *Tree {
	return &Tree{}
}
