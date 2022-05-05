package algs

func SelectSort(arr []int) []int {
	var max, index int = arr[0], 0
	for j, _ := range arr {
		for i := j; i < len(arr)-1; i++ {
			if arr[i] > max {
				index = i
			}
		}
		arr[j], arr[index] = arr[index], arr[j]
	}
	return arr
}
