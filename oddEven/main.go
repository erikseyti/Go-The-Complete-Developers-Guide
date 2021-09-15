package main

func main() {
	num := createGroupOfNumbers(10)

	for _, n := range num {
		isOddOrEven(n)
	}
}
