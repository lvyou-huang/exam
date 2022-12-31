package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var a int = 1
	var ch = make(chan int, 1)
	go work(&a, ch)
	go work(&a, ch)
	go work(&a, ch)
	go work(&a, ch)
	go work(&a, ch)
	<-ch

}

func work(a *int, ch chan int) {
	var flag bool = true
	for ; *a < 1000; *a++ {
		for i := 2; i <= *a/2; i++ {
			if *a%i == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Println(*a)
		} else {
			flag = true
		}
	}
	ch <- 1

}
