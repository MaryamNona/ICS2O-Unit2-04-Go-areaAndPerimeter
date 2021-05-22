// Created by: Maryam Nona
// Created on: May 2021
//
// This program does basic math

package main

import (
	"fmt"
)

func main() {
	// This function does addition
	var length int
	var width int

	// input
	fmt.Println("This program finds the area and perimeter of a rectangle")
	fmt.Println()
	fmt.Print("Enter the length (mm): ")
	fmt.Scanln(&length)
	fmt.Println("Enter the width (mm): ")
	fmt.Scanln(&width)

	// output
	fmt.Println("The area is: ", length*width, "mm²")
	fmt.Println("The perimeter is: ", (length+width)*2, "mm")
	fmt.Println("Done.")
}
