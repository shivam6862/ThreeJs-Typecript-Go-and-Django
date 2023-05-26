package main
  
import (
    "fmt"
    "time"
)
  
func display(str string) {
    for w := 0; w < 6; w++ {
        time.Sleep(1 * time.Millisecond)
        fmt.Println(w)
    }
}
  
func main() {
    go display("Welcome")
    // display("Shivam")

	fmt.Println("Welcome!! to Main function")
    go func() {
        fmt.Println("Welcome!! to GeeksforGeeks")
    }()
    time.Sleep(50* time.Millisecond)
    fmt.Println("GoodBye!! to Main function")
}