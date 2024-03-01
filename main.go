package main

import "fmt"

func main() {
	err := printPattern(10)
	if err != nil {
		fmt.Println(err)
	}
}

func printPattern(length int) error {
	if length%2 == 0 {
		return fmt.Errorf("length must be an odd number")
	}
	middle := length / 2

	for i := 0; i < length; i++ {
		fmt.Println()

		for j := 0; j < length; j++ {
			if i == middle {
				fmt.Print("* ")
			} else if j == 0 || j == length-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("= ")
			}
		}
	}

	return nil
}
