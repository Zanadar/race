package main

import (
	"fmt"
	"time"
)

type intCounter int64

func main() {
	counter := intCounter(0)

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter++
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(counter)

}
