package newprinter

import "fmt"

// Closures
type Printer func(string) ()

func PrintNoLine (s string) {
    fmt.Print(s)
}

func PrintLine (s string) {
    fmt.Println(s)
}

func CustomPrintLine (custom string) Printer {
    return func (s string) {
        fmt.Println(s + " " + custom)
    }
}

func Print(message string, printer Printer) {
    printer(message)
}
