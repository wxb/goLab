package main

import "fmt"

// 计算字符串中不包含重复字符的最长子串的长度
func lengthOfLongestSubstring(s string) int {

	right, left := 0, 0
	max := 0
	exists := map[rune]int{}

	for i, c := range s {
		if p, ok := exists[c]; ok {
			left = p + 1
		}

		exists[c] = i
		right = i
		current := right - left + 1
		if current > max {
			max = current
		}
	}

	return max
}

func main() {
	s := "abcabcbb"
	max := lengthOfLongestSubstring(s)

	fmt.Println("最长子串为：", max)
}
