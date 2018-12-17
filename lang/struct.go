package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
	int64
}

func main() {
	h := &Human{name: "wangxb"}
	fmt.Println(*h, h.name)
	// 我们初始化一个学生
	mark := Student{Human{"Mark", 25, 120}, "Computer Science", 77}
	marker := Student{
		Human: Human{
			name:   "wangxb",
			age:    27,
			weight: 50,
		},
		speciality: "sssss",
		//int64:      45,
	}
	marker.int64 = 89
	fmt.Println("ssssss", &marker.speciality)

	m := new(Student)
	m.age = 1
	fmt.Println(m.age)

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)
}
