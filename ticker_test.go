package goroutines

import (
	"fmt"
	"testing"
	"time"
)

//ticker digunakan untuk waktu yang berulang, kalo time.timer digunakan untuk satu waktu
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	//untuk menghentikan ticker
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
	//kode ini akan menghasilkan deadlock, jika ingin kode berhentitanpa deadlock maka gamntifor range dengan select
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
