/*
1.Suppose we have function that speaks word by word with some pauses :
2. Now lets' create two talkers each saying their own phrase
Not bad but the function speaks one after another. Tome make them speak at the same time, let's add "go" before calling the function.
Now they complete for out attention ! when we write go f() the function f() runs independently
Function that runs with go are called goroutines. The Go runtime judges these goroutines and distribute them among operating system threads running on CPU cores. Compare to os threads. goroutines are lightweight threads, so you can create hundreds of them.
You may wondering why we need time.Sleep() in the main function. Let's clarify that.
Go routines are completely independent, when we call go say() the function runs on its own. main doesn't wait for it. So if we write like this:
like without time.Sleep()
func main() {

	go say(1, "go is awesome")
	go say(2, "citty are cutie")

}
- The program won't print anything. main finishes before our go routines speak and since the main is done the whole program terminates.
The main function is also a goroutine but starts only when the program start.
So we have three go routine all of them which are independent. The catch is  that when main() ends everything else ends too.
Using time.Sleep() to wait for goroutines is a bad idea because we can't predict how long they will take.
So a better approach is wait group:
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//	func say(id int, phrase string) {
//		for _, word := range strings.Fields(phrase) {
//			fmt.Printf("Worker %d and Harsh says : %s...\n", id, word)
//			dur := time.Duration(rand.Intn(1000)) * time.Millisecond
//			time.Sleep(dur)
//		}
//	}
//
//	func main() {
//		/*
//			say(1, "go is awesome")
//			say(2, "citty are cutie")
//		*/
//		go say(1, "go is awesome")
//		go say(2, "citty are cutie")
//		time.Sleep(500 * time.Millisecond)
//	}
func say(wg *sync.WaitGroup, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worder #%d says :%s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say(&wg, 1, "go is awesome")

	wg.Add(1)
	go say(&wg, 2, "citty are cute")
	wg.Wait()
}
