package main

import "fmt"

/*
Given an array of sorted numbers, remove all duplicates from it.
You should not use any extra space; after removing the duplicates in-place return the length of the subarray
that has no duplicate in it.


Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].


Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].
*/

func remove(arr []int) int {
	newInd := 0
	for curInd := 1; curInd < len(arr); curInd++ {
		if arr[curInd] == arr[newInd] {
			continue
		}
		newInd++
		arr[newInd] = arr[curInd]
	}
	return newInd + 1
}

func main() {
	arr1 := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Printf("Input: %v\n", arr1)
	length1 := remove(arr1)
	fmt.Printf("Output: %d, Array: %v\n", length1, arr1[:length1])

	arr2 := []int{2, 2, 2, 11}
	fmt.Printf("Input: %v\n", arr2)
	length2 := remove(arr2)
	fmt.Printf("Output: %d, Array: %v\n", length2, arr2[:length2])
}
