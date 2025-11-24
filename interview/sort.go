package main

import "fmt"

func sortWithAsc(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println("----", arr)
}

// BubbleSort 升序冒泡排序
// 冒泡排序核心：从右边开始比较相邻两个数字的大小，再根据结果交换两个数字的位置
func BubbleSort(arr []int) {
	// 边界处理：空数组或单元素数组无需排序
	if len(arr) <= 1 {
		return
	}

	length := len(arr)
	// 新增：交换标志位，标记本轮是否发生交换
	var swapped bool

	for i := 0; i < length; i++ {
		swapped = false // 每轮开始前重置为false
		for j := length - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swapped = true // 发生交换，标记为true
			}
		}
		// 新增：本轮无交换，说明数组已完全有序，直接退出
		if !swapped {
			break
		}
	}
}

// SelectionSortAsc 对整型数组进行升序选择排序
// 选择排序核心：每轮从待排序区找到最小值的索引，仅做一次交换
func SelectionSortAsc(arr []int) {
	// 鲁棒性处理：空数组或单元素数组无需排序，直接返回
	if len(arr) <= 1 {
		return
	}

	arrLen := len(arr)
	// 外层循环：划分已排序区(0~i-1)和待排序区(i~arrLen-1)
	for i := 0; i < arrLen-1; i++ { // 优化：只需遍历到arrLen-1，最后一个元素自然有序
		minIdx := i // 初始化最小值索引为待排序区第一个元素

		// 内层循环：遍历待排序区，找到最小值的索引
		for j := i + 1; j < arrLen; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j // 更新最小值索引，不立即交换
			}
		}

		// 优化：仅当最小值索引不等于i时，才交换（避免无意义的自我交换）
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// InsertSortAsc 对整型数组进行升序插入排序
// 插入排序核心：从右侧的未排序区域内取出一个数据，将它插入到已排序区域内合适的位置上
func InsertSortAsc(arr []int) {

	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	// i从1开始，已排序区域长度默认是i（i=1时已排序长度1，i=2时已排序长度2，和你的sortedLen完全同步）
	for i := 1; i < arrLen; i++ {
		insertVal := arr[i] // 暂存待插入元素，避免多次交换
		j := i - 1          // 已排序区域的最后一个元素索引
		// 比insertVal大的元素统一后移，无需交换
		for j >= 0 && arr[j] > insertVal {
			arr[j+1] = arr[j] // 单赋值操作，比交换轻量
			j--
		}
		arr[j+1] = insertVal // 插入到最终位置
	}
}

func main() {

	s := []int{3, 1, 8, 2, 0, 11}
	InsertSortAsc(s)
	fmt.Println("------->>>>> ", s)
}
