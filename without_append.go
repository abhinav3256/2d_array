package main

import "fmt"

func main() {

	input := [][]int{
		{4, 8, 6},
		{3, 5, 4},
	}
	//fmt.Println(input)

	var sum int
	//var sumArray = []int{}
	var largest int

	for i := 0; i < len(input); i++ {
		sum = 0

		for j := 0; j < len(input[i]); j++ {
			//fmt.Println(input[i][j])
			sum = input[i][j] + sum

		}
		fmt.Println(sum)
		if largest < sum {

			largest = sum
		}

		//sumArray = append(sumArray, sum)
	}

	fmt.Println("largest wealth:", largest)
}
