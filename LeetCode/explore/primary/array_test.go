package primary_test

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	f := func(nums []int) int {

		i := 0
		has := map[int]bool{}
		for _, v := range nums {
			if has[v] {
				continue
			}

			has[v] = true
			nums[i] = v
			i++
		}

		return len(has)
	}

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size := f(nums)
	for i := 0; i < size; i++ {
		fmt.Println(nums[i])
	}

	f1 := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}

		slow, fast := 0, 1
		for ; fast < len(nums); fast++ {
			if nums[slow] != nums[fast] {
				slow++
				nums[slow] = nums[fast]
			}
		}

		return slow + 1
	}

	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size1 := f1(nums1)
	for i := 0; i < size1; i++ {
		fmt.Println(nums1[i])
	}
}

func TestMaxProfit(t *testing.T) {

}
