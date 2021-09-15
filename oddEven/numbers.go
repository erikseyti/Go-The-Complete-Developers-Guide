package main

import (
	"fmt"
)

type numbers []int

func createGroupOfNumbers(maxNumber int) numbers {
	num := numbers{}

	for i := 0; i < maxNumber+1; i++ {
		num = append(num, i)
	}

	return num
}

func isOddOrEven(n int) {
	if n%2 == 0 {
		// fmt.Sprintln("the number" + n + "is Even")
		fmt.Printf("The number %v is Even \n", n)
	} else {
		fmt.Printf("The number %v is Odd \n", n)
	}

}
