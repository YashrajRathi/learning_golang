package main

import (
	// "errors"
	"bufio"
	"os"
	"fmt"
	// "reflect"
	// "strconv"
	// "strings"
)

func convert_string_to_char_array(s string) []int32 {
	return_array := make([]int32,len(s))
	for a,b := range s {
		return_array[a] = b
	}
	return return_array
}

func convert_char_array_to_string(arr []int32) (outString string) {
	for _,cc := range arr {
		outString += string(cc)
	}
	return outString
}

func permute(given_string string,target_index int) {
	if((target_index) == len(given_string)){
		fmt.Println(given_string)
	} else {
		char_arr := convert_string_to_char_array(given_string)
		for i:= target_index ; i <	 len(given_string) ; i++ {
			char_arr = swap(target_index,i,char_arr)
			same_string := convert_char_array_to_string(char_arr)
			permute(same_string,target_index + 1)
			char_arr = swap(target_index,i,char_arr)
		}
	}

}

func swap(x,y int,arr []int32) []int32{
	dummy := arr[x]
	arr[x] = arr[y]
	arr[y] = dummy
	return arr
}



func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the string that you want to see permutations of >")
	for scanner.Scan() {
		var target_string string
		target_string = scanner.Text()
		permute(target_string,0)
		fmt.Print("Enter the string that you want to see permutations of >")
	}
	// test := "abc"
	// // fmt.Println(print_string_from_char_array(convert_char_array(test)))
	// permute(test,0)

	// arr:= convert_string_to_char_array(test)
	// fmt.Println(convert_char_array_to_string(swap(0,1,arr)))

	// permute(test,2)
}