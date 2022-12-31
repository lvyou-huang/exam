package main

import "fmt"

// a/b抛出错误 后，程序往上知行defer语句，故不会执行下一步
func main() {
	a := 2
	b := 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("2/0=", a/b)
	fmt.Println("3/2 =", 3/2)
}
