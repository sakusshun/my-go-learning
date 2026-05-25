package main

import "fmt"

func main() {
	fruits := []string{"яблоко", "банан", "апельсин"}
	fmt.Println(fruits)

	fruits = append(fruits, "киви", "манго")
	fmt.Println(fruits)
	fmt.Println(fruits[1])
}
