package sort

func SelectSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if arr[k] > arr[j] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
	return arr
}
