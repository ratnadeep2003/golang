// Every executable Go program must belong to the "main" package.
// This tells the compiler: "This is a standalone program, not a library."
package main

// "fmt" is imported so we can use formatting and printing functions.
import "fmt" //fmt = format

// main is the special entry point function. 
// When you run the program, execution starts right here inside main().
func main() {
    // Println means "Print Line". It prints the text to the console 
    // and automatically adds a newline at the end.
    fmt.Println("hello from go program")
}