package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const randomNumbersCount int = 10
	const randomRange int = 100
	randomNumberChannel := make(chan int, randomNumbersCount)
	powChannel := make(chan int, randomNumbersCount)
	result := make([]int, 0)

	for i := 1; i <= randomNumbersCount; i++ {
		go genRandomNumbers(randomRange, randomNumberChannel)
	}

	for range randomNumbersCount * 2 {
		select {
		case randomNumber := <-randomNumberChannel:
			go powNumbers(randomNumber, powChannel)
		case powNumber := <-powChannel:
			result = append(result, powNumber)
		}
	}

	fmt.Printf("%v\n", result)
}

func genRandomNumbers(randomRange int, channel chan int) {
	digit := rand.Intn(randomRange)
	channel <- digit
}

func powNumbers(number int, channel chan int) {
	channel <- number * number
}
