package avltree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/RedAFD/treeprint"
)

func TestAVLTree(t *testing.T) {
	number := make([]int, 0)
	for i := 0; i < 15; i++ {
		rand.Seed(time.Now().UnixNano())
		number = append(number, rand.Intn(100000))
	}
	tree := NewTree()
	for _, val := range number {
		tree.Append(val, nil)
	}
	t.Logf("\n%v", number)
	t.Logf("\n%s", treeprint.Sprint(tree.root))
	t.Logf("\ntotal node:%d", tree.count)

	uniq := make(map[int]struct{})
	removeNode := make([]int, 0)
	for {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(number) - 1)
		if _, ok := uniq[index]; !ok {
			uniq[index] = struct{}{}
			removeNode = append(removeNode, number[index])
			tree.Remove(number[index])
		}
		if len(removeNode) >= 5 {
			break
		}
	}
	t.Logf("\nremoved: %v", removeNode)
	t.Logf("\n%s", treeprint.Sprint(tree.root))
	t.Logf("\ntotal node:%d", tree.count)
}
