package main

import "fmt"

func mul(nums ...int) {
	fmt.Print(nums, " ")
	a := nums[0]

	for _, num := range nums {
		a *= num
	}
	fmt.Println(a)
}

func main() {

	mul(1, 2)
	mul(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	mul(nums...)
}
