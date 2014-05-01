package main

import (
	"fmt"
	"launchpad.net/goyaml"
	"net/http"
	"os"
	"local/rahulkj/person"
	"local/rahulkj/newprinter"
	"local/rahulkj/newmath"
)

// Declaring constants
const (
   PI = 3.14
   Language = "Go"
)

func main() {

    PointerExample()

    MathExample()

    StructureExample()

    ConstantUsage()

    ClosureExample()

	chann := make(chan bool)

    persons := person.Persons{{"Amanda", "X", 20}, {"Gerry", "Y", 32}, {"John", "Z", 45}}
    go func() {
		persons.BranchingExample()
		chann <- true
	}()

	persons.BranchingExample()

    TypeSwitchExample(3.20)

    fmt.Println(MapExample("US"))

	SliceExample()

	<- chann

}

func ConstantUsage() {
    fmt.Println(PI, Language)
}

func StructureExample() {
    var p = person.Person{"Mary", "Dis", 30}
    fmt.Println(p.FName, p.LName, p.Age)

    // Update the fName of variable person
    p.FName = "Test"
    fmt.Println(p.FName, p.LName, p.Age)

    // Update the values of variable person by passing in the reference
    person.UpdatePerson(&p)
    fmt.Println("Updated Person", p.FName, p.LName, p.Age)

	p.Rename("RJ")
	fmt.Println(p.FName, p.LName, p.Age)

}

func MathExample() {
    // short form of declaring int variables
    a, b, c := 1, 2, 3

    fmt.Println(a, b, c)

    // Perform addition of two variables
    fmt.Println("Sum of", a, "and", b, "is =", newmath.Addition(a, b))

    // Perform math operations on two variables
    add, sub, multi, div := newmath.Compute(c, b)
    fmt.Println(add, sub, multi, div)

    // Perform math operations on two variables - call EnhancedCompute
    add, sub, multi, div = newmath.EnhancedCompute(a, b)
    fmt.Println(add, sub, multi, div)

}

func PointerExample() {
    // short form of declaring string variables
    message := "Test hello \n"
    // Declaring a variable of type Person

    var greeting *string = &message

    fmt.Printf(message)
    fmt.Printf("Hello world! \n")

    fmt.Println(greeting, *greeting)
    fmt.Println(message, *greeting)

    *greeting = "Updated"
    fmt.Println(message, *greeting)
}

func ClosureExample() {
	newprinter.Print("Rahul", newprinter.PrintNoLine)

	newprinter.Print("Rahul", newprinter.PrintLine)
	newprinter.Print("Rahul", newprinter.CustomPrintLine("LOL"))
}

func TypeSwitchExample(x interface{}) {
    switch x.(type) {
        case int: fmt.Println("int")
        case string: fmt.Println("string")
        default: fmt.Println("unknown")
    }
}

func MapExample(countryName string) string {
    countries := map[string]string {
        "CH" : "China", "JP" : "Japan",
    }
//    var countries map[string]string
//    countries = make(map[string]string)
    countries["US"] = "United States"
    countries["AU"] = "Australia"
    countries["IN"] = "India"

//    countries["JP"] = "", false
    delete(countries, "JP")
    fmt.Println(countries)

	if value, exists := countries[countryName]; exists {
		fmt.Println("Found the country", countryName)
		return value
	}

    return countries[countryName]
}

func SliceExample() {
	var s []int
	s = make([]int, 3)
	s[0] = 1
	s[1] = 10
	s[2] = 500

	// s := []int {1, 10, 500}
	fmt.Println(s)

	s = s[1:2]
	fmt.Println(s)

	// Adding element to a slice
	s = append(s, 20, 40)
	fmt.Println(s)

	// Deleting an element from a slice
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
}

type T struct {
	A string
	B []int
}

func hello(res http.ResponseWriter, req *http.Request) {
	// Dump ENV
	fmt.Fprint(res, "ENV:\n")
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintln(res, e)
	}
	fmt.Fprint(res, "\nYAML:\n")

	//Dump some YAML
	t := T{A: "Foo", B: []int{1, 2, 3}}
	if d, err := goyaml.Marshal(&t); err != nil {
		fmt.Fprintf(res, "Unable to dump YAML")
	} else {
		fmt.Fprintf(res, "--- t dump:\n%s\n\n", string(d))
	}
}