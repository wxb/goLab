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

func BubbleSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := length - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

}

func main() {

	s := []int{3, 1, 8, 2, 0, 11}
	sortWithAsc(s)
}
