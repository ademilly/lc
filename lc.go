// lc is a ligne counter command line utilitary
// usage: lc [pathtofile]
// Output: number of lines

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// usage print usage string to standard output
func usage() {
	fmt.Println("Usage: lc some_file")
}

// countLines scan r io.Reader and count number of lines in r
// returns number of lines in r as an int
func countLines(r io.Reader) int {
	n := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}

	return n
}

// main checks command line arguments and compute line count from them
func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(countLines(f))
}
