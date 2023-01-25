package main

import (
	"fmt"
)

func main() {
	// Regular declaration, but the size is twice
	grades :=  [3]int{97, 85, 93}

	// If you initialize it, the size is optional
	grades2 := [...]int{97, 85, 93}

	// Empty array
	var students [3]string
		
	fmt.Printf("Grades: %v (this is a type: %T)\n", grades, grades)
	fmt.Printf("Grades: %v (this is a type: %T)\n", grades2, grades2)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Marcin"
	students[1] = "Lisa"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))
}
