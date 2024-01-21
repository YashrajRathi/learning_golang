package main

import (
	"fmt"
        //This will path of your model package
	"learn/model"
)

//Test function
func main() {
	//STRUCTURE IDENTIFIER
	p := &model.Person{
		Name: "test",
		Age:  21,
	}
	fmt.Println(p)
	c := &model.Company{}
	fmt.Println(c)

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}