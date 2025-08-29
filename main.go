package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Flags
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	filePath := flag.String("file", "", "File to read (if empty, use stdin or args)")
	flag.Parse()

	// Obtener reader según prioridad
	reader, err := getReader(*filePath, flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
		os.Exit(1)
	}

	// Contar
	n, err := count(reader, *lines, *bytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Println(n)
}

func getReader(filePath string, args []string) (io.Reader, error) {
	switch {
	case filePath != "":
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		return file, nil

	case len(args) > 0:
		input := strings.Join(args, " ")
		return strings.NewReader(input), nil

	default:
		info, err := os.Stdin.Stat()
		if err != nil {
			return nil, err
		}
		if (info.Mode() & os.ModeCharDevice) != 0 {
			return nil, fmt.Errorf("no input provided")
		}
		return os.Stdin, nil
	}
}

func count(r io.Reader, countLines, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	n := 0
	for scanner.Scan() {
		n++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return n, nil
}
