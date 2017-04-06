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

// source return the source for countLines depending on os.Args[1] value
func source(arg string) io.Reader {
	switch arg {
	case "-":
		return os.Stdin
	default:
		f, err := os.Open(arg)
		if err != nil {
			log.Fatalln(err)
		}
		return f
	}
}

// countLines scan r io.Reader and count number of lines in r
// returns number of lines in r as an int
func countLines(r io.Reader) int64 {
	n := int64(0)

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

	acc := int64(0)
	for _, arg := range os.Args[1:] {
		acc += countLines(source(arg))
	}
	fmt.Printf("%v\n", acc)
}
