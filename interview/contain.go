package main

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。

// 示例 1：
// 输入：s1 = "abb" s2 = "eidbabooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("bab").
// 示例 2：
// 输入：s1= "abb" s2 = "eidbboaoo"
// 输出：false

import (
	"fmt"
	"sort"
)

func checkInclusion(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false // s1 比 s2 长，直接不可能
	}

	// 1. 把 s1 排序（转成切片排序后再转回字符串）
	sortedS1 := sortString(s1)

	// 2. 滑动窗口：遍历 s2 所有长度为 n1 的子串
	for i := 0; i <= n2-n1; i++ {
		// 切出当前窗口的子串（s2[i:i+n1]）
		subStr := s2[i : i+n1]
		// 子串排序后和 sortedS1 对比
		if sortString(subStr) == sortedS1 {
			return true
		}
	}

	return false
}

// 辅助函数：把字符串排序（比如 "bab"→"abb"）
func sortString(s string) string {
	// 字符串转 rune 切片（支持中文，这里虽然用不到，但更通用）
	runes := []rune(s)
	// 排序切片
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// 转回字符串
	return string(runes)
}

func main() {
	fmt.Println(checkInclusion("abb", "eidbabooo")) // 输出 true
	fmt.Println(checkInclusion("abb", "eidbboaoo")) // 输出 false
	fmt.Println(checkInclusion("a", "ab"))          // 输出 true
}
