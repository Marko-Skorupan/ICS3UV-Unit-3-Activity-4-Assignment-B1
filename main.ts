/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview "Sweater Colour Selector Program."
 */

//INPUT
let colour: string | null = prompt("Please choose a sweater colour from the available choices: Blue, Black, Red, White.");
let choice = (colour ?? "").toLowerCase();

//PROCESS
if (choice == "black" || choice == "white") {
  console.log("You have enough sweaters in this colour");
} else if (choice === "red") {
  console.log("This colour will look good on you.");
  console.log("How about a pair of jeans to go with the sweater?");
} else if (choice === "blue") {
  console.log("This colour doesn't do well with your eyes.");
}else {
  console.log("Your colour choice is invalid.");
}