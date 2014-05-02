package main

import (
	"code.google.com/p/log4go"
	"fmt"
//	"launchpad.net/goyaml"
	"net/http"
	"os"
	"github.com/rahulkj/hello"
)

// Declaring constants
const (
   PI = 3.14
   Language = "Go"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

type T struct {
	A string
	B []int
}

func main() {
	log := make(log4go.Logger)
	log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())

	http.HandleFunc("/", helloThere)
	http.HandleFunc("/all", all)
	http.HandleFunc("/pointer", PointerExample)
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	log.Debug("Listening at port %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func helloThere(res http.ResponseWriter, req *http.Request) {
	// Dump ENV
	fmt.Fprint(res, "Find the source code at : https://github.com/rahulkj/hello.git\n")
	fmt.Fprint(res, "Try appending the following to this URL: \n", "/all\n", "/pointer\n")
	fmt.Fprint(res, "Use the following command to push the code to cloudfoundy\n")
	fmt.Fprint(res, "gcf push server -b http://github.com/ryandotsmith/null-buildpack.git -c ./bin/server \n")
}

func ConstantUsage(res http.ResponseWriter, req *http.Request) {
    fmt.Println(res, PI, Language)
}

func StructureExample(res http.ResponseWriter, req *http.Request) {
    var p = hello.Person{"Mary", "Dis", 30}
    fmt.Fprintf(res, p.FName, p.LName, p.Age)

    // Update the fName of variable person
    p.FName = "Test"
    fmt.Fprintf(res, p.FName, p.LName, p.Age)

    // Update the values of variable person by passing in the reference
    hello.UpdatePerson(&p)
    fmt.Fprintf(res, "Updated Person", p.FName, p.LName, p.Age)

	p.Rename("RJ")
	fmt.Fprintf(res, p.FName, p.LName, p.Age)

}

func MathExample(res http.ResponseWriter, req *http.Request) {
    // short form of declaring int variables
    a, b, c := 1, 2, 3

    fmt.Println(a, b, c)

    // Perform addition of two variables
    fmt.Fprintf(res, "Sum of", a, "and", b, "is =", hello.Addition(a, b))

    // Perform math operations on two variables
    add, sub, multi, div := hello.Compute(c, b)
    fmt.Println(res, add, sub, multi, div)

    // Perform math operations on two variables - call EnhancedCompute
    add, sub, multi, div = hello.EnhancedCompute(a, b)
    fmt.Println(res, add, sub, multi, div)

}

func PointerExample(res http.ResponseWriter, req *http.Request) {
    // short form of declaring string variables
    message := "Test hello \n"
    // Declaring a variable of type Person

    var greeting *string = &message

    fmt.Fprintf(res, message)
    fmt.Fprintf(res, "Hello world! \n")

    fmt.Println(res, greeting, *greeting)
    fmt.Println(res, message, *greeting)

    *greeting = "Updated"
    fmt.Fprintf(res, message, *greeting)
}

func ClosureExample(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Rahul", hello.PrintNoLine)

	fmt.Fprintf(res, "Rahul", hello.PrintLine)
	fmt.Fprintf(res, "Rahul", hello.CustomPrintLine("LOL"))
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

func all(res http.ResponseWriter, req *http.Request) {

    PointerExample(res, req)

    MathExample(res, req)

    StructureExample(res, req)

    ConstantUsage(res, req)

    ClosureExample(res, req)

	chann := make(chan bool)

    persons := hello.Persons{{"Amanda", "X", 20}, {"Gerry", "Y", 32}, {"John", "Z", 45}}
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

