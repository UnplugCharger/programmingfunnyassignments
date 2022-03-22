package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Palindrome Checker")
	var word string
	fmt.Scanln(&word)

	isPalindrome(word)
}

func isPalindrome(s string) string {

	if len(os.Args) < 1 {
		log.Println("Please enter a valid string")
	}
	reversedStr := ""

	// loop over the string from the last index and append the letters to the empty reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversedStr += string(s[i])
	}
	for i := range s {
		if s[i] != reversedStr[i] {
			return "is not palindrome"
		}
	}

	return fmt.Println("is palindrome")
}
