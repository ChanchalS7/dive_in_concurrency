package main

import (
	"fmt"
	"strings"
	"unicode"
)

//Generator
//Counter stores the number of digits in each word.
// The key is the word and the value is the number of digits.

type counter map[string]int

//countDigitsInWords counts the number of digits in words,
//fetching the next word with next().

func countDigitsInWords(next func() string) counter {

	stats := counter{}

	for {
		word := next()
		if word == "" {
			break
		}
		count := countDigits(word)
		stats[word] = count

	}
	return stats
}
func countDigits(str string) int {
	count := 0

	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats print the number of digits in words.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator returns a function that yeilds one word at time from the phrase.
func wordGenerator(phrase string) func() string {

	words := strings.Fields(phrase)
	index := 0

	return func() string {
		if index >= len(words) {
			return ""
		}
		word := words[index]
		index++
		return word
	}

}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
