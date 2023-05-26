package main
  
import "fmt"

type details struct {
    name   string
    age    int
    gender string
}
type student struct {
    branch string
    year   int
    details
}

// Funcuntion value will print
func (d student) promotmethod(number int) int {
	return d.age * number*2
}
func (d details) promotmethod(number int) int {
	return d.age * number
}

func main() {
    values := student{
        branch: "CSE",
        year:   2010,
        details: details{
            name:   "Sumit",
            age:    28,
            gender: "Male",
        },
    }
    fmt.Println("Name: ", values.name)
    fmt.Println("Age: ", values.age)
    fmt.Println("Gender: ", values.gender)
    fmt.Println("Year: ", values.year)
    fmt.Println("Branch : ", values.branch)

	fmt.Println("promotmethod value: ", values.promotmethod(20))
}