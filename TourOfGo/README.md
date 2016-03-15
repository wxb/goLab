# Go 语言指南


## 包

Go程序是由包组成的    
程序运行的入口是包 `main`    
按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始。     

## 导入包  

Go语言使用 `import` 导入包，也可以用圆括号组合了导入，”打包”导入     

    import(
        "fmt"
        "math/rand"
    )

## 函数   

* Go语言中函数的定义用：`func`    
* 函数可以没有参数也可以有多个参数；参数变量的类型在**变量名之后**，例如：
```
func add(x int, y int) int {
return x + y
}
```  
* 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略，例如：      
```
    func add(x, y int) int {
    	return x + y
    }
```

## 基本类型   

    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // uint8 的别名

    rune // int32 的别名
     // 代表一个Unicode码

    float32 float64

    complex64 complex128
