package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readfile(path string) {
	f, err := os.Open(path)

	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "byron") {
			fmt.Printf("your string found on line %s", line)
		}
		line++

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
func main() {
	args := os.Args
	fileName := args[1]

	readfile(fileName)
	//how do I read the second command as a string to parse it in the string.Contains
}
