package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%d ", A[i])
	}
}
