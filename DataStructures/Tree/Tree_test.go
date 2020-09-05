package tree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/RedAFD/treeprint"
)

func TestTree(t *testing.T) {
	tree := NewTree()
	rand.Seed(time.Now().UnixNano())
	getRandValue := func() int {
		return rand.Intn(100000)
	}
	// root
	err := tree.Append([]int{}, getRandValue())
	if err != nil {
		t.Fatal(err)
	}
	// layer 2
	for i := 0; i < 5; i++ {
		err := tree.Append([]int{i}, getRandValue())
		if err != nil {
			t.Fatal(err)
		}
	}
	// layer 3
	layer2Index := rand.Intn(5)
	for i := 0; i < 5; i++ {
		err := tree.Append([]int{layer2Index, i}, getRandValue())
		if err != nil {
			t.Fatal(err)
		}
	}
	// layer 4
	layer3Index := rand.Intn(5)
	for i := 0; i < 5; i++ {
		err := tree.Append([]int{layer2Index, layer3Index, i}, getRandValue())
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("\n%s", treeprint.Sprint(tree.root))
	// search
	index := []int{layer2Index, layer3Index, rand.Intn(5)}
	node, err := tree.Search(index)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%v %v", index, node.Value)
	// remove
	err = tree.Remove([]int{1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n%s", treeprint.Sprint(tree.root))
}
