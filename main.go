package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")
	// Parsing the flags provided by the user
	flag.Parse()
	// Calling the count funtion to count the number of words
	// received from the Standar Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader(such as files)
	scanner := bufio.NewScanner(r)
	// If countBytes is set then we split the scanner by bytes,
	// el if the count lines is not set, we want to count words so we define
	// the scanner split types to words (default is split by lines)
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	// Defining a counter
	wc := 0
	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}
	// Return the total
	return wc
}
