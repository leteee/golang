package leetcode

import (
	"fmt"
)

//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。
//比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	res := make([][]byte, numRows)

	i, flag := 0, -1
	for _, v := range []byte(s) {
		res[i] = append(res[i], v)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	result := ""
	for _, v := range res {
		result += string(v)
	}
	return result
}

func Test0006() {
	s := "PAYPALISHIRING"
	ret := convert(s, 3)
	fmt.Printf("%s", ret)
}
