package Advanced

import "fmt"

func AdultYears() {
	age := 32

	// result := calcAdultYears(&age)
	editAdultYears(&age)
	fmt.Println(age)
	// fmt.Println("Current age is:",age,"and Adult Age is: ", result)
}

func editAdultYears(age *int) {
	*age -= 18
}
