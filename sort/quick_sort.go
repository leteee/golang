package sort

func QuickSort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	quickSort(arr, 0, n-1)
	return arr
}

func partition(arr []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= arr[left] {
			j--
		}
		for i < j && arr[i] <= arr[left] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[left] = arr[left], arr[i]
	return i
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}
