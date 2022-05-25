package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup //Counter for waiting.
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done() //Decrements counter by 1
	}()

	wg.Wait() // Wait until count is done.

	//go count("sheep")
	//go count("fish")
	//time.Sleep(time.Millisecond * 200)

	//fmt.Scanln()
}

func count(thing string) {  //String channel to communicate to the main routine
	for i:= 1; i <= 5; i++ {		
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}


}