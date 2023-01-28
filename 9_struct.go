package main

import (
	"fmt"
)

type Doctor struct {
	// field names with lowercase are not exported
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	// Positional syntax also available, but it's shitty when you change the struct
	// e.g., change the order, add some fields etc.
	aDoctor2 := Doctor{
		4,
		"Tom Baker",
		[]string{
			"Sarah Jane Smith",
			"Harry Sullivan",
			"Leela",
			"K-9",
			"Romana",
			"Adric",
			"Nyssa",
			"Tegan Jovanka",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	fmt.Println(aDoctor2)

	// Anonymous structs are also possible
	aDoctor3 := struct{ name string }{name: "John Pertwee"}
	// Struct are value types, so not reference ones like splices/maps
	anotherDoctor := aDoctor3
	// accessing by reference is done as usual
	aDoctorRef := &aDoctor3
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor3)
	fmt.Println(anotherDoctor)
	fmt.Println(aDoctorRef)
}
	