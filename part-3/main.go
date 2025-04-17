package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		fmt.Println("B: Sending message...")
		message <- "ping"              //1
		fmt.Println("B: Message sent") //2
	}()

	fmt.Println("A: Doing some work...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-message //3

	fmt.Println("A: Message received!")
	time.Sleep(100 * time.Millisecond)
}

// // Result channel.
// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// // counter stores the number of digits in each word.
// // The key is the word, and the value is the number of digits.
// type counter map[string]int

// // solution start

// // countDigitsInWords counts the number of digits in the words of a phrase.
// func countDigitsInWords(phrase string) counter {
// 	words := strings.Fields(phrase)
// 	counted := make(chan int)

// 	go func() {
// 		// Loop through the words,
// 		// count the number of digits in each,
// 		// and write it to the counted channel.
// 	}()

// 	// Read values from the counted channel
// 	// and fill stats.

// 	// As a result, stats should contain words
// 	// and the number of digits in each.

// 	return stats
// }

// // solution end

// // countDigits returns the number of digits in a string.
// func countDigits(str string) int {
// 	count := 0
// 	for _, char := range str {
// 		if unicode.IsDigit(char) {
// 			count++
// 		}
// 	}
// 	return count
// }

// // printStats prints the number of digits in words.
// func printStats(stats counter) {
// 	for word, count := range stats {
// 		fmt.Printf("%s: %d\n", word, count)
// 	}
// }

// func main() {
// 	phrase := "0ne 1wo thr33 4068"
// 	stats := countDigitsInWords(phrase)
// 	printStats(stats)
// }
