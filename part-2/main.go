package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

func countDigitsInWords(phrase string) counter {
	var wg sync.WaitGroup

	syncStats := new(sync.Map)

	words := strings.Fields(phrase)

	//count the number of digits in words
	//using a separate goroutine  for each word
	// to store the result of the count
	// use syncStates.Store(word,count)
	// As a result, syncStats should contain words
	// and the number of digits in each
	return asStats(syncStats)
}

// return the number of digits in a string
func countDigits(str string) int {
	count := 0

	for _, char := range str {

		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// asStats converts statistics from sync.Map to a regular map.
func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s:%d\n", word, count)
	}
}
func main() {
	// To create a channel, use make(chan typr)
	// channel can only accept values of the specified type:
	// message := make(chan string)

	//to send value to a channel
	// use channel <- syntax
	// Lets send "ping"
	// go func() { message <- "ping" }()

	//to receive a value from  a channel,
	// use the <- channel syntax.
	//Let's receive "ping" and print it.
	// msg := <-message
	// fmt.Println(msg)

	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)

}
