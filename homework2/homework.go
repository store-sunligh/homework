package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func point(a *int) {
	*a *= 2
}

func slicessss(a *[]int) {
	for i := range *a {
		(*a)[i] *= 2
	}

}

func printoddeven() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				continue
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("work odd", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i <= 10; i++ {
			if i%2 == 1 {
				continue
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("work even", i)
		}
	}()
	wg.Wait()
}

type Job struct {
	id     int
	during int //耗时
}

func workings(jobs []Job, result chan<- Job, wg *sync.WaitGroup) {

	for _, job := range jobs {
		(*wg).Add(1)
		go func() {
			defer (*wg).Done()
			job.during = rand.Int()
			result <- job
		}()
	}

}

func workPollDemo() {
	fmt.Println("======任务调度器====")
	jobs := make([]Job, 0)
	for i := 0; i < 10; i++ {
		jobs = append(jobs, Job{
			id: i,
		})
	}
	var wg sync.WaitGroup
	result := make(chan Job)
	workings(jobs, result, &wg)
	go func() {
		wg.Wait()
		close(result)
	}()
	for v := range result {
		fmt.Println(v)
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	height float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}
func (r *Rectangle) Perimeter() float64 {
	return (r.height + r.width) * 2
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func info(s Shape) {
	fmt.Println(s.Area(), s.Perimeter())
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Println("Employee:", *e)
}

func channal1() {
	chanint := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			chanint <- i
		}
		close(chanint)
	}()
	for v := range chanint {
		fmt.Println(v)
	}
}
func product(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}
func custom(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}
func channal2() {
	channal := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go product(channal, &wg)
	go custom(channal, &wg)
	wg.Wait()
	fmt.Println("finish")
}

func mutex1() {
	var sm sync.Mutex
	var wg sync.WaitGroup
	a := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Lock()
			defer sm.Unlock()
			for i := 0; i < 1000; i++ {
				a++
			}

		}()
	}
	wg.Wait()
	fmt.Println(a)

}

func mutex2() {
	var wg sync.WaitGroup
	var count int32 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func main() {
	/*a := 5
	point(&a)
	fmt.Println(a)*/

	/*b := []int{1, 2, 3}
	slicessss(&b)
	fmt.Println(b)*/

	//printoddeven()

	/*a := &Circle{3}
	info(a)
	b := &Rectangle{2.0, 4.0}
	info(b)*/

	/*e := Employee{
		Person:     Person{"jindou", 10},
		EmployeeID: "123",
	}
	e.PrintInfo()*/

	//channal1()
	//channal2()

	//mutex1()
	mutex2()
}
