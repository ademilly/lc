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
func countLines(r io.Reader, acc chan int64) {
	n := int64(0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}

	acc <- n
	close(acc)
}

// main checks command line arguments and compute line count from them
func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	acc := make([]chan int64, 0)

	for i, arg := range os.Args[1:] {

		acc = append(acc, make(chan int64))
		go countLines(source(arg), acc[i])
	}

	count := int64(0)
	for i := range acc {
		for c := range acc[i] {
			count += c
		}
	}
	fmt.Printf("%v\n", count)
}
