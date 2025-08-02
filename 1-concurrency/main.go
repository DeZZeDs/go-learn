package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const randomNumbersCount int = 10
	const randomRange int = 100
	randomNumberChannel := make(chan int, randomNumbersCount)
	powChannel := make(chan int, randomNumbersCount)
	wg := sync.WaitGroup{}
	result := make([]int, 0)

	for i := 1; i <= randomNumbersCount; i++ {
		wg.Add(1)
		go genRandomNumbers(randomRange, randomNumberChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(randomNumberChannel)
		close(powChannel)
	}()

	for randomNumber := range randomNumberChannel {
		wg.Add(1)
		go powNumbers(randomNumber, powChannel, &wg)
	}

	for powedNumber := range powChannel {
		result = append(result, powedNumber)
	}

	fmt.Printf("%v\n", result)
}

func genRandomNumbers(randomRange int, channel chan int, wg *sync.WaitGroup) {
	digit := rand.Intn(randomRange)
	channel <- digit
	wg.Done()
}

func powNumbers(number int, channel chan int, wg *sync.WaitGroup) {
	channel <- number * number
	wg.Done()
}
