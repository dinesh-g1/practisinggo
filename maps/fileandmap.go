package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("ramayan.txt")
	if err != nil {
		log.Printf("can't open the file. err : %v", err)
	}
	words, _ := wordFrequency(f)
	for k, v := range words {
		fmt.Println(k, "-", v)
	}
}

func wordFrequency(content io.Reader) (map[string]int, error) {
	sc := bufio.NewScanner(content)
	wordFreq := make(map[string]int)
	for sc.Scan() {
		line := sc.Text()
		line = strings.ToLower(line)
		words := strings.Split(line, " ")

		for _, w := range words {
			wordFreq[w] = 1
		}
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}
