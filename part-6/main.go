package main

import (
	"fmt"
	"strings"
	"unicode"
)

//counter stores the number of digits in each word
// the key is the word, and the value is the number of digits.

type counter map[string]int

// pair stores a word and number of digits
type pair struct {
	word  string
	count int
}

//countDigitsInWords counts the number of digits in words,
//fetching the next word with next().

func countDigitsInWords(next func() string) counter {
	counted := make(chan pair)

	//start a goroutine to read words and send digit count
	go func() {
		for {
			word := next()
			if word == "" {
				break
			}
			counted <- pair{word, countDigits(word)}
			// close(counted) //signal that no more data will be sent
		}
		close(counted) //only close once all send are done
	}()
	stats := counter{}
	for p := range counted {
		stats[p.word] = p.count
	}
	return stats
}

// countDigits returns the number of digits in a string
func countDigits(str string) int {
	count := 0
	for _, char := range str {

		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// print stats prints the number of digits in words.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s:%d\n", word, count)
	}
}

// wordGenerator returns a generator that yields words from a phrase.
func wordGenerator(phrase string) func() string {

	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}

}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
