package section4

func ExampleArrayFilterBySubstring() {
	var elements = []string{"Hello", "world", "bla", "data", "bla", "foo", "boo"}
	search := "a"

	ArrayFilterBySubstring(elements, search)
	// Output:
	// bla
	// data
	// bla

}
