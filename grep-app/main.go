package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//os.Args provides access to raw command-line arguments. Note that the first
	//value in this slice is the path to the program, and os.Args[1:] holds the
	//arguments to the program.
	args := os.Args

	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println(scanner.Text()) // token in unicode-char

	}
}
