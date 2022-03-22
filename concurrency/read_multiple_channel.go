package main

import (
	"fmt"
	"time"
)

func writeInt(c chan int) {
	for i := 1; i <= 3; i++ {
		c <- i
		fmt.Printf("write int = %v\n", i)
	}
}

func writeString(c chan string) {
	for i := 97; i <= 99; i++ {
		c <- string(i)
		fmt.Printf("write rune %v\n", string(i))
	}
}

func main() {

	intChan := make(chan int, 3)
	stringChan := make(chan string, 3)

	go writeInt(intChan)
	go writeString(stringChan)

	i := 0
	for {
		select {
		case v := <-intChan:
			fmt.Printf("Read int data = %v\n", v)
			time.Sleep(time.Second)

		case n := <-stringChan:
			fmt.Printf("Read string data = %v\n", n)
			time.Sleep(time.Second)
		default:
			fmt.Println("No input")
		}

		i++
		if i == 10 {
			break
		}
	}

	fmt.Println("End process")
}
