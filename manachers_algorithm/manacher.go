package manachers_algorithm

import (
	"strings"
)

/**
	1. Preprocessing the string
  2. Iterating N times and expanding from the center
	3. Picking the next center
*/

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	// Preprocess the palindrome so that the string will have a length of 2 * N + 1 for the algorithm to work
	var testString string
	for _, v := range s {
		testString += "#" + string(v)
	}
	testString += "#"
	// will contain the lengths of the palindromes
	P := make([]int, len(testString))

	pos, maxRight := 0, 0
	center, maxLen := 0, 0
	for i := range P {
		if i < maxRight {
			P[i] = Min(P[2*pos-i], maxRight-i)
		} else {
			P[i] = 1
		}
		for i-P[i] >= 0 && i+P[i] < len(testString) && testString[i-P[i]] == testString[i+P[i]] {
			P[i]++
		}
		if P[i]+i-1 > maxRight {
			maxRight = P[i] + i - 1
			pos = i
		}
		if maxLen < P[i]-1 {
			maxLen = P[i] - 1
			center = i
		}
	}

	return strings.Replace(testString[center-maxLen:center+maxLen], "#", "", -1)
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
