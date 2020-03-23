package ds

// 单向链表 Linked List

import (
	"errors"
	"fmt"
)

type LList struct {
	head, tail, current *LLNode
	length uint
}

type LLNode struct {
	Value interface{}
	Next *LLNode
}

func (l *LList) Len() uint {
	return l.length
}

func (l *LList) Reset() {
	l.current = l.head
}

func (l *LList) Range() (*LLNode, bool) {
	r := l.current
	GoOn := true
	if r == nil {
		l.current = l.head
		GoOn = false
	} else if l.current.Next == nil {
		l.current = nil
	} else {
		l.current = l.current.Next
	}
	return r, GoOn
}

func (l *LList) LLNode(index uint) (n *LLNode, err error) {
	if l.length == 0 {
		err = errors.New("the list is empty")
		return
	} else if index >= l.length {
		err = errors.New("out of range")
		return
	}
	switch {
	case index == 0:
		n = l.head
	case index > 0:
		l.Reset()
		for {
			node, GoOn := l.Range()
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
	}
	return
}

func (l *LList) Append(val interface{}) *LList {
	if l.tail == nil {
		l.head = &LLNode{
			Value: val,
		}
		l.tail = l.head
		l.current = l.head
	} else {
		l.tail.Next = &LLNode {
			Value: val,
		}
		l.tail = l.tail.Next
	}
	l.length++
	return l
}

func (l *LList) Insert(val interface{}, index uint) *LList {
	node := &LLNode{
		Value: val,
	}
	switch {
	case index == 0:
		if l.head != nil {
			node.Next = l.head
		}
		l.head = node
		l.current = node
	case index > 0:
		prevLLNode, err := l.LLNode(index - 1)
		if err != nil {
			panic(err)
		}
		node.Next = prevLLNode.Next
		prevLLNode.Next = node
	}
	l.length++
	return l
}

func (l *LList) Remove(index uint) *LList {
	switch {
	case index == 0:
		if l.length == 0 {
			panic(errors.New("the list is empty"))
		}
		l.head = l.head.Next
		l.current = l.head
		if l.head == nil {
			l.tail = nil
		}
	case index > 0:
		prevLLNode, err := l.LLNode(index - 1)
		if err != nil {
			panic(err)
		}
		prevLLNode.Next = prevLLNode.Next.Next
		if prevLLNode.Next == nil {
			l.tail = prevLLNode
		}
	}
	l.length--
	return l
}

func (l *LList) Reverse() *LList {
	l.Reset()
	var prevNode *LLNode
	for {
		node, GoOn := l.Range()
		if GoOn == false {
			break
		}
		node.Next = prevNode
		prevNode = node
	}
	l.head, l.tail = l.tail, l.head
	l.Reset()
	return l
}

func (l *LList) Slice() []interface{} {
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

func (l *LList) Print() {
	fmt.Println(l.Slice())
}

func NewLList() *LList {
	return &LList{}
}