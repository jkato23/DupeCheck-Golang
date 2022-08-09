package main

import "fmt"

func main() {
	var numArray [5]int
	var numContent int
	var dupeCount int
	fmt.Println("Hello")
	for i := 0; i < 5; i++ {
		fmt.Print("Enter a number to be added to the array to be checked. You will be prompted five times: ")
		_, err := fmt.Scanln(&numContent)
		if err != nil {
			fmt.Println("You entered in something that was not a number, please try again.")
			var discard string
			fmt.Scanln(&discard)
			i--
		} else {
			numArray[i] = numContent
		}
	}
	for i := 0; i < len(numArray); i++ {
		for j := i + 1; j < len(numArray); j++ {
			if numArray[i] == numArray[j] {
				dupeCount++
				break
			}
		}
	}
	fmt.Println("The number of duplicate entries in the array is", dupeCount)
}
