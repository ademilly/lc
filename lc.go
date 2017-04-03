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

func usage() {
	fmt.Println("Usage: lc some_file")
}

func countLines(r io.Reader) int {
	n := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}

	return n
}

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
