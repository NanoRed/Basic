package quicksort

// 快速排序 Quick Sort

// Sort sort the data
func Sort(arr *[]int) {
	section := make([][]int, 0)
	section = append(section, []int{1, len(*arr) - 1})
	for {
		current := section[0]
		begin := current[0]
		end := current[1]
		compare := (*arr)[begin-1]
		for {
			if (*arr)[begin] <= compare {
				begin++
			}
			if (*arr)[end] > compare {
				end--
			}
			if end > begin {
				(*arr)[begin], (*arr)[end] = (*arr)[end], (*arr)[begin]
			} else {
				break
			}
		}
		*arr = append((*arr)[current[0]:begin+1], append([]int{compare}, (*arr)[begin+1:current[1]+1]...)...)
		if begin+1-current[0] > 1 {
			section = append(section, []int{current[0] - 1, begin - 1})
		}
		if current[1]+1-begin+1 > 1 {
			section = append(section, []int{begin + 1, current[1]})
		}
		section = section[1:]
		if len(section) == 0 {
			break
		}
	}
}
