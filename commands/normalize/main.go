package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// tr '[:upper:]' '[:lower:]' |
// sed -e 's/^/__label__/g' |
// sed -e "s/'/ ' /g"
// 	-e 's/"//g'
// 	-e 's/\./ \. /g'
// 	-e 's/<br \/>/ /g'
// 	-e 's/,/ , /g'
// 	-e 's/(/ ( /g'
// 	-e 's/)/ ) /g'
// 	-e 's/\!/ \! /g'
// 	-e 's/\?/ \? /g'
// 	-e 's/\;/ /g'
// 	-e 's/\:/ /g'
// tr -s " " | myshuf

func main() {
	label := flag.Bool("l", true, `If true, adds the line prefix "__label__"; use -l=0 for prediction input`)
	flag.Parse()

	space := regexp.MustCompile(`\s+`)

	// `ABC'''"""...,,,()!?;:    END`
	// normalized: `__label__ABC ' ' ' . . . , , , ( ) ! ? END`

	file := os.Stdin
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		if *label {
			line = "__label__" + line
		}
		line = strings.ReplaceAll(line, "'", " ' ")
		line = strings.ReplaceAll(line, `"`, "")
		line = strings.ReplaceAll(line, ".", " . ")
		line = strings.ReplaceAll(line, "<br />", " ")
		line = strings.ReplaceAll(line, ",", " , ")
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")
		line = strings.ReplaceAll(line, "!", " ! ")
		line = strings.ReplaceAll(line, "?", " ? ")
		line = strings.ReplaceAll(line, ";", " ")
		line = strings.ReplaceAll(line, ":", " ")
		line = space.ReplaceAllString(line, " ")
		line = strings.ReplaceAll(line, "0", "@")
		line = strings.ReplaceAll(line, "1", "@")
		line = strings.ReplaceAll(line, "2", "@")
		line = strings.ReplaceAll(line, "3", "@")
		line = strings.ReplaceAll(line, "4", "@")
		line = strings.ReplaceAll(line, "5", "@")
		line = strings.ReplaceAll(line, "6", "@")
		line = strings.ReplaceAll(line, "7", "@")
		line = strings.ReplaceAll(line, "8", "@")
		line = strings.ReplaceAll(line, "9", "@")
		// line = strings.ToValidUTF8(line, "")
		fmt.Println(line)
	}
}
