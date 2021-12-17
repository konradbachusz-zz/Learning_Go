package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Konrad"
	fmt.Println(firstName)

	ptr := &firstName 

	
	fmt.Println(ptr,*ptr)
}