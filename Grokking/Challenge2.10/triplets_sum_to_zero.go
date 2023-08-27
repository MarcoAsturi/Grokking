package main

import (
	"fmt"
	"sort"
)

/*
Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.
*/

func searchTriplets(arr []int) [][3]int {
	// 1. sort array
	sort.Ints(arr)
	triplets := make([][3]int, 0)

	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		triplets = searchPair(arr, -arr[i], i+1, triplets)
	}

	return triplets
}

func searchPair(arr []int, targetNum int, left int, triplets [][3]int) [][3]int {
	right := len(arr) - 1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetNum {
			triplets = append(triplets, [3]int{-targetNum, arr[left], arr[right]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if currentSum > targetNum {
			right--
		} else {
			left++
		}
	}
	return triplets
}

func main() {
	arr1 := []int{-3, 0, 1, 2, -1, 1, -2}
	fmt.Printf("Input: %v\n", arr1)
	triplets1 := searchTriplets(arr1)
	fmt.Println("Output:")
	for _, triplet := range triplets1 {
		fmt.Println(triplet)
	}

	arr2 := []int{-5, 2, -1, -2, 3}
	fmt.Printf("Input: %v\n", arr2)
	triplets2 := searchTriplets(arr2)
	fmt.Println("Output:")
	for _, triplet := range triplets2 {
		fmt.Println(triplet)
	}
}
