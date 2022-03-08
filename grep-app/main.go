package main

import (
	"fmt"
	"io/ioutil"
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
	b, err := ioutil.ReadAll(file)
	fmt.Println(b)
}
