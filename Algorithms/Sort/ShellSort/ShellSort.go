package ShellSort

// 希尔排序 Shell Sort

func Sort(arr *[]int) {
	count := len(*arr)
	gap := count/2
	for {
		if gap == 0 {
			break
		}
		for i := 0; i < gap; i++ {
			for j := i; j < count; j += gap {
				for k := j; k > 0; k -= gap {
					if k - gap >= 0 && (*arr)[k] < (*arr)[k - gap] { // > for reverse
						(*arr)[k], (*arr)[k - gap] = (*arr)[k - gap], (*arr)[k]
					} else {
						break
					}
				}
			}
		}
		gap = gap/2
	}
}