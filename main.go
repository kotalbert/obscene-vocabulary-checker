package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fn := readUserInput() // name of file to read
	ws := readLines(fn)   // list of taboo words
	tw := readUserInput() // word to check if taboo

	ts := buildTabooSet(ws)
	if _, ok := ts[tw]; ok {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}

func buildTabooSet(ws []string) map[string]struct{} {
	set := make(map[string]struct{}, len(ws))
	for _, w := range ws {
		set[w] = struct{}{}
	}
	return set
}

func readUserInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}

// readLines reads a file into a slice of strings, separated by newlines.
func readLines(filename string) []string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var ls []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	return ls

}
