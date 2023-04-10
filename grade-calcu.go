package main

import (
	"fmt"
	"os"

	"github.com/inancgumus/screen"
)

func main() {
	var prelim, midterm, prefinal, final int

	for {
		fmt.Println("Enter Prelim grade: ")
		fmt.Scan(&prelim)
		fmt.Println("Enter Midterm grade: ")
		fmt.Scan(&midterm)
		fmt.Println("Enter Prefinal grade: ")
		fmt.Scan(&prefinal)
		fmt.Println("Enter Final grade: ")
		fmt.Scan(&final)
		average := (prelim + midterm + prefinal + final) / 4
		fmt.Printf("Your Average is %d \n", average)

		if average >= 75 {
			fmt.Println("You passed!")
		} else {
			fmt.Println("You failed!")
		}

		fmt.Println("Do you want to compute again? y/n")

		var answer string
		fmt.Scan(&answer)

		if answer != "y" {
			// Clear the console screen
			screen.Clear()
			os.Exit(0)
		} else {
			// Clear the console screen
			screen.Clear()
		}
	}
}
