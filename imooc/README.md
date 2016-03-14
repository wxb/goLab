## Go语言的优势

* 脚本化的语法   
* 静态类型+编译型，程序运行速度有保障     
* 原生支持并发编程：这使得go语言在服务器端软件开发更有优势，降低开发，维护成本；程序可以更好地执行   

## Go语言的劣势   

* go语言并没有Python和ruby那么多语法糖   
* go语言开发的程序运行速度目前还不及C    
* go语言第三方函数库暂时还不像主流的编程语言那么多    


## Go语言的安装与设置  

go语言在Linux、Windows和Mac上的安装参见go语言[官方安装](https://golang.org/doc/install)说明;这里不再赘述    

go语言设置：
* `GOROOT`-go语言的当前安装目录   
    export GOROOT=/usr/local/go   
* `GOPATH`-go语言工作区的集合(go源码文件目录)    
    export GOPATH=~/golib:~/goproject    
* `GOBIN`-存放go编译后的可执行文件目录,    
    export GOBIN=~/gobin   
* `PATH`- 为方便使用Go语言命令和Go程序的可执行文件，需要追加其值到环境变量中    
    export PATH=$PATH:$GOROOT/bin:$GOBIN     

## 工作区和PATH

GOPATH：工作区集合，存放go源码文件    

![工作区结构](1.png)      

src: 用于存放源码文件，以代码包为组织形式    
pkg: 用于存放归档文件(名称以 .a 为后缀的文件)    
![](2.png)    
bin: 存放当前工作区中的Go程序的可执行文件    

## 源码文件的分类和含义   

* Go源码文件，名称以 `.go` 为后缀，内容就是go语言代码；多个文件需要用代码包组织起来的          
* Go源码文件分为三类：命令源码文件，库源码文件，测试源码文件    

> 2-2
