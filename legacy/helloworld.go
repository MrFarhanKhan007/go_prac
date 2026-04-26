package main

import (
	"fmt"
)

func main()  {
	fmt.Println("I am farhan")
	// This is a comment
	// This is another comment

	var name string = "Farhan"
	fmt.Println("My name is", name)

	var age int = 25
	fmt.Println("My age is", age)

	var isEmployed bool = true
	fmt.Println("Am I employed?", isEmployed)


	name = "Farhan2"
	fmt.Println("My name is", name)

	age= 26
	fmt.Println("My age is", age)

	isEmployed= false
	fmt.Println("Am I employed?", isEmployed)

	location:= "New York"
	fmt.Println("My location is", location)

	// multiple variable declaration
	var (
		name1     string
		age1      int
		isEmployed1 bool
		location1 string
	)

	name1 = "Farhan1"
	age1 = 30
	isEmployed1 = true
	location1 = "Los Angeles"

	fmt.Println("My name is", name1)
	fmt.Println("My age is", age1)
	fmt.Println("Am I employed?", isEmployed1)
	fmt.Println("My location is", location1)

	main2()
	main3()


	var name string = "farhan"
	name2 :="farhan"
}
