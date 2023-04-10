package leetcode

import (
	"fmt"
	"math"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围[−231, 231− 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
//输入：x = -123
//输出：-321

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func Test0007() {
	i := -123
	ret := reverse(i)
	fmt.Printf("%d", ret)
}
