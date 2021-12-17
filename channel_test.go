package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	//membuat variabel untuk channel
	channel := make(chan string)
	//defer untuk memastikan, mau eror/sukse akan dieksekusi
	defer close(channel)

	//goroutines dengan anonimus function
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Galih Setiadi"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	/*
		//mengirim data ke chanel
		channel <- "Galih"
		//mengambil data dari channel
		//dari variabel
		data := <- string
		data <- channel
		//langsung ke parameter
		fmt.Println(<- channel)*/
}

//channel sebagai parameter. Tidak butuh pointer
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Galih Setiadi"
}

//Test channel as parameter
func TestChannelAsParameter(t *testing.T) {
	//membuat variabel untuk channel
	channel := make(chan string)
	//defer untuk memastikan, mau eror/sukse akan dieksekusi
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
