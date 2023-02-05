package main

import (
	"fmt"
)

func point1() {
	a := 42
	// this copies the value
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
}

func point2() {
	var a int = 42
	var b *int = &a
	// compare addresses, get address via &
	fmt.Println(&a, b)
	a = 27
	// dereference a pointer via *
	fmt.Println(a, *b)
	// directly change value
	*b = 14
	fmt.Println(a, *b)
}

func point3() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	// Pointer arithmetic is not allowed
	// so the following would fail
	// c := &a[1] - 4
	// to point to a[0]
	// there's an option to use the "unsafe" package if necessary
	fmt.Printf("%v %p %p\n", a, b, c)
}

func point4() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	// ms is storing a pointer with an address to the sctruct
	fmt.Println(ms)

	// alternative is to use the new keyword
	// but it's zero initialized, so <nil> for starters
	// and then a struct with a zero when object is created
	// as it's a struct containing an int
	var ms2 *myStruct
	fmt.Println(ms2)
	ms2 = new(myStruct)
	fmt.Println(ms2)
	// the parentheses is necessary for operation order - dot would go before star
	(*ms2).foo = 42
	fmt.Println(ms2, (*ms).foo)
	// compiler knows a shorter version though
	// so this does not try to access field of a pointer, but the underlying object
	fmt.Println(ms.foo)
}

func point5() {
	// in case of a table a copy of the object is done
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	// slice is already a reference, so assigning s to t is also copying a reference
	// hence s[1] = 42 affects both s and t
	s := []int{1, 2, 3}
	t := s
	fmt.Println(s, t)
	s[1] = 42
	fmt.Println(s, t)

	// same goes for maps, also a reference type
	m := map[string]string{"foo": "bar", "baz": "buz"}
	n := m
	fmt.Println(m, n)
	m["foo"] = "qux"
	fmt.Println(m, n)
}

func main() {
	point1()
	point2()
	point3()
	point4()
	point5()
}

type myStruct struct {
	foo int
}
