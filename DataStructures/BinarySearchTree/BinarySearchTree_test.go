package binarysearchtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBinarySearchTree(t *testing.T) {

	number := make([]int, 0)
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		number = append(number, rand.Intn(100000))
	}

	tree := NewTree()
	for _, val := range number {
		tree.Append(val, nil)
	}

	fmt.Printf("rand number: %v \n", number)
	fmt.Printf("print tree:\n%s\n", tree.Sprint())

	removeNode := make(map[int]struct{})
	for {
		rand.Seed(time.Now().UnixNano())
		removeNode[rand.Intn(len(number)-1)] = struct{}{}
		if len(removeNode) >= 5 {
			break
		}
	}
	fmt.Printf("remove node: ")
	for key := range removeNode {
		fmt.Printf("%d ", number[key])
		tree.Remove(number[key])
	}
	fmt.Printf("\nprint tree:\n%s\n", tree.Sprint())
}
