package main

import (
	"fmt"
	"math"
)

func calTotal(a int, res int) int {

	return int(math.Pow(float64(a), 2)) + res
}

func getnumbers(cnt int, res int) int {
	var tt int
	for cnt < 0 {
		if _, err := fmt.Scan(&tt); err == nil {
			calTotal(tt, res)
		}
	}
	//test
	return res

}
func fory(cnt int, tot int) int {

	for cnt < 0 {
		getnumbers(cnt, tot)
		cnt--
	}
	return tot

}
func main() {
	//reader := bufio.NewReader(os.Stdin)

	var n int

	if _, err := fmt.Scan(&n); err == nil {
		total := make([]int, n)
		for i := 0; i <= n; i++ {
			var x int
			var tot int
			if _, err := fmt.Scan(&x); err == nil {
				total[i] = fory(x, tot)
			}
		}
		fmt.Println("Result!!", total)
	} else {
		fmt.Println(err)
	}

}
