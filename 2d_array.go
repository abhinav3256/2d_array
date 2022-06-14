package main

import "fmt"

func main() {

	input := [][]int{
		{4, 8, 6},
		{3, 5, 4},
	}
	//fmt.Println(input)

	var sum int
	var sumArray = []int{}
	for i := 0; i < len(input); i++ {
		sum = 0
		for j := 0; j < len(input[i]); j++ {
			//fmt.Println(input[i][j])
			sum = input[i][j] + sum

		}
		sumArray = append(sumArray, sum)
	}
	fmt.Println(sumArray)

	var largest = sumArray[0]
	for k := 0; k < len(sumArray); k++ {
		if largest < sumArray[k] {
			largest = sumArray[k]
		}
	}
	fmt.Println("largest wealth:", largest)
}
