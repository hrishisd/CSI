package main

import "fmt"

func main() {
	source := make(chan int)
	ints := make(chan int)
	squares := make(chan int)

	// Squarer
	go func() {
		for {
			i, ok := <-ints
			if !ok {
				break
			}
			squares <- (i * i)
		}
		close(squares)
	}()

	// Printer
	go func() {
		for {
			i, ok := <-squares
			if !ok {
				break
			}
			fmt.Println(i)
		}
	}()

	go pipe(source, ints)

	// Counter
	for i := 0; i < 10; i++ {
		source <- i
	}
	close(source)
}

func pipe(in <-chan int, out chan<- int) {
	for {
		out <- (<-in)
	}
}
