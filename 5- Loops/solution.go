package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n) // Membaca sebuah bilangan bulat dari pengguna

	// Perulangan dari 1 hingga 10 dan cetak tabel perkalian
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
