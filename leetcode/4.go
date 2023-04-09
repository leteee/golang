package leetcode

import (
	"fmt"
	"math"
)

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	left, right := 0, m
	medianIdx := (m + n + 1) >> 1
	for left < right {
		i := (left + right + 1) >> 1
		j := medianIdx - i
		nums1Im1 := math.MinInt32
		if i != 0 {
			nums1Im1 = nums1[i-1]
		}
		nums2j := nums2[j]
		if nums1Im1 < nums2j {
			left++
		} else {
			right--
		}
	}
	i, j := left, medianIdx-left
	median := max(If(i == 0, math.MinInt32, nums1[i-1]).(int), nums2[j-1])
	if (m+n)%2 == 0 {
		median2 := min(If(i == m, math.MaxInt32, nums1[i]).(int), nums2[j])
		return float64(median+median2) / 2.0
	} else {
		return float64(median)
	}
}

func Test4() {
	nums1 := []int{1, 3}
	num2 := []int{2}
	ret := findMedianSortedArrays(nums1, num2)
	fmt.Printf("%f", ret)
}
