package main

import (
	"fmt"
	"strings"
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
	// string in Go is a sequesnce of byte values
	var name string = "String -> Hello, World!"
	fmt.Println(name, len(name))
	// Slice [:]
	name = "Var reassign"
	fmt.Println(name[:])
	// Strings are immutable you can not update value but can reassign. Also they are reference type so so string is copied not its value. 
	var first = "test1"
	var second = first
	second = "test2"
	first = "123"
	fmt.Println(first)
	fmt.Println(second)


	// strings utlity package
	str3 := "aquickbrownfoxjumpedoveralazydog"
	ct := strings.Count(str3, "ui")
	fmt.Println(ct)
	


	fmt.Println("Hello, World!", age, age1)
}