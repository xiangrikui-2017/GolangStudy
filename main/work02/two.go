package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var num int
	addNum(&num)
	fmt.Println("num + 10 =", num)

	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	addSlice(&myslice)
	fmt.Println("myslice * 2 =", myslice)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go printOddOrEven(1, &wg) // 奇数
	go printOddOrEven(0, &wg) // 偶数
	wg.Wait()

	tasks := []func(){
		func() {
			fmt.Println("My Task 1")
			time.Sleep(100 * time.Millisecond)
		},
		func() {
			fmt.Println("My Task 2")
			time.Sleep(200 * time.Millisecond)
		},
		func() {
			fmt.Println("My Task 3")
			time.Sleep(300 * time.Millisecond)
		},
		func() {
			fmt.Println("My Task 4")
			time.Sleep(400 * time.Millisecond)
		},
		func() {
			fmt.Println("My Task 5")
			time.Sleep(500 * time.Millisecond)
		},
	}
	exeTasks(&tasks)
	fmt.Println("All Tasks Completed")

}

/*
 * 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
 */
func addNum(num *int) {
	*num += 10
}

/*
 * 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
 */
func addSlice(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		(*slice)[i] *= 2
	}
}

func printOddOrEven(remainder int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == remainder {
			fmt.Println("偶数：", i)
		} else if i%2 == remainder {
			fmt.Println("奇数：", i)
		}
	}
}

func exeTasks(tasks *[]func()) {
	ch := make(chan struct{})

	for _, task := range *tasks {
		go func(t func()) {
			start := time.Now()
			t()
			fmt.Printf("Task completed in %v\n", time.Since(start))
			ch <- struct{}{}
		}(task)
	}
	for range *tasks {
		<-ch
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Shape
}

func (r Rectangle) Area() float64 {
	fmt.Println("调用Rectangle的Area方法")
	return 20
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("调用Rectangle的Perimeter方法")
	return 20
}

type Circle struct {
	Shape
}

func (c Circle) Area() float64 {
	fmt.Println("调用Circle的Area方法")
	return 20
}

func (c Circle) Perimeter() float64 {
	fmt.Println("调用Circle的Perimeter方法")
	return 20
}
