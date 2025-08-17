package main

import (
	"fmt"
	"slices"
)

func binarySearch(arr []int, left int, right int, target int) bool {
	for left <= right {
		mid := left + (right-left)/2
		//fmt.Println(mid)
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	return false

}

func twoSum(nums []int, target int) bool {
	slices.Sort(nums)
	fmt.Println(nums)

	for index, value := range nums {
		complement := target - value
		if binarySearch(nums, index+1, len(nums)-1, complement) {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []int{0, -1, 2, -3, 1}
	target := -2
	fmt.Println(twoSum(arr1, target))

}
