package rotate

import "fmt"

func Rotate(nums []int, k int) []int {
	n := len(nums)


	// Rotate the array
	rotateArray := make([]int, n)
	copy(rotateArray, nums[n-k:])
	copy(rotateArray[k:], nums[:n-k])

	fmt.Println(rotateArray)

	return rotateArray
}

