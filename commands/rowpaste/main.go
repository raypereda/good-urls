// This command pastes two CSV files row wise.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	usage := func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "    rowpaste <file1> <file2>\n")
		fmt.Fprintf(flag.CommandLine.Output(), "<file1> and <file2> are a required.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Paste together CSV files row wise.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Output is set to standard output.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Example:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "$ cat file1.txt\n")
		fmt.Fprintf(flag.CommandLine.Output(), "AAA\n")
		fmt.Fprintf(flag.CommandLine.Output(), "BBB\n")
		fmt.Fprintf(flag.CommandLine.Output(), "CCC\n")
		fmt.Fprintf(flag.CommandLine.Output(), "$ cat file2.txt\n")
		fmt.Fprintf(flag.CommandLine.Output(), "111\n")
		fmt.Fprintf(flag.CommandLine.Output(), "222\n")
		fmt.Fprintf(flag.CommandLine.Output(), "333\n")
		fmt.Fprintf(flag.CommandLine.Output(), "$ rowpaste file1.txt file2.txt\n")
		fmt.Fprintf(flag.CommandLine.Output(), "AAA,111\n")
		fmt.Fprintf(flag.CommandLine.Output(), "BBB,222\n")
		fmt.Fprintf(flag.CommandLine.Output(), "CCC,333\n")
		flag.PrintDefaults()
	}
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) != 3 {
		fmt.Fprintf(flag.CommandLine.Output(), "Input file(s) is missing from command-line.\n")
		usage()
		os.Exit(1)
	}

	inputfile1 := os.Args[len(os.Args)-2]
	lines1, err := readLines(inputfile1)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	inputfile2 := os.Args[len(os.Args)-1]
	lines2, err := readLines(inputfile2)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	if len(lines1) != len(lines2) {
		log.Fatalf("File lengths are not equal: %d != %d", len(lines1), len(lines2))
	}

	for i := range lines1 {
		fmt.Printf("%s,%s\n", lines1[i], lines2[i])
	}
}
