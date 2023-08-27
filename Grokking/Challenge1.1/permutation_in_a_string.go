package main

import "fmt"

/*
Given a string and a pattern, find out if the string contains any permutation of the pattern.

Permutation is defined as the re-arranging of the characters of the string.
For example, “abc” has the following six permutations:
abc
acb
bac
bca
cab
cba
If a string has ‘n’ distinct characters, it will have n!n! permutations.

Input: String="oidbcaf", Pattern="abc"
Output: true
Explanation: The string contains "bca" which is a permutation of the given pattern.

Input: String="odicf", Pattern="dc"
Output: false
Explanation: No permutation of the pattern is present in the given string as a substring.

Input: String="bcdxabcdy", Pattern="bcdyabcdx"
Output: true
Explanation: Both the string and the pattern are a permutation of each other.

Input: String="aaacb", Pattern="abc"
Output: true
Explanation: The string contains "acb" which is a permutation of the given pattern.
*/

func isSamePattern(s1, s2 map[byte]int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for k, v := range s1 {
		if v != s2[k] {
			return false
		}
	}

	return true
}

func findPermutation(str string, pattern string) bool {
	var (
		windowStart = 0
		charMap     = make(map[byte]int)
		patternMap  = make(map[byte]int)
	)

	// convert pattern to patternMap
	for i := 0; i < len(pattern); i++ {
		patternMap[pattern[i]]++
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {

		charMap[str[windowEnd]]++

		if windowEnd < len(pattern)-1 {
			continue
		}

		if isSamePattern(charMap, patternMap) {
			return true
		}

		delChar := str[windowStart]
		charMap[delChar]--
		if charMap[delChar] == 0 {
			delete(charMap, delChar)
		}
		windowStart++
	}

	return false
}

func main() {
	// Chiamata a findPermutation e gestione dell'output
	str := "oidbcaf"
	pattern := "abc"
	result := findPermutation(str, pattern)

	fmt.Printf("Input: String=\"%s\", Pattern=\"%s\"\n", str, pattern)
	if result {
		fmt.Println("Output: true")
	} else {
		fmt.Println("Output: false")
	}
}
