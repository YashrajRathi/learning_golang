package main

import (
	"fmt"
	"os"
)

func convert_string_to_char_array(s string) (*[]int32,int) {
	return_array := make([]int32,len(s))
	for a,b := range s {
		return_array[a] = b
	}

	return &return_array,len(s)
}

func convert_char_array_to_string(arr *[]int32,len int) (outString string) {
	for i:=0;i<len;i++ {
		outString += string((*arr)[i])
	}
	return outString
}

func swap(x,y int,arr *[]int32){
	dummy := (*arr)[x]
	(*arr)[x] = (*arr)[y]
	(*arr)[y] = dummy
}

func pointer_permute(char_arr *[]int32,target_index int,len int) {
	if((target_index) == len){
		fmt.Println(convert_char_array_to_string(char_arr,len))
	}else{
		for i:= target_index ; i <len ; i++ {
			swap(target_index,i,char_arr)
			pointer_permute(char_arr,target_index+1,len)
			swap(target_index,i,char_arr)
		}
	}
}

func main(){
	var target_string string
	target_string = os.Args[1]
	var tt,len = convert_string_to_char_array(target_string)
	pointer_permute(tt,0,len)
}