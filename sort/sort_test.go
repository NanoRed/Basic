package sort_test

import (
	"GoRadix/sort"
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := []int{57, 1, 997, 325, 25, 77, 123, 98, 23, 88, 32}
	sort.ShellSort(&arr)
	fmt.Println(arr)
}
