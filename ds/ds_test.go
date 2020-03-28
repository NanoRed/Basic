package ds_test

import (
	"GoRadix/ds"
	"testing"
)

func TestTree(t *testing.T) {
	tr := ds.NewTree()
	tr.Append(nil, 1) // 根
	tr.Append(&[]int{}, 2)
	tr.Append(&[]int{}, 3)
	tr.Append(&[]int{0}, 4)
	tr.Append(&[]int{0}, 5)
	tr.Append(&[]int{0}, 6)
	tr.Append(&[]int{0, 1}, 9)
	tr.Append(&[]int{0, 1, 0}, 13)
	tr.Append(&[]int{0, 1, 0}, 14)
	tr.Append(&[]int{0, 2}, 10)
	tr.Append(&[]int{0, 2}, 11)
	tr.Append(&[]int{1}, 7)
	tr.Append(&[]int{1}, 8)
	tr.Append(&[]int{1, 1}, 12)
	tr.DepthFirstSearch()
}

func TestBSTree(t *testing.T) {
	bst := ds.NewBSTree()
	bst.Append(150, nil)
	bst.Append(102, nil)
	bst.Append(12, nil)
	bst.Append(101, nil)
	bst.Append(3, nil)
	bst.Append(564, nil)
	bst.Append(777, nil)
	bst.Append(666, nil)
	bst.Append(888, nil)
	bst.Append(73, nil)
	bst.Append(66, nil)
	bst.Append(1, nil)
	bst.Append(665, nil)
	bst.Append(664, nil)
	bst.Append(663, nil)
	bst.Remove(73)
	bst.Remove(564)
	bst.Append(662, nil)
	bst.Append(661, nil)
	bst.Print()
}

func TestLList(t *testing.T) {
	d := ds.NewLList()
	d.Append(0)
	d.Append(1)
	d.Insert(2, 0)
	d.Insert(3, 1)
	n1, err := d.Search(0)
	if n1.Value.(int) != 2 || err != nil {
		t.Fatal("此元素应该为2")
	}
	n2, err := d.Search(1)
	if n2.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n3, err := d.Search(2)
	if n3.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	n4, err := d.Search(3)
	if n4.Value.(int) != 1 || err != nil {
		t.Fatal("此元素应该为1")
	}
	d.Remove(0)
	d.Remove(2)
	n5, err := d.Search(0)
	if n5.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n6, err := d.Search(1)
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
	n1, err := d.Search(0)
	if n1.Value.(int) != 2 || err != nil {
		t.Fatal("此元素应该为2")
	}
	n2, err := d.Search(1)
	if n2.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n3, err := d.Search(-2)
	if n3.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	n4, err := d.Search(-1)
	if n4.Value.(int) != 1 || err != nil {
		t.Fatal("此元素应该为1")
	}
	d.Remove(0)
	d.Remove(-1)
	n5, err := d.Search(0)
	if n5.Value.(int) != 3 || err != nil {
		t.Fatal("此元素应该为3")
	}
	n6, err := d.Search(-1)
	if n6.Value.(int) != 0 || err != nil {
		t.Fatal("此元素应该为0")
	}
	d.Print()
	t.Log("测试成功")
}