package algs

func SelectSort(arr []int) []int {
	sortArr := make([]int, len(arr))
	var max int = arr[0]
	for j, _ := range sortArr {
		for i, _ := range arr {
			if arr[i] > max {
				max = arr[i]
			}
		}
		sortArr[j] = max
	}
	return sortArr
}
