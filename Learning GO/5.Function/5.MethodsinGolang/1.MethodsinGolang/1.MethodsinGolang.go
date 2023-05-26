package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    particles int
    salary    int
}
 
// Method with a receiver of author type
func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.particles)
    fmt.Println("Salary: ", a.salary)
}

// Type definition
type data int
 
// Defining a method with non-struct type receiver
func (d1 data) multiply(d2 data) data {
    return d1 * d2
}

func main() {
    // Initializing the values of the author structure
    res := author{
        name:      "Shashank",
        branch:    "ME",
        particles: 21117119,
        salary:    334000,
    }

    res.show()


	value1 := data(23)
    value2 := data(20)
    res1 := value1.multiply(value2)
    fmt.Println("Final result: ", res1)
}