package ds_test

import (
	"fmt"
	"radix/ds"
	"testing"
)

func TestBSTree(t *testing.T) {
	var d *[]interface{}
	tmp := make([]interface{}, 1)
	d = &tmp
	*d = append(*d, 123)
	fmt.Println(*d)
}

func TestLList(t *testing.T) {
	d := ds.NewLList()
	d.Append(0)
	d.Append(1)
	d.Insert(2, 0)
	d.Insert(3, 1)
	n1, err := d.LLNode(0)
	if n1.Value.(int) != 2 || err != nil {
		t.Fatal("此元素应该为2")
	}
	n2, err := d.LLNode(1)
	if n2.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n3, err := d.LLNode(2)
	if n3.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	n4, err := d.LLNode(3)
	if n4.Value.(int) != 1 || err != nil {
		t.Fatal("此元素应该为1")
	}
	d.Remove(0)
	d.Remove(2)
	n5, err := d.LLNode(0)
	if n5.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n6, err := d.LLNode(1)
	if n6.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	d.Append(4)
	d.Append(5)
	d.Append(6)
	d.Append(7)
	d.Reverse()
	d.Print()
	t.Log("测试成功")
}

func TestDLList(t *testing.T) {
	d := ds.NewDLList()
	d.Append(0)
	d.Append(1)
	d.Insert(2, 0)
	d.Insert(3, -2)
	n1, err := d.DLLNode(0)
	if n1.Value.(int) != 2 || err != nil {
		t.Fatal("此元素应该为2")
	}
	n2, err := d.DLLNode(1)
	if n2.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n3, err := d.DLLNode(-2)
	if n3.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	n4, err := d.DLLNode(-1)
	if n4.Value.(int) != 1 || err != nil {
		t.Fatal("此元素应该为1")
	}
	d.Remove(0)
	d.Remove(-1)
	n5, err := d.DLLNode(0)
	if n5.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n6, err := d.DLLNode(-1)
	if n6.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	d.Print()
	t.Log("测试成功")
}