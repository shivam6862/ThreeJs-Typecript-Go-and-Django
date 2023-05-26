package main

import (
    "fmt"
    "strings"
)
// Variadic function to join strings
func joinstr(elements ...string) string {
    return strings.Join(elements, "-")
}
 
func main(){
 
    // zero argument
    fmt.Println(joinstr())
 
    // multiple arguments
    fmt.Println(joinstr("GEEK", "GFG"))
    fmt.Println(joinstr("Geeks", "for", "Geeks"))
    fmt.Println(joinstr("G", "E", "E", "k", "S"))


	// pass a slice in variadic function
    elements := []string{"geeks", "FOR", "geeks"}
    fmt.Println(joinstr(elements...))
}