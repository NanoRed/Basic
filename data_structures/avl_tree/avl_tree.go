package avl_tree

// AVL树（平衡二叉搜索树） AVL Tree

type Tree struct {
	root *Node
}

type Node struct {
	Key int
	Value interface{}
	Left *Node
	Right *Node
}

