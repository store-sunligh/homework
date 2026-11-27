package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Pay interface {
	pay()
}

type Alipay struct {
}

func (a *Alipay) pay() {
	fmt.Println("支付宝支付")
}

type WeiChat struct {
}

func (w *WeiChat) pay() {
	fmt.Println("微信支付")
}

type Bank struct {
}

func (b *Bank) pay() {
	fmt.Println("银行卡支付")
}
func paying(p Pay) {
	p.pay()
}
func pay() {

	fmt.Println("======支付多态作业======")

	p := &Alipay{}
	paying(p)
	a := &WeiChat{}
	paying(a)
	b := &Bank{}
	paying(b)
}

type SafeLog struct {
	lock sync.RWMutex
	logs []string
}

func NewSafeLog() *SafeLog {
	return &SafeLog{
		logs: make([]string, 0),
	}
}

func (s *SafeLog) AppendLog(log string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.logs = append(s.logs, log)
}

func (s *SafeLog) GetAll() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.logs
}

func log() {
	fmt.Println("=========安全日志作业===========")
	log := NewSafeLog()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log.AppendLog(fmt.Sprintf("%d线程加入日志完成", id))
		}(i)
	}
	wg.Wait()
	fmt.Println("所有日志写入完成")
	for v, value := range log.GetAll() {
		fmt.Println(v, value)
	}
}

func work(ctx context.Context, id int, job <-chan int, result chan<- int, group *sync.WaitGroup) {
	defer group.Done()

	for {
		select {
		case val, ok := <-job:
			if !ok {
				fmt.Println("已关闭")
				return
			}
			fmt.Printf("work %d 处理数据 %d\n", id, val)

			select {
			case result <- val * 2:
				fmt.Printf("work %d 处理数据 %d 结果是%d\n", id, val, val*2)
			case <-ctx.Done():
				fmt.Println("结束")
				return
			}

		case <-ctx.Done():
			fmt.Println("结束")
			return
		}
	}
}

func workPoll() {
	fmt.Println("=======任务调度作业==========")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	job := make(chan int)
	result := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go work(ctx, i, job, result, &wg)
	}
	go func() {
		for i := 0; i < 10; i++ {
			job <- i
		}
		close(job)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		fmt.Println("收到结果", v)
	}

	fmt.Println("完成")

}

func main() {
	//pay()
	//log()
	workPoll()
}
