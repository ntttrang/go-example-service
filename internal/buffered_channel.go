package internal

import (
	"fmt"
	"time"
)

func TestBufferedChannel() {
	ch := make(chan int, 1)

	// read from empty buffered channel => blok current goroutines
	<-ch

	// write to full buffered channel -> blok current goroutines
	ch <- 1

}

func TestBufferedChannel2() {
	ch := make(chan int, 10) // Work as a queue ( FIFO)

	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			fmt.Println("Wrote: ", i)
		}
	}()

	for i := 1; i <= 20; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Chan value: ", <-ch)
	}

}

func TestBufferedChannel3() {
	ch := make(chan int, 10) // Work as a queue ( FIFO)

	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			fmt.Println("Wrote: ", i)
		}
		close(ch)
	}()

	// Cách 1: Đọc giá trị từ channel
	for v := range ch {
		fmt.Println("Chan value: ", v)
	}
	// Cách 2: Đọc giá trị từ channel ( Hiếm dùng)
	// for {
	// 	v, isActive := <-ch
	// 	if !isActive {
	// 		break
	// 	}
	// 	fmt.Println("Chan value: ", v)
	// }
}
