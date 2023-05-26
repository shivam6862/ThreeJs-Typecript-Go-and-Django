package main
 
import "fmt"
// Author structure
type author struct {
    name   string
    branch string
}
 
// Method with a pointer receiver of author type
func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}
 
// Method with a value receiver of author type
func (a author) show_2() {
    a.name = "Shyam"
    fmt.Println("Author's name(Before) : ", a.name)
}

func main() {
    res := author{
        name:   "Shivam Kumar",
        branch: "ME",
    }
    fmt.Println("Branch Name(Before): ", res.branch)
    // Calling the show_1 method (pointer method) with value
    res.show_1("ECE")
    fmt.Println("Branch Name(After): ", res.branch)
    // Calling the show_2 method (value method) with a pointer
    (&res).show_2()
    fmt.Println("Author's name(After): ", res.name)
}