package goroutines

import (
	"fmt"
	"strconv"
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

//cuma boleh mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Galih Setiadi"

}

//cuma boleh menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

//test dari InOutChannel
func TestInOutChannel(t *testing.T) {
	//membuat variabel untuk channel
	channel := make(chan string)
	//defer untuk memastikan, mau eror/sukse akan dieksekusi
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

//Membuat buffered channel. Buffered channel ibaratkan memory untuk channel. Kalau tidak diberikan buffered,
//maka channel hanya akan bisa menerima dan mengirim satu data.
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Galih"
		channel <- "Minavf"
		channel <- "Setiadi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("")
	fmt.Println("Selesai")

}

//Range Channel.Range Channel untuk membuat perulangan, ketika untuk menanggulangi data yang dikirim terus
//menerus dan kita tidak tahu jelas kapan channel tersebut akan berhenti menerima data

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			//Itoa merupakan konversi dari integer ke string
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		//setiap channel harus di close, dalam kasus ini kalau tidak di close maka perulangan dibawahnya tidak akan pernah berhenti
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

//Select channel untuk mendapatkan data dari semua channel
func TestChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	//pemberian counter untuk menghindari deadlock
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}
