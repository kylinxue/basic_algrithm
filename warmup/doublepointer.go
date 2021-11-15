package warmup

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	n := len(s)
	i, j := n-1, n-1
	res := make([]string, 0)

	for ; i >= 0; i-- {
		// i为空格，i+1不为空格，找到一个单词
		if i < n-1 && s[i] == ' ' && s[i+1] != ' ' {
			res = append(res, s[i+1:j+1])
		} else if i == 0 && s[i] != ' ' { // i=0，i不为空格，找到一个单词
			res = append(res, s[i:j+1])
		}

		// 通过i判断单词尾部
		if i > 0 && s[i] == ' ' && s[i-1] != ' ' {
			j = i - 1
		}
	}

	return strings.Join(res, " ")
}
