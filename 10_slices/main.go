package main

import (
	"fmt"
	"slices"
)

func main() {
	// Uninitialized slice is nil slice
	var nums []int
	fmt.Println("nums:", nums)
	fmt.Println("Is nums nil?", nums == nil)
	fmt.Println("Length of nums:", len(nums))

	var nums2 = make([]int, 2)
	fmt.Println("nums2:", nums2)
	fmt.Println("Is nums2 nil?", nums2 == nil)
	fmt.Println("Length of nums2:", len(nums2))

	var nums3 = make([]int, 2, 5)
	nums3 = append(nums3, 1)
	nums3 = append(nums3, 2)
	nums3 = append(nums3, 3)
	nums3 = append(nums3, 4)
	fmt.Println("nums3:", nums3)
	fmt.Println(cap(nums3))

	nums4 := []int{}
	nums4 = append(nums4, 1)
	fmt.Println("nums4 after append:", nums4)
	fmt.Println("Length of nums4:", len(nums4))
	fmt.Println("Capacity of nums4:", cap(nums4))

	//copy function
	var arr = make([]int, 0, 5)
	arr = append(arr, 2)
	var brr = make([]int, len(arr))
	copy(brr, arr)
	fmt.Println("arr:", arr)
	fmt.Println("brr:", brr)

	// Slice Operations
	slice := []int{10, 20, 30, 40, 50}
	subSlice1 := slice[1:4]
	fmt.Println("subSlice1:", subSlice1)

	subSlice2 := slice[:3]
	fmt.Println("subSlice2:", subSlice2)

	subSlice3 := slice[2:]
	fmt.Println("subSlice3:", subSlice3)

	subSlice4 := slice[:]
	fmt.Println("subSlice4:", subSlice4)

	// Slice Inbuilt Functions
	var ex1 = []int{5, 10, 15, 20, 25}
	var ex2 = []int{30, 35, 40}

	fmt.Println(slices.Equal(ex1, ex2))

	//2D Slice
	var nums2D = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Slice:", nums2D)

}
