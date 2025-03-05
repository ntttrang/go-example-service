package internal

import (
	"fmt"
	"time"
)

func CrawerUrl() {
	// Em hãy viết một chương trình mô phỏng việc crawl một lúc 10K URL,
	// Tuy nhiên là trong cùng một lúc chỉ tối đa có 5 URLs được crawl,
	// Hãy dùng channel để mô phỏng, lưu ý là URL có thể là một mảng số

	// đổ vào 1 queue => buffered channel
	const numberOfURL = 10000
	const maxConncurrent = 5

	queueChan := make(chan int, numberOfURL)

	for i := 1; i <= numberOfURL; i++ {
		queueChan <- i
	}

	close(queueChan)

	checkDoneChan := make(chan int)

	// 5 URLs tối đa được chạy trong 1 lúc
	// => only 5 goroutines running on the same time
	for i := 1; i <= maxConncurrent; i++ {
		// process/thread
		go func(name string) {
			for urlValue := range queueChan {
				time.Sleep(time.Second / 2)
				msg := fmt.Sprintf("Goroutine %s is crawling url: %d. \n", name, urlValue)
				fmt.Println(msg)
			}

			fmt.Printf("%s done.\n", name)
			checkDoneChan <- 1

		}(fmt.Sprintf("%d", i))
	}

	for i := 1; i <= maxConncurrent; i++ {
		<-checkDoneChan
	}

	fmt.Println("All urls finished.")

	// Cach 1: Su dung channel
	// Cach 2: Wait Group
	// Cach 3: checkCount -- => Racing condition. Neu muon su dung, can them mutex Lock/ UnLock
	// Cach 4: Integer Atomic

}
