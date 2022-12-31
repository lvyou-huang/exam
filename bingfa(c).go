package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 1)
	go work("1", ch)
	<-ch
	go work("2", ch)
	<-ch
	go work("3", ch)
	<-ch
	go work("4", ch)
	<-ch
	go work("5", ch)
	<-ch
	go work("6", ch)
	<-ch
	go work("7", ch)
	<-ch
	go work("8", ch)
	<-ch
	go work("9", ch)
	<-ch
	go work("10", ch)
	<-ch

}

func work(s string, ch chan int) {
	fmt.Println(s)
	ch <- 1
}
