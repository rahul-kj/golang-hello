package person

import "fmt"

// Define a structure type called Person
type Person struct {
   FName string
   LName string
   Age int
}

type Persons []Person

// Function that updates the values of the passed in pointer of the variable
func UpdatePerson(person *Person) {
   fmt.Println(person.FName, person.LName, person.Age)

   person.Age = 40
   fmt.Println(person.FName, person.LName, person.Age)
}

func (person *Person) Rename(fName string) {
	person.FName = fName
}

func (persons Persons) BranchingExample() {
	// Use len function to get the size of the slice/array
	fmt.Println("Length of the persons slice", len(persons))

	if persons[1].FName == "Gerry" {
		fmt.Println("Yes Gerry is here")
	} else {
		fmt.Println("Oops Gerry is not here")
	}

	// Conditional case which has no condition on switch, but has on case that returns a boolean
	for _, p := range persons {
		switch {
		case p.Age <= 20 :
			fmt.Println(p.FName, "belongs to Group 1")
		case p.Age <= 30, p.Age > 40 :
			fmt.Println(p.FName, "belongs to Group 2")
		case p.Age <= 40 :
			fmt.Println(p.FName, "belongs to Group 3")
		default :
			fmt.Println(p.FName, "belongs to Group X")
		}
	}
}
