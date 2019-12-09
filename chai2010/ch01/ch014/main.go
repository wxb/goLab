package main

import (
	"errors"
	"fmt"
)

// File 文件对象
type File struct {
	Path string
}

// Close 关闭文件
func (f *File) Close() error {
	defer fmt.Println("Close: ", f.Path)

	return nil
}

// Read 读文件数据
func (f *File) Read(offset int64, data []byte) int {
	defer fmt.Println("Read: ", f.Path)

	return len(data)
}

func (f *File) Write(data []byte) error {
	if len(data) > 1024 {
		return errors.New("Out of write range 1024")
	}

	return nil
}

// OpenFile 打开文件
func OpenFile(name string) (f *File, err error) {
	return &File{Path: name}, nil
}

func main() {
	// f, _ := OpenFile("/path/prod.ini")
	// f.Read(1, []byte("abc"))
	// f.Close()

	var CloseFile = (*File).Close

	// 不依赖具体的文件对象
	// func ReadFile(f *File, offset int64, data []byte) int
	var ReadFile = (*File).Read

	// 文件处理
	f, _ := OpenFile("/path/prod.ini")
	ReadFile(f, 0, []byte("abc"))
	CloseFile(f)

	// 通过叫方法表达式的特性可以将方法还原为普通类型的函数：
	ff := File{"/path/test.ini"}
	ReadFile(&ff, 0, []byte("aabbcc"))
	CloseFile(&ff)
}
