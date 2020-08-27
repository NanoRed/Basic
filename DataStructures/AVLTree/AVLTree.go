package avltree

// AVL树（平衡二叉搜索树） AVL Tree

// Tree tree structure
type Tree struct {
	root  *Node
	count uint
}

// Node tree node structure
type Node struct {
	Key    int
	Value  interface{}
	Height uint
	Left   *Node
	Right  *Node
}

// Append append a new node to the tree
func (t *Tree) Append(key int, val interface{}) *Tree {
	var appendNode func(int, interface{}, *Node) *Node
	appendNode = func(key int, val interface{}, node *Node) *Node {
		if node == nil {
			node = &Node{
				Key:    key,
				Value:  val,
				Height: 1,
			}
			t.count++
			return node
		} else if key == node.Key {
			node.Value = val
			return node
		}

		if key < node.Key {
			node.Left = appendNode(key, val, node.Left)
			if node.Left.Height == node.Height {
				node.Height++
			}
		} else {
			node.Right = appendNode(key, val, node.Right)
			if node.Right.Height == node.Height {
				node.Height++
			}
		}

		if node.isBalanced() {

		}

		return node
	}
	t.root = appendNode(key, val, t.root)
	return t
}

func (n *Node) isBalanced() bool {
	if n.Left.Height > n.Right.Height {
		return n.Left.Height-n.Right.Height < 2
	}
	return n.Right.Height-n.Left.Height < 2
}
