package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id                int
	schedulerToWorker chan int
	workerToScheduler chan int
	doneWorker chan bool
	times int
}

func createWorker(id int, schedulerToWorker chan int,
	workerToScheduler chan int, doneWorker chan bool) {

	// 创建 Worker，绑定管道
	w := Worker{
		id:                id,
		schedulerToWorker: schedulerToWorker,
		workerToScheduler: workerToScheduler,
		doneWorker: doneWorker,
		times: 0,
	}

	// 开始工作
	go w.workLoop()
}

func (w *Worker) workLoop() {

	// 一直工作
	for {
		w.times++
		fmt.Println(w.id, w.times, "A")
		// 存储从schedulerToWorker拿到的数据
		inData := <-w.schedulerToWorker
		fmt.Println(w.id, w.times, "AAA")
		// 进行数据处理，并获取需要传进workerToScheduler的数据
		ret := w.process(inData)

		fmt.Println(w.id, w.times, "B")

		if ret == nil || len(ret) == 0 {
			break
		}

		fmt.Println(w.id, w.times, "C")

		go func() {
			// 将处理好的数据传回给 workerToScheduler
			for _, item := range ret {
				w.workerToScheduler <- item
			}
		}()
	}

	fmt.Println(w.id, w.times, "D")
	// 往 doneWorker 管道中送入一个数据，代表该Worker工作已经完成
	w.doneWorker <- true

}

// 数据处理过程
func (w *Worker) process(data int) []int {

	rand.Seed(time.Now().UnixNano())

	var ret []int

	retSize := rand.Intn(3)

	for i := 0; i < retSize; i++ {
		ret = append(ret, rand.Intn(100))
	}

	fmt.Printf("Worker %v : Accept %v : Return %v\n", w.id, data, ret)

	return ret
}

func scheduler(seed int, num int, finish chan bool) {

	// 创建 Scheduler --> Worker 管道
	schedulerToWorker := make(chan int,num)
	// 创建 Worker --> Scheduler 管道
	workerToScheduler := make(chan int,num)
	// 创建 DoneWorker 管道，用于记录完成工作的Worker数量
	doneWorker := make(chan bool)
	// 用于统计完成工作的Worker数量
	doneWorkerCount := 0

	// 创建Worker，每个Worker与Scheduler通过管道连接
	// Worker一旦创建，就开始工作
	for i := 0; i < num; i++ {
		go createWorker(i, schedulerToWorker, workerToScheduler, doneWorker)
	}

	schedulerToWorker <- seed

	// 开一个 goroutine 收集workerToScheduler的数据，加入schedulerToWorker中
	for {
		select {
		case n := <-workerToScheduler:
			schedulerToWorker <- n
		case <-doneWorker:
			doneWorkerCount++
			fmt.Println("doneWorkerCount：",doneWorkerCount)
			if doneWorkerCount == num {
				break
			}
		}
	}

	//for {
	//	schedulerToWorker <- <-workerToScheduler
	//}

    //doneWorkerCount = doneWorkerCount
    //finish <- true
}

func main() {


	// 创建 finish 管道
	//finish := make(chan bool)

	go scheduler(666, 2 ,nil)

	time.Sleep(100 * time.Second)

	//<- finish
}
