package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	//perulangan yang akan menambah jumlah Goroutine
	for i := 0; i < 1; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu", totalCpu)

	//kalo diatas 0 akan mengubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	//perulangan yang akan menambah jumlah Goroutine
	for i := 0; i < 1; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu", totalCpu)

	//menambah jumlah thread
	runtime.GOMAXPROCS(20)
	//kalo diatas 0 akan mengubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}
