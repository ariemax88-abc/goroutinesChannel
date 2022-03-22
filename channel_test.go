package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreteChannel(t *testing.T) {
	channel := make(chan string)
	//defer close(channel)

	// channel <- "Ari"
	// fmt.Println(<-channel)

	//jika data disimpan dalam variabel
	// data := <-channel
	// fmt.Println(data)

	go func() {
		time.Sleep(2 * time.Second)
		//channel <- "Ari Bayu Prasetyo"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
