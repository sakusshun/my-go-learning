package main

import "fmt"

func main() {
	friends := []string{"Аня", "Кирилл", "Настя", "Анатолий"}

	for i := 0; i <= 3; i++ {
		fmt.Println("Привет! ", friends[i])
	}
}
