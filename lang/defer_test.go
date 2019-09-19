package lang_test

import (
	"fmt"
	"testing"
)

// return 不是原子的语句：
// 	实际执行分为:设置返回值 --> ret，defer语句实际执行在返回前(ret)，即拥有defer的函数返回过程是：设置返回值-->执行defer-->ret
// 	有一个事实必须要了解，关键字return不是一个原子操作，实际上return只代理汇编指令ret，即将跳转程序执行。
// 	比如语句return i，实际上分两步进行，即将i值存入栈中作为返回值，然后执行跳转，而defer的执行时机正是跳转前，所以说defer执行时还是有机会操作返回值的。

// defer规则:
// 	规则一：延迟函数的参数在defer语句出现时就已经确定下来了
// 	规则二：延迟函数执行按后进先出顺序执行，即先出现的defer最后执行
// 	规则三：延迟函数可能操作主函数的具名返回值

func TestDeferParams(t *testing.T) {
	a := 1
	// 作为函数参数，则在defer定义时就把值传递给defer，并被cache起来；
	defer func(v int) {
		fmt.Println("second step:", v)
	}(a)

	a++
	fmt.Println("first step:", a)

	// Output:
	// first step: 2
	// second step: 1
}

func TestDeferClosure(t *testing.T) {
	a := 1
	// 作为闭包引用的话，则会在defer函数真正调用时根据整个上下文确定当前的值
	defer func() {
		fmt.Println("second step:", a)
	}()

	a++
	fmt.Println("first step:", a)

	// Output:
	// first step: 2
	// second step: 2
}

type number int

func (n number) print() {
	fmt.Println(n)
}
func (n *number) pprint() {
	fmt.Println(*n)
}
func TestDeferValueAndPointer(t *testing.T) {
	var n number

	defer n.print()               // 调用n值方法，defer会保存n此时值
	defer n.pprint()              // 调用n指针方法，defer会保存n此时的引用，so在函数退出调用时, n内存地址中保存的值变化了
	defer (&n).pprint()           // 同上
	defer func() { n.print() }()  // 闭包n:3, 然后调用n为3的值方法打印出 3
	defer func() { n.pprint() }() // 闭包n:3, 然后调用n为3的指针方法打印出 3

	n = 3
}

// --------------------------------------------------------------------------------
// 函数中：
// 	return xxx
// 进过编译后，变成了三条指令：
// 	返回值 = xxx
// 	调用 defer 函数
// 	空return

func TestDeferReturn01(t *testing.T) {

	f0 := func() (r int) {
		t := 5
		defer func() {
			t = t + 5
		}()
		return t
	}

	t.Log(f0())
	// Output：5

	//  相当于
	f1 := func() (r int) {
		t := 5

		// ...

		// 1. 赋值指令
		r = 5

		// 2. defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		defer func() {
			t = t + 5
		}()

		// 3. 空的return指令
		return
	}

	t.Log(f1())
	// Output：5
}

func TestDeferReturn02(t *testing.T) {
	f0 := func() (r []int) {
		t := make([]int, 1, 5)
		t[0] = 1

		defer func() {
			t = append(t, 4)
		}()

		return t
	}
	t.Log(f0()) // [1]

	f00 := func() (r []int) {
		t := make([]int, 1, 5)
		t[0] = 1

		r = t

		func() {
			// 对于引用类型切片，虽然切片结构中指向底层数组的指针值都相同，但是append后的切片t已经发生了变化，已经不是原来的t了，不同的就是len这个属性，其他pointer和cap都没变
			t = append(t, 4)
		}()

		fmt.Println(len(r), cap(r), r)
		fmt.Println(len(t), cap(t), t)
		fmt.Printf("%p %p\n", t, r)
		return
	}
	t.Log(f00()) // [1]

	f1 := func() (r []int) {
		t := make([]int, 2, 5)
		t[0] = 1

		defer func() {
			t[1] = 4
		}()

		return t
	}
	t.Log(f1()) // [1 4]
	f11 := func() (r []int) {
		t := make([]int, 2, 5)
		t[0] = 1

		r = t

		func() {
			// t这个切片pointer、len、cap都没有发生变化；只是对pointer指向的内存的+1位置赋值了一个4；so r也跟随变化了
			t[1] = 4
		}()

		fmt.Println(len(r), cap(r))
		fmt.Printf("%p %p\n", t, r)
		return
	}
	t.Log(f11()) // [1 4]

}

func TestDeferFuncReturn(t *testing.T) {
	f := func() (result int) {
		i := 1

		defer func() {
			result++
			fmt.Println(result)
		}()

		return i
	}

	fmt.Println(f())

	f1 := func() (result int) {
		i := 1

		result = i

		func() {
			result++
			fmt.Println(result)
		}()

		return
	}

	fmt.Println(f1())
}
