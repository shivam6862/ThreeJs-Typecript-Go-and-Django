package main
 
import "fmt"

type Author struct {
    name   string
    branch string
    year   int
}
type HR struct {
    details Author
}


type Student struct {
    name   string
    branch string
    year   int
}
type Teacher struct {
    name    string
    subject string
    exp     int
    details Student
}

func main() {
    result1 := HR{
        details: Author{"Sona", "ECE", 2013},
    }
    fmt.Println("\nDetails of Author")
    fmt.Println(result1)


	result2 := Teacher{
        name:    "Suman",
        subject: "Java",
        exp:     5,
        details: Student{"Bongo", "CSE", 2},
    }
	fmt.Println("\nDetails of Teacher")
    fmt.Println(result2)
	fmt.Println("Student's name: ", result2.details.name)
}