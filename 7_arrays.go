package main

import (
	"fmt"
)

func main() {
	// Regular declaration, but the size is twice
	grades := [3]int{97, 85, 93}

	// If you initialize it, the size is optional
	grades2 := [...]int{97, 85, 93}

	// Empty array
	var students [5]string

	fmt.Printf("Grades: %v (this is a type: %T)\n", grades, grades)
	fmt.Printf("Grades: %v (this is a type: %T)\n", grades2, grades2)
	fmt.Printf("Students: %v\n", students)
	students[0] = "Marcin"
	students[1] = "Lisa"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	// Array of arrays, plus initialization
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	a := [...]int{1, 2, 3}
	// This copies the whole array, not a reference
	b := a
	// This is  a reference
	c := &a

	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Slices

	slicea := []int{1, 2, 3}
	sliceb := slicea
	// Slice does not copy data as an array, so all references will change
	sliceb[1] = 5
	fmt.Println(slicea)
	fmt.Println(sliceb)
	fmt.Printf("Length: %v\n", len(slicea))
	fmt.Printf("Capacity: %v\n", cap(slicea))

	sa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sb := sa[:]   // slice of all elements
	sc := sa[3:]  // slice from 4h element to end
	sd := sa[:6]  // slice frist 6 elements
	se := sa[3:6] // slice including 4th to 6th elements
	// remember, they all point to the same underlying data
	sa[5] = 42
	fmt.Println(sa)
	fmt.Println(sb)
	fmt.Println(sc)
	fmt.Println(sd)
	fmt.Println(se)

	// Create a slice via make
	ma := make([]int, 3)
	fmt.Println(ma)
	fmt.Printf("Length: %v\n", len(ma))
	fmt.Printf("Capacity: %v\n", cap(ma))

	// And set the underlying table to be 100 elements long
	ma1 := make([]int, 3, 100)
	fmt.Println(ma1)
	fmt.Printf("Length: %v\n", len(ma1))
	fmt.Printf("Capacity: %v\n", cap(ma1))

	// Normal empty slice
	aa1 := []int{}
	fmt.Println(aa1)
	fmt.Printf("Length: %v\n", len(aa1))
	fmt.Printf("Capacity: %v\n", cap(aa1))

	// We can append an item (it actually creates a new array underneath)
	// This can get quite expensive with bigger arrays, hence the make function
	aa1 = append(aa1, 1)
	fmt.Println(aa1)
	fmt.Printf("Length: %v\n", len(aa1))
	fmt.Printf("Capacity: %v\n", cap(aa1))

	// Not sure how the extra space in the array is handled
	// in an Internet exapmle it got to cap 8, mine was 6 (so 5 + 1 extra)
	aa1 = append(aa1, 2, 3, 4, 5)
	fmt.Println(aa1)
	fmt.Printf("Length: %v\n", len(aa1))
	fmt.Printf("Capacity: %v\n", cap(aa1))

	// Interesting feature - the ... will spread an array
	aa1 = append(aa1, []int{6, 7, 8, 9, 10}...)
	fmt.Println(aa1)
	fmt.Printf("Length: %v\n", len(aa1))
	fmt.Printf("Capacity: %v\n", cap(aa1))

	// Removing leading/trailing elements from a slice
	aa2 := []int{1, 2, 3, 4, 5}
	ab2 := aa2[1 : len(aa2)-1]
	fmt.Println(aa2)
	fmt.Println(ab2)
	// Removing a central element is a bit more tricky
	// WARNING!!! this will fuck up any other references to the underlying array
	ac2 := append(aa2[:2], aa2[3:]...)
	fmt.Println(ac2)
	fmt.Println(aa2)
}
