package main

import (
	"fmt"
)


func increment(pointer *int){
	*pointer = *pointer + 1
}

func increment_with_declared_variable(pointer *int){
	dd := *pointer
	dd = dd + 1
}

func main(){

	ii := 10
	pp := &ii
	increment_with_declared_variable(pp) //Does nothing to orignal variable
	fmt.Println(ii)
	fmt.Println(pp)

	increment(pp) //Actually increments the variable
	fmt.Println(ii)
	fmt.Println(pp)

}

func print(dd int){
	fmt.Println(dd)
}