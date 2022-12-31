package main

import "fmt"

func main() {
	var a, b, n, password int = 0, 0, 0, 0
	fmt.Scanf("%d %d\n", &a, &b)
	buildingStaircase := make([][]int, a)
	buildingId := make([][]int, a)
	for i := range buildingStaircase {
		buildingStaircase[i] = make([]int, b)
	}
	for i := range buildingId {
		buildingId[i] = make([]int, b)
	}
	fmt.Println(buildingStaircase)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Scanf("%d %d\n", &buildingStaircase[i][j], &buildingId[i][j])

		}
	}
	fmt.Println(buildingStaircase, buildingId)
	fmt.Scanf("%d\n", &n)
	i := 0
	password = buildingId[i][n]
	for {
		if buildingStaircase[i][n] == 1 {

			i++
			password += buildingId[i][n]

		} else {
			
			n -= buildingId[i][n] - 1
			if n < 0 {
				n = b + n
			}
		}
		if i == a-1 {
			break
		}
	}
	fmt.Println(password)

}
