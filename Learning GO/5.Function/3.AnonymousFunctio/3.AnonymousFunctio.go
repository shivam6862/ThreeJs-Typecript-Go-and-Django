// AnonymousFunctio.go
package main
  
import (
  "fmt"
  "sort"
  "strings"
  "time"
)
 
func main() { 
    // Anonymous function
   func(){
      fmt.Println("Welcome! to AnonymousFunctio.go")
    }()

    value := func(){
    fmt.Println("Welcome! to AnonymousFunctio")
    }
    value()

    // Passing arguments in anonymous function
    func(data string){
    fmt.Println(data)
    }("AnonymousFunction")


    // Sorting the given slice
    s := []int{345, 78, 123, 10, 76, 2, 567, 5}
    sort.Ints(s)
    fmt.Println("Sorted slice: ", s)
    // Finding the index
    fmt.Println("Index value: ", strings.Index("GeeksforGeeks", "ks"))
    // Finding the time
    fmt.Println("Time: ", time.Now())
    fmt.Println("Time: ", time.Now().Unix())

    var Z int = 10
    fmt.Printf("Before Function Call, value of Z is = %d", Z)
    // call by Reference
    // by passing the address
    // of the variable Z
    modify(&Z)
    fmt.Printf("\nAfter Function Call, value of Z is = %d", Z)
}

func init() {
  fmt.Println("Welcome to init() function")
}
func modify(Z *int) {
  *Z = 70
}