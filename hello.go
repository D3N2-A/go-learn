package main

import (
	"fmt"
)


func main(){
	// Declaring variable with datatype
	var age int 
	// Initializing that variable
	age = 10
	// Walrus operator -> Auto infer type
	age1 := 20
	// Operators + can join strings
	str1  := "Hello,"
	str2  := "World!"
	fmt.Println(str1+" "+str2)

	// ++ -- -> incr or decr but can't prepend
	// && || ! AND, OR, NOT
	age++
	age--

	// Strings -> defined only by double quotes
	var name string = "String -> Hello, World!"
	fmt.Println(name, len(name))
	// Slice [:]
	name = "Var reassign"
	fmt.Println(name[:])


	fmt.Println("Hello, World!", age, age1)
}