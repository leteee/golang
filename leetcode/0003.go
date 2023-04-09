package leetcode

import "fmt"

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func lengthOfLongestSubstring(s string) int {
	//哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	//右指针，初始值为-1，相当于我们在字符串的左边界的左侧，还没开始移动
	right, ans := -1, 0
	for left := 0; left < n; left++ {
		if left != 0 {
			delete(m, s[left-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}
		ans = max(ans, right-left+1)

	}
	return ans
}

func Test3() {
	s := "abcabcbb"
	ret := lengthOfLongestSubstring(s)
	fmt.Printf("%d", ret)
}
