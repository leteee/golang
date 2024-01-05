package main

import (
	"fmt"
	mysort "lang-go/sort"
)

func main() {
	arr := []int{5, 4, 3, 2, 1}
	arr = mysort.QuickSort(arr)
	fmt.Println("hello world", arr)

	mysort.TestReverseList()
}
