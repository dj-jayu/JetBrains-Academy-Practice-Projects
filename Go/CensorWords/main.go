package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func createMap(file *os.File) map[string]bool {
	tabooWords := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tabooWords[strings.ToLower(scanner.Text())] = true
	}
	if scanner.Err() != nil {
		log.Fatal("scanner Error")
	}
	return tabooWords
}

func censorPhrases(tabooWords map[string]bool) {
	var inputWords []string
	var builder strings.Builder
	// var endChar string
	var scanner = bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		inputWords = strings.Fields(scanner.Text())
		if len(inputWords) == 1 && inputWords[0] == "exit" {
			break
		}
		for _, word := range inputWords {
			if _, ok := tabooWords[strings.ToLower(word)]; ok {
				builder.WriteString(strings.Repeat("*", len(word)) + " ")
			} else {
				builder.WriteString(word + " ")
			}
		}
		fmt.Println(strings.TrimSpace(builder.String()))
		builder.Reset()
	}
}
func main() {
	var fileName string
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var tabooWords = createMap(file)
	censorPhrases(tabooWords)
	fmt.Println("Bye!")
}
