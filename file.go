package main

import (
	"fmt"
	"os"
)

func findKthLargest(nums []int, k int) int {
	if nums == nil {
		fmt.Println("Массив пуст")
		os.Exit(0)
	}
	if k > len(nums) {
		fmt.Println("k больше длины массива")
		os.Exit(0)
	}

	n := len(nums)

	if k > n/2 {
		lowLimit := k - 1
		for i := n; i > lowLimit; i-- {
			for j := n - 1; j > n-i; j-- {
				if nums[j] < nums[j-1] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				}
			}
		}
		return nums[n-k]
	} else {
		lowLimit := n - k
		for i := n; i > lowLimit; i-- {
			for j := n - 1; j > n-i; j-- {
				if nums[j] > nums[j-1] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				}
			}
		}
		return nums[k-1]
	}
}
