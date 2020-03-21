package main

/*
	the goal is to read in stdin, and read in an argument, and then add all the lines
	from stdin to an array and add all the lines from the target file to an array, then
	merge them together, then sort them, then write them back out to the file
	represented by the program argument.
*/

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)

	var lines = make(map[string]bool)
	var fileName = os.Args[1]
	var file, _ = os.Open(fileName)

	dedupe(bufio.NewScanner(os.Stdin), lines)
	dedupe(bufio.NewScanner(file), lines)

	keys := make([]string, len(lines))
	i := 0
	for k := range lines {
		if k!="" {
			keys[i] = k
			i++
		}
	}

	join := strings.Join(keys, "\n")
	_ = ioutil.WriteFile(fileName, []byte( join ), 0644)
}

func dedupe(scanner *bufio.Scanner, lines map[string]bool) {
	for scanner.Scan() {
		text := scanner.Text()
		if !lines[text] {
			lines[text] = true
		}
	}
}
