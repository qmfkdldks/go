package main

import (
    "fmt"
)

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	for _, value := range(ar) {
		sum += value
	}
    return sum
}


func main() {
	array := []int32 { 2, 3, 6, 2 }
	result := simpleArraySum(array)
	fmt.Printf("result: %d \n", result)
}