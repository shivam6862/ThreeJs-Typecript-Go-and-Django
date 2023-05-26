package main
  
import "fmt"

type Finalsalary func(int, int) int

type Author struct {
    name      string
    language  string
    Marticles int
    Pay       int
    salary Finalsalary
}

func main() {
    result := Author{
        name:      "Sonia",
        language:  "Java",
        Marticles: 120,
        Pay:       500,
        salary: func(Ma int, pay int) int {
            return Ma * pay
        },
    }
    fmt.Println("Author's Name: ", result.name)
    fmt.Println("Language: ", result.language)
    fmt.Println("Total number of articles published in May: ", result.Marticles)
    fmt.Println("Per article pay: ", result.Pay)
    fmt.Println("Total salary: ", result.salary(result.Marticles, result.Pay))
}