package sort

func InsertSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	for i := 1; i < n; i++ {
		base := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > base; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = base
	}
	return arr
}
