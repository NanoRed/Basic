package doublylinkedlist

// 双向链表 Doubly Linked List

import (
	"errors"
	"fmt"
	"math"
)

// Direction list operating direction
type Direction int

const (
	// FromHead do something with the list from head to the tail
	FromHead Direction = iota
	// FromTail same from tail to the head
	FromTail
)

// List list structure
type List struct {
	head, tail, current *Node
	length              uint
}

// Node list node structure
type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

// Len total count of the list nodes
func (l *List) Len() uint {
	return l.length
}

// Reset reset the current node to head node
func (l *List) Reset() {
	l.current = l.head
}

// End set the current node to tail node
func (l *List) End() {
	l.current = l.tail
}

// Range traverse the list, one called return one node
// param d means traversal direction
// it returns a node pointer
// and a boolean data which means whether it reachs to the end
func (l *List) Range(d Direction) (*Node, bool) {
	r := l.current
	ctn := true
	switch d {
	case FromHead:
		if r == nil {
			l.current = l.head
			ctn = false
		} else if l.current.Next == nil {
			l.current = nil
		} else {
			l.current = l.current.Next
		}
	case FromTail:
		if r == nil {
			l.current = l.tail
			ctn = false
		} else if l.current.Prev == nil {
			l.current = nil
		} else {
			l.current = l.current.Prev
		}
	}
	return r, ctn
}

// Search search node from the list with index
func (l *List) Search(index int) (n *Node, err error) {
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
			node, ctn := l.Range(FromHead)
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
	case index < 0:
		l.End()
		for {
			node, ctn := l.Range(FromTail)
			if ctn == false {
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

// Append append a new node to the list
func (l *List) Append(val interface{}) *List {
	if l.tail == nil {
		l.head = &Node{
			Value: val,
		}
		l.tail = l.head
		l.current = l.head
	} else {
		l.tail.Next = &Node{
			Value: val,
			Prev:  l.tail,
		}
		l.tail = l.tail.Next
	}
	l.length++
	return l
}

// Insert insert a node to pecific location of the list
func (l *List) Insert(val interface{}, index int) *List {
	node := &Node{
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
		origNode, err := l.Search(index)
		if err != nil {
			panic(err)
		}
		node.Next = origNode
		node.Prev = origNode.Prev
		if origNode.Prev != nil {
			origNode.Prev.Next = node
		}
		origNode.Prev = node
	}
	l.length++
	return l
}

// Remove remove a pecific node from the list
func (l *List) Remove(index int) *List {
	rNode, err := l.Search(index)
	if err != nil {
		panic(err)
	}
	if rNode.Prev != nil {
		rNode.Prev.Next = rNode.Next
	} else {
		l.head = rNode.Next
		l.current = rNode.Next
	}
	if rNode.Next != nil {
		rNode.Next.Prev = rNode.Prev
	} else {
		l.tail = rNode.Prev
	}
	l.length--
	return l
}

// Slice return the slice format data of the list
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

// Print print the list
func (l *List) Print() {
	fmt.Println(l.Slice())
}

// NewList create a new list object
func NewList() *List {
	return &List{}
}
