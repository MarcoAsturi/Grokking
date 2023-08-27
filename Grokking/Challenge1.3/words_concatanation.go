package main

import "fmt"

/*
Given a string and a list of words,
find all the starting indices of substrings in the given string that are a concatenation of all the given words
exactly once without any overlapping of words. It is given that all words are of the same length.

Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".

Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".
*/

func findWordConcatenation(str string, words []string) []int {
	if len(words) == 0 || len(str) == 0 {
		return []int{}
	}
	var (
		wordFreqMap = make(map[string]int)
		wordsCount  = len(words)
		wordLength  = len(words[0])
		startIndex  = make([]int, 0, 10) // store final return result
	)

	// Keep the frequency of every word in a Map.
	for _, v := range words {
		wordFreqMap[v]++
	}

	for i := 0; i <= len(str)-wordLength*wordsCount; i++ {
		wordsSeenMap := make(map[string]int)

		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength
			word := str[nextWordIndex : nextWordIndex+wordLength]

			if _, ok := wordFreqMap[word]; !ok {
				break
			}
			wordsSeenMap[word]++
			if wordsSeenMap[word] > wordFreqMap[word] {
				break
			}

			if j == wordsCount-1 {
				startIndex = append(startIndex, i)
			}
		}
	}

	return startIndex
}

func main() {
	str := "catfoxcat"
	words := []string{"cat", "fox"}
	result := findWordConcatenation(str, words)

	fmt.Printf("Input: String=\"%s\", Words=%v\n", str, words)
	fmt.Printf("Output: %v\n", result)
}
