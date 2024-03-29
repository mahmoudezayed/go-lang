package main

import (
	"fmt"
	"log"
	"time"
)

func sumArray(data []int) int {
	size := len(data)
	total := 0

	for i := 0; i < size; i++ {
		total += data[i]
	}

	return total
}

func main() {
	// Helper to calculate the excution time.
	start := time.Now()
	// time.Sleep(time.Second * 2)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("The sum of all values in array:", sumArray(data))

	// Helper to calculate the excution time.
	log.Printf("Excution took %s \n", time.Since(start))
}
