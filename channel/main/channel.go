package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int, 20)
	go writeChannel(channel)
	go readChannel(channel)
	wg.Add(2)
	wg.Wait()
}

func writeChannel(channel chan int) {
	defer close(channel)
	defer wg.Done()
	for i := 0; i < 20; i++ {
		channel <- i
		fmt.Println("写入数据：", i)
	}
}

func readChannel(channel chan int) {
	defer wg.Done()
	for value := range channel {
		fmt.Println("读出数据：", value)
	}
}
