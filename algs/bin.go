package algs

func BinSearch(arr []int, el int) int {
	var minI, maxI, medI int
	maxI = len(arr) - 1
	for minI = 0; minI <= maxI; {
		medI = ((maxI - minI) / 2) + ((maxI - minI) % 2)
		if arr[medI] > el {
			maxI = medI
		} else if arr[medI] < el {
			minI = medI
		} else {
			return medI
		}
	}
	return medI
}
