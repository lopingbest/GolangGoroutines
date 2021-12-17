package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutines(t *testing.T) {
	//secara synchronous
	//RunHelloWorld()
	// penambahan "go" di depan untuk membuat kode asynchronous
	go RunHelloWorld()
	fmt.Println("Ups")

	//penundaan untuk memastikan goroutines dieksekusi
	time.Sleep(1 * time.Second)
}

// Membuat banyak goroutines untuk mengetes memori
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
