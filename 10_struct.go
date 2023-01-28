package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // This is an embeded struct
	SpeedKPH float32
	CanFly   bool
}

// You can also limit fields, by defining tags
// Format is conveniont based, so it might be messy
// To get the tag you need reflection (sigh)
type Dragon struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	// Go does not support inheritance
	// But there is composition via embedding
	// The objects are not interchangable so Bird is not a type of Animal in this example

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b.Name)

	// In case of a literal syntax, you need to explicitly define the internal structs (like the Animal one)
	b2 := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b2.Name)

	// Embedding is not a good way to model common behavior (interfaces are a better fit)

	t := reflect.TypeOf(Dragon{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
