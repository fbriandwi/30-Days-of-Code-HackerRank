package main

import "fmt"

func main() {
	var tests int
	fmt.Scan(&tests)

	for i := 0; i < tests; i++ {
		var s string
		fmt.Scan(&s)

		evenChars := ""
		oddChars := ""

		for j, char := range s {
			if j%2 == 0 {
				evenChars += string(char)
			} else {
				oddChars += string(char)
			}
		}

		fmt.Println(evenChars, oddChars)
	}
}
