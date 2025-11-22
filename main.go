/*
 * @author Marko Skorupan
 * @version 1.0.0
 *@date 2025-11-21
 * @fileoverview "Sweater Colour Selector Program."
 */

package main

import "fmt"

func main () {
	var colour string

	//INPUT
	fmt.Print("Please choose a sweater colour from the available choices: Blue, Black, Red, White. ")
	fmt.Scan(&colour)

	//PROCESS
	if colour == "black" || colour == "white" {
		fmt.Println("You have enough sweaters in this colour")
	}else if colour == "red" {
		fmt.Println("This colour will look good on you.")
		fmt.Println("How about a pair of jeans to go with the sweater?")
	}else if colour == "blue" {
		fmt.Println("This colour doesn't do well with your eyes.")
	} else {
		fmt.Println("Your colour choice is invalid.")
	}
}
