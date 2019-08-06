package main

import "fmt"

func main() {

	container := map[int]string{0: "zero", 1: "one"}

	c, e := interface{}(container).(map[int]string)
	fmt.Println(c, e)

	var srcInt int16 = -255
	dstInt := int8(srcInt)
	fmt.Println(dstInt)
	type sss = int16
	aa := sss(1)
	bb := int16(1)
	fmt.Println(aa == bb, &aa, &bb)
	cc := byte(2)
	dd := uint8(2)
	fmt.Println(cc == dd, &cc, &dd)
	ee := rune(3)
	ff := int32(3)
	fmt.Println(ee == ff, &ee, &ff)
	gg := int(4)
	hh := int32(4)
	fmt.Println(&gg, &hh)
}
