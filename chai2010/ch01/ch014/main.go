package main

import "fmt"

type File struct {
	abc string
}

// 关闭文件
func (f *File) Close() error {
	// ...
	fmt.Println(f.abc)
	return nil
}

// 读文件数据
func (f *File) Read(offset int64, data []byte) int {
	// ...

	return 0
}

func main() {
	var CloseFile = (*File).Close

	// 不依赖具体的文件对象
	// func ReadFile(f *File, offset int64, data []byte) int
	var ReadFile = (*File).Read

	// 文件处理
	// f, _ := OpenFile("foo.dat")

	// 通过叫方法表达式的特性可以将方法还原为普通类型的函数：
	f := File{"123"}
	ReadFile(&f, 0, []byte("abc"))
	CloseFile(&f)

}
