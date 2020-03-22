package ds

// 双向链表 Doubly Linked List

import (
	"errors"
	"fmt"
	"math"
)

type Direction int

const (
	DLLFromHead Direction = iota
	DLLFromTail
)

type DLList struct {
	head, tail, current *DLLNode
	length uint
}

type DLLNode struct {
	Value interface{}
	Prev *DLLNode
	Next *DLLNode
}

func (l *DLList) Len() uint {
	return l.length
}

func (l *DLList) Reset() {
	l.current = l.head
}

func (l *DLList) End() {
	l.current = l.tail
}

func (l *DLList) Range(d Direction) (*DLLNode, bool) {
	r := l.current
	GoOn := true
	switch d {
	case DLLFromHead:
		if r == nil {
			l.current = l.head
			GoOn = false
		} else if l.current.Next == nil {
			l.current = nil
		} else {
			l.current = l.current.Next
		}
	case DLLFromTail:
		if r == nil {
			l.current = l.tail
			GoOn = false
		} else if l.current.Prev == nil {
			l.current = nil
		} else {
			l.current = l.current.Prev
		}
	}
	return r, GoOn
}

func (l *DLList) DLLNode(index int) (n *DLLNode, err error) {
	if l.length == 0 {
		err = errors.New("the list is empty")
		return
	} else if index > 0 && index >= int(l.length) {
		err = errors.New("out of range")
		return
	} else if index < 0 {
		abs := uint(math.Abs(float64(index)))
		if abs > l.length {
			err = errors.New("out of range")
			return
		}
	}
	switch {
	case index == 0:
		n = l.head
	case index > 0:
		l.Reset()
		for {
			node, GoOn := l.Range(DLLFromHead)
			if GoOn == false {
				err = errors.New("out of range")
				break
			} else if index == 0 {
				n = node
				l.Reset()
				break
			} else {
				index--
			}
		}
	case index < 0:
		l.End()
		for {
			node, GoOn := l.Range(DLLFromTail)
			if GoOn == false {
				err = errors.New("out of range")
				break
			} else {
				index++
			}
			if index == 0 {
				n = node
				l.Reset()
				break
			}
		}
	}
	return
}

func (l *DLList) Append(val interface{}) *DLList {
	if l.tail == nil {
		l.head = &DLLNode{
			Value: val,
		}
		l.tail = l.head
		l.current = l.head
	} else {
		l.tail.Next = &DLLNode {
			Value: val,
			Prev: l.tail,
		}
		l.tail = l.tail.Next
	}
	l.length++
	return l
}

func (l *DLList) Insert(val interface{}, index int) *DLList {
	node := &DLLNode{
		Value: val,
	}
	switch {
	case index == 0:
		if l.head != nil {
			l.head.Prev = node
			node.Next = l.head
		}
		l.head = node
		l.current = node
	default:
		origDLLNode, err := l.DLLNode(index)
		if err != nil {
			panic(err)
		}
		node.Next = origDLLNode
		node.Prev = origDLLNode.Prev
		if origDLLNode.Prev != nil {
			origDLLNode.Prev.Next = node
		}
		origDLLNode.Prev = node
	}
	l.length++
	return l
}

func (l *DLList) Remove(index int) *DLList {
	rDLLNode, err := l.DLLNode(index)
	if err != nil {
		panic(err)
	}
	if rDLLNode.Prev != nil {
		rDLLNode.Prev.Next = rDLLNode.Next
	} else {
		l.head = rDLLNode.Next
		l.current = rDLLNode.Next
	}
	if rDLLNode.Next != nil {
		rDLLNode.Next.Prev = rDLLNode.Prev
	} else {
		l.tail = rDLLNode.Prev
	}
	l.length--
	return l
}

func (l *DLList) Slice() []interface{} {
	s := make([]interface{}, l.length)
	var i uint
	l.Reset()
	for i = 0; i < l.length; i++ {
		s[i] = l.current.Value
		l.current = l.current.Next
	}
	l.Reset()
	return s
}

func (l *DLList) Print() {
	fmt.Println(l.Slice())
}

func NewDLList() *DLList {
	return &DLList{}
}
