package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
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

	r := Rectangle{Width: 5, Height: 3}
	c := Circle{Radius: 2}
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", r.Area(), r.Perimeter())
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", c.Area(), c.Perimeter())
	fmt.Println()

	p := Person{Name: "悦", Age: 20}
	e := Employee{EmployeeId: "28674386723", Person: p}
	e.printEmployee()

	ch := make(chan int, 10)
	go func() {
		// defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for val := range ch {
			fmt.Println("channel val :", val)
		}
	}()
	time.Sleep(1 * time.Second)

	size := 100
	wg.Add(2)

	go func(size int) {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= size; i++ {
			ch <- i
		}
	}(size)

	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println("channel val :", val)
		}
	}()
	wg.Wait()

	lock := sync.Mutex{}
	var total int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			for i := 1; i <= 1000; i++ {
				total++
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Total: %d\n", total)
	fmt.Println()

	counter := int32(0)
	atomicAdd(&counter)
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
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId string
}

func (e *Employee) printEmployee() {
	fmt.Printf("EmployeeName : %v, EmployeeAge : %v, EmployeeId : %v\n", e.Person.Name, e.Person.Age, e.EmployeeId)
}

func atomicAdd(counter *int32) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Counter: %d\n", *counter)
	fmt.Println()
}
