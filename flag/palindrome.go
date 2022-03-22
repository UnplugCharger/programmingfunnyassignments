package main

import (
	"fmt"
)

func main() {

	fmt.Println("Palindrome Checker")
	var word string
	fmt.Scanln(&word)

	fmt.Println(isPalindrome(word))
}

func isPalindrome(s string) bool {

	reversedStr := ""

	// loop over the string from the last index and append the letters to the empty reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversedStr += string(s[i])
	}
	for i := range s {
		if s[i] != reversedStr[i] {
			return false
		}
	}

	return true
}
