package template_method

import "fmt"

func Example() {
	fqt := NewQuotes(&FrenchQuotes{})
	fmt.Println("French quotes: " + fqt.Quotes("Test String"))

	gqt := NewQuotes(&GermanQuotes{})
	fmt.Println("German quotes: " + gqt.Quotes("Test String"))
}
