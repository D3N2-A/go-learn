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

	// Array var varName [size]dataType
	// Array var varName = [size]dataType{initialize with items or omit brackets for just declaration}
	var fixed = [3]string{"Hello", " ", ", World!"}
	var counted = [...]string{"Test"}
	fmt.Println(fixed, counted)
	// Array cannot be resized, you have to exlicitly define length of array. 
	// This is a big limitation so we rarely use array, we use slizes
	// Arrays are value types -> passing it to functions, returnging from a function creates a copy of original array


	// Slices -> Data structure similar to array but can change size
	// Slice is defined similar to array, you just omit sizze
	// You can create empty slice using make -> Make returns type of arg not pointer to it (unlike new)

	var mySlice = make([]string, 3)
	mySlice[0] = "Hello"
	mySlice[1] = " "
	mySlice[2] = ", World!"
	fmt.Println(mySlice)
	mySlice = append(mySlice, " This is great")
	fmt.Println(mySlice)
	// You can create a slice with fixed size -> Good practice since extending slice requires allocating new memory and copying content
	var newSlice = make([]string, 0, 10) // Initializes slice with 0 empty elements, memory allocation for 10 elements, for 10 empty element ommit size arg
	newSlice[0] = "Hello, World!"
	// You can create slice from array [:] operator



	


	fmt.Println("Hello, World!", age, age1)
}