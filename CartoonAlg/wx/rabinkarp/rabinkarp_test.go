// 漫画：什么是字符串匹配算法？
// https://mp.weixin.qq.com/s/67uf7pRxXh7Iwm7MMpqJoA
package rabinkarp_test

import (
	"fmt"
	"testing"
)

// RK - Rabin-Karp算法
func rabinKarp(str, pattern string) int {
	//主串长度
	m := len(str)
	//模式串的长度
	n := len(pattern)
	//计算模式串的hash值
	patternCode := hash(pattern)
	//计算主串当中第一个和模式串等长的子串hash值
	strCode := hash(str[:n])

	//用模式串的hash值和主串的局部hash值比较。
	//如果匹配，则进行精确比较；如果不匹配，计算主串中相邻子串的hash值。
	for i := 0; i < m-n+1; i++ {
		if strCode == patternCode && compareString(i, str, pattern) {
			return i
		}

		//如果不是最后一轮，更新主串从i到i+n的hash值
		if i < m-n {
			strCode = nextHash(str, strCode, i, n)
		}
	}

	return -1
}

//这里采用最简单的hashcode计算方式：
//把a当做1，把b当中2，把c当中3.....然后按位相加
func hash(str string) (hashcode int32) {
	for _, c := range str {
		hashcode += int32(c) - int32('A')
	}
	return
}

func nextHash(str string, hash int32, index, n int) int32 {
	hash -= int32(str[index]) - int32('A')
	hash += int32(str[index+n]) - int32('A')
	return hash
}

func compareString(i int, str, pattern string) bool {
	subStr := str[i : i+len(pattern)]
	return subStr == pattern
}

func TestMatchSubStr(t *testing.T) {
	str := "aacdesadsdfer"
	pattern := "adsd"
	fmt.Println("第一次出现的位置:", rabinKarp(str, pattern))
}
