package main

import (
	"fmt"
)


func main(){
	// Declare variables in the beginning
	var main_data int 
	var pointer_variable *int
	// var dereferencing_value &int

	// Assign them values
	main_data = 9
	pointer_variable = &main_data
	// dereferencing_value = *main_data

	// Attempt to print the main variable
	fmt.Println(main_data) // prints the obvious value [B]
	fmt.Println(&main_data) // prints the memory address of the variable [A]
	// fmt.Println(*main_data) // Not allowed since the int value doesn't contain a valid memory address

	// Attempt to print memory addresses
	fmt.Println(&pointer_variable) // prints the memory address of pointer variable
	fmt.Println(pointer_variable) // prints the memory address of the earlier variable [A]
	fmt.Println(*pointer_variable) // prints the value of the variable [B]
}
