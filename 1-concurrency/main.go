package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const randomNumbersCount int = 10
	const randomRange int = 101
	randomNumberChannel := make(chan int, randomNumbersCount)
	powChannel := make(chan int, randomNumbersCount)
	result := make([]int, 0)

	go genRandomNumbers(randomRange, randomNumbersCount, randomNumberChannel)
	go powNumbers(randomNumberChannel, powChannel)

	for powedNumber := range powChannel {
		result = append(result, powedNumber)
	}

	for i := range result {
		fmt.Printf("%d ", result[i])
	}
}

func genRandomNumbers(randomRange int, randomNumbersCount int, channel chan int) {
	for i := 1; i <= randomNumbersCount; i++ {
		digit := rand.Intn(randomRange)
		channel <- digit
	}
	close(channel)
}

func powNumbers(channel chan int, powChannel chan int) {
	for number := range channel {
		powChannel <- number * number
	}
	close(powChannel)
}
