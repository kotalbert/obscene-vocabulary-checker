package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fn := readUserInput() // name of file to read
	ws := readLines(fn)   // list of taboo words

	for {
		s := readUserInput() // sentence to censor
		if s == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		} else {
			var censored string
			censored = censorSentence(ws, s)
			fmt.Println(censored)
		}
	}

}

func censorSentence(ws []string, s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		if isTabooWord(ws, w) {
			words[i] = censor(w)
		}
	}
	return strings.Join(words, " ")
}

// censor replaces each character in a string with an asterisk.
func censor(tw string) string {
	return strings.Repeat("*", len(tw))
}

// isTabooWord checks if a word is in a list of taboo words.
func isTabooWord(ws []string, tw string) bool {
	utw := strings.ToLower(tw)
	for _, w := range ws {
		uw := strings.ToLower(w)
		if uw == utw {
			return true
		}
	}
	return false
}

// readUserInput reads a line of input from the user.
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
