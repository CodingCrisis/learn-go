func main() {
	// Regular declaration with initialization
	var age int = 41

	// Declare with inferred type
	var doprint = true

	// Declare two variables at the same time
	var like1, like2 = "playing cards", "synthwave music"
	// Shorthand declaration with initialization
	name := "Marcin"

	// Update value
	age = 41

	if doprint {
		fmt.Printf("My name is %s and I'm %d\nI like %s and %s\n", name, age, like1, like2)
	}

	// Variables are zero initialized
	var emptyint int
	var emptystring string
	var istrue bool

	fmt.Println(emptyint, emptystring, istrue)
}
