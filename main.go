package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func add(a,b int) string {
	return strconv.Itoa(a + b)
}

// func main(){
// 	fmt.Println(add(1,2))
// }

func hello(w http.ResponseWriter, req *http.Request) {
	num1,_:=strconv.Atoi(req.URL.Query().Get("num1"))
	num2,_:=strconv.Atoi(req.URL.Query().Get("num2"))
    fmt.Fprintf(w, add(num1,num2))
}

func main() {

    http.HandleFunc("/add", hello)

    http.ListenAndServe(":8090", nil)
}