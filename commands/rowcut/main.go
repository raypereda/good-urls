// This command cuts out one column from a csv.
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var columnNum int
	flag.IntVar(&columnNum, "c", 1, "the column to cut out of the CSV file")
	usage := func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "    rowcut [-c=N] <file>\n")
		fmt.Fprintf(flag.CommandLine.Output(), "<file> is required.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Cuts out one(1) column from csv.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Output is set to standard output.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Example:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "$ cat file.csv\n")
		fmt.Fprintf(flag.CommandLine.Output(), "AAA,111,ABC\n")
		fmt.Fprintf(flag.CommandLine.Output(), "BBB,222,MNO\n")
		fmt.Fprintf(flag.CommandLine.Output(), "CCC,333,XYZ\n")
		fmt.Fprintf(flag.CommandLine.Output(), "$ rowcut -c=2 file.csv\n")
		fmt.Fprintf(flag.CommandLine.Output(), "111\n")
		fmt.Fprintf(flag.CommandLine.Output(), "222\n")
		fmt.Fprintf(flag.CommandLine.Output(), "333\n")
		flag.PrintDefaults()
	}
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintf(flag.CommandLine.Output(), "Input file is missing from command-line.\n")
		usage()
		os.Exit(1)
	}

	csvFile, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bufio.NewReader(csvFile))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record[columnNum-1])
	}
}
