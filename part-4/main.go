package main

import (
	"fmt"
	"strings"
	"unicode"
)

//type counter stores the number of digits in each word

type counter map[string]int

// countDigitsInWords counts the number of digits in word of a phrase.
func countDigitsInWords(phrase string) counter {

	words := strings.Fields(phrase)
	stats := make(counter)

	type result struct {
		word  string
		count int
	}
	counted := make(chan result)

	//Launch a goroutine to process words and send result to the channel
	go func() {
		for _, word := range words {
			digitCount := countDigits(word)
			counted <- result{word, digitCount}
		}
		close(counted)
	}()

	//collect result from the channel and populate the stats map

	for res := range counted {
		stats[res.word] = res.count
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

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s:%d\n", word, count)

	}

}

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)

}
