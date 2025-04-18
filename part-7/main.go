package main

import (
	"fmt"
	"strings"
	"unicode"
)

//counter stores the number of digits in each word
//the key is the word and value is the number of digits.

type counter map[string]int

type pair struct {
	word  string
	count int
}

//solution start
//countDigitsInWords count the number of digits in words,
//fetching the next word with next()

func countDigitsInWords(next func() string) counter {
	pending := make(chan string)
	counted := make(chan pair)

	//sends words to be counted
	//Reader goroutine : fetches from next() and send them to pending channel.
	go func() {
		//fetch words from generator
		//and send them to the pending channel,

		for {
			word := next()

			if word == "" {
				break

			}
			pending <- word
		}
		close(pending) //Important to close so worker knows there are no more words.
	}()
	//counts digits in words
	// Worker goroutine : reads from pending, counts digits, sends result to counted.
	go func() {
		//read the words from the pending channel.
		//count the number of digits in each word.
		//and send the results to the counted channel.

		for word := range pending {
			count := countDigits(word)
			counted <- pair{word, count}
		}
		close(counted) //close after all words are processed.
	}()
	//read values from the counted channel
	//and fill stats.
	//as a result, stats should contain words
	// and the number of digits in each.

	//Main outer function : reads results from counted channel and updates stats.

	stats := counter{}
	for p := range counted {
		stats[p.word] = p.count
	}

	return stats
}

//solution end

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

// printStats print the number of digits in words.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s:%d\n", word, count)
	}
}

//wordGenerator returns a generator that yields words from a phrase.

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
