package linked_list

// 单向链表 Linked List

import (
	"errors"
	"fmt"
)

type List struct {
	head, tail, current *Node
	length uint
}

type Node struct {
	Value interface{}
	Next *Node
}

func (l *List) Len() uint {
	return l.length
}

func (l *List) Reset() {
	l.current = l.head
}

func (l *List) Range() (*Node, bool) {
	r := l.current
	ctn := true
	if r == nil {
		l.current = l.head
		ctn = false
	} else if l.current.Next == nil {
		l.current = nil
	} else {
		l.current = l.current.Next
	}
	return r, ctn
}

func (l *List) Search(index uint) (n *Node, err error) {
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
			node, ctn := l.Range()
			if ctn == false {
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

func (l *List) Append(val interface{}) *List {
	if l.tail == nil {
		l.head = &Node{
			Value: val,
		}
		l.tail = l.head
		l.current = l.head
	} else {
		l.tail.Next = &Node {
			Value: val,
		}
		l.tail = l.tail.Next
	}
	l.length++
	return l
}

func (l *List) Insert(val interface{}, index uint) *List {
	node := &Node{
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
		prevNode, err := l.Search(index - 1)
		if err != nil {
			panic(err)
		}
		node.Next = prevNode.Next
		prevNode.Next = node
	}
	l.length++
	return l
}

func (l *List) Remove(index uint) *List {
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
		prevNode, err := l.Search(index - 1)
		if err != nil {
			panic(err)
		}
		prevNode.Next = prevNode.Next.Next
		if prevNode.Next == nil {
			l.tail = prevNode
		}
	}
	l.length--
	return l
}

func (l *List) Reverse() *List {
	l.Reset()
	var prevNode *Node
	for {
		node, ctn := l.Range()
		if ctn == false {
			break
		}
		node.Next = prevNode
		prevNode = node
	}
	l.head, l.tail = l.tail, l.head
	l.Reset()
	return l
}

func (l *List) Slice() []interface{} {
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

func (l *List) Print() {
	fmt.Println(l.Slice())
}

func NewList() *List {
	return &List{}
}