package main

import (
	"fmt"
	"sync"
)

func printMe(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("I'm a go routine %d\n", i)
}

func main() {

	rounds := 10

	var wg sync.WaitGroup

	wg.Add(rounds)

	for i := 0; i < rounds; i++ {
		go func(rI int) {
			printMe(rI, &wg)
		}(i)
	}

	wg.Wait()
}
