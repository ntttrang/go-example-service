package internal

import (
	"fmt"
	"time"
)

func TestChannel() {
	// Block current goroutines ( lightweight thread)
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Second / 2)
			fmt.Println("Goroutine 2:", <-ch)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Goroutine 1:", <-ch)
		time.Sleep(time.Second) // Sau 1s, nó sẽ lấy từ channel ra. Ko phụ thuộc khi nào put vào channelchannel
		ch <- i * 2
	}
}
