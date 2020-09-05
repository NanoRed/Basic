package avltree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/RedAFD/treeprint"
)

func TestAVLTree(t *testing.T) {
	number := make([]int, 0)
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		number = append(number, rand.Intn(100000))
	}
	tree := NewTree()
	// number = []int{53413, 31067, 29775, 69461, 56657, 61903, 62722, 23300, 81750, 9574, 15534, 26020, 99445, 81178, 59742, 59697, 1785, 45199, 2592, 98034}
	number = []int{11426, 17591, 87028, 47559, 21154, 71726, 27663, 12594, 51009, 97706, 80429, 38555, 19034, 4074, 61225, 27922, 89508, 78150, 93108, 49460}
	number = []int{11426, 17591, 87028, 47559}
	t.Logf("\n%v", number)
	for _, val := range number {
		tree.Append(val, nil)
	}
	t.Logf("\n%s", treeprint.Sprint(tree.root))
}
