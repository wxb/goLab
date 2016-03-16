# Go 语言指南


## 包

Go程序是由包组成的    
程序运行的入口是包 `main`    
按照惯例，包名与导入路径的最后一个目录一致。例如，`"math/rand"` 包由 package rand 语句开始。     

## 导入包  

Go语言使用 `import` 导入包，也可以用圆括号组合了导入，”打包”导入     
```go
    import(
        "fmt"
        "math/rand"
    )
```

## 变量
* Go语言使用 `var` 定义一个变量列表，变量类型在变量后面,`var` 语句可以定义在包或函数级别。        
```go
    package main

    import "fmt"

    var c, python, java bool

    func main() {
    	var i int
    	fmt.Println(i, c, python, java)
    }
```
* 初始化变量：变量定义可以包含初始值，每个变量对应一个。如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。    
```go
    package main

    import "fmt"

    var i, j int = 1, 2

    func main() {
    	var c, python, java = true, false, "no!"
    	fmt.Println(i, j, c, python, java)
    }
```
* 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。`:=` 结构不能使用在函数外。   
```go
    func main() {
    	var i, j int = 1, 2
    	k := 3
    	c, python, java := true, false, "no!"

    	fmt.Println(i, j, k, c, python, java)
    }
```

## 函数/方法   

* Go语言中函数的定义用：`func`    
* 函数可以没有参数也可以有多个参数；参数变量的类型在**变量名之后**，例如：
```go
    func add(x int, y int) int {
        return x + y
    }
```  
* 当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略，例如：      
```go
    func add(x, y int) int {
    	return x + y
    }
```   
* Go函数可以返回任意数量的返回值，例如：    
```go   
    func swap(x, y string)(string, string){
        return "Go from", "Google"
    }
```
* 命名返回值：Go语言中方法的返回值可以被命名，并且可以像变量一样使用；**没有参数的 return 语句返回结果的当前值。也就是`直接`返回。** 例如：     
```go
    package main

    import "fmt"

    func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
    }

    func turn(i, j int)(ivalue, jvalue int){
    i = i+i
    j = j-j
    return j, i
    }

    func main() {
    fmt.Println(split(17))
    fmt.Println(turn(1, 2))
    }
```

## 基本类型   
```go
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // uint8 的别名

    rune // int32 的别名
     // 代表一个Unicode码

    float32 float64

    complex64 complex128
```

## 类型转换

表达式 T(v) 将值 v 转换为类型 `T`。    
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
// 或者更加简洁
i := 42
f := float64(i)
u := uint(f)

```

## 常量

* 常量的定义与变量类似，只不过使用 const 关键字。      
* 常量可以是字符、字符串、布尔或数字类型的值。   
* 常量不能使用 := 语法定义。    
```go
    const Pi = 3.14
    const World = "世界"
    const Truth = true
```

# 流程控制

## for
* Go语言**只有一种**循环结构-`for`循环；相对于C,php,java等语言来说，Go的`for`循环只是没有了`()`,其他用法都是一样的。     
* Go语言中`for`跟 C 或者 Java 中一样，可以让前置、后置语句为空。例如：    
```go
    func main() {
    	sum := 1
    	for ; sum < 1000; {
    		sum+=sum
    		fmt.Println(sum)
    	}
    	fmt.Println(sum)
    }
```
* Go语言中只有一种循环结构`for`，并没有`while`,`do-while`类似的结构，但是可以通过一种变形的`for`达到同样的效果：     
```go
    func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
    }
```
