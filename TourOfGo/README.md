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
