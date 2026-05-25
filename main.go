package main

import "fmt"

func main() {
	numbers := []int{15, 22, 37, 40, 55, 62}

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i])
		}
	}

}
