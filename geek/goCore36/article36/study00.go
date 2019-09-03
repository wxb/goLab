package main

import "fmt"

func main() {
	name := "王晓勃[wangxiaobo]🏷"
	fmt.Printf("%q \n", name)
	fmt.Printf("rune(char): %q \n", []rune(name))
	fmt.Printf("rune(hex): %x \n", []rune(name))

	// 注意：len得到的是字节长度，在对多字节字符使用时需注意
	fmt.Println(name, len(name))

	// 对于多字节字符在 range 时按照rune类型输出，注意i键的值在多字符的跨越
	for i, c := range name {
		fmt.Printf("%d: %q \n", i, c)
	}

	// 为防止range隐晦的rune类型规则，建议明确使用[]rune对字符串进行转换，此时i键的值也变成了我们通常想到的样子
	for i, c := range []rune(name) {
		fmt.Printf("%d, %v \n", i, c)
	}
}
