package algs

func SelectSort(arr []int) []int {
	for j, _ := range arr {
		max := arr[j]
		index := j
		for i := j + 1; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
				index = i
			}
		}
		arr[j], arr[index] = arr[index], arr[j]
	}
	return arr
}
