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
	"path"
	"strings"
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

// handleArgs handles command line arguments
// returns a slice of strings where
// - head contains behavior tags
// - tails is the list of files to count lines from
func handleArgs(cmdArgs []string) []string {
	args := []string{
		"",
	}

	for _, in := range cmdArgs {
		if in != "-d" {
			args = append(args, in)
		} else {
			args[0] += "DETAILS|"
		}
	}

	return args
}

// main checks command line arguments and compute line count from them
func main() {

	if len(os.Args) < 2 {
		usage()
		return
	}

	args := handleArgs(os.Args[1:])
	acc := make([]chan int64, len(args[1:]))

	for i, arg := range args[1:] {
		acc[i] = make(chan int64)
		go countLines(source(arg), acc[i])
	}

	count := int64(0)
	for i := range acc {
		for c := range acc[i] {
			count += c

			if strings.Contains(args[0], "DETAILS") {
				fmt.Println(c, "\t", path.Base(args[1+i]))
			}
		}
	}
	fmt.Printf("%v\n", count)
}
