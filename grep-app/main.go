package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing two required positional arguments: fileName  and searchString")
	}
	//os.Args provides access to raw command-line arguments. Note that the first
	//value in this slice is the path to the program, and os.Args[1:] holds the
	//arguments to the program.
	args := os.Args
	fileName := args[1]
	searchTerm := args[2:]
	file, err := os.Open(fileName)
	if err != nil {

		fmt.Printf("failed to open file : %s\n", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on separator
		fmt.Println(scanner.Text()) // token in unicode-char

	}
	fmt.Println(searchTerm)
}
