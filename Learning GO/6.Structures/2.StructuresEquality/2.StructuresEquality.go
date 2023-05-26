package main
 
import (
    "fmt"
    "reflect"
)
 
type Author struct {
    name      string
    branch    string
    language  string
    Particles int
}

func main() {
    a1 := Author{
        name:      "Moana",
        branch:    "CSE",
        language:  "Python",
        Particles: 38,
    }
    a2 := Author{
        name:      "Moana",
        branch:    "CSE",
        language:  "Python",
        Particles: 38,
    }
    a3 := Author{
        name:      "Dona",
        branch:    "CSE",
        language:  "Python",
        Particles: 38,
    }
	if a1 == a2 {
        fmt.Println("Variable a1 is equal to variable a2")
    } else {
        fmt.Println("Variable a1 is not equal to variable a2")
    }
    if a2 == a3 {
        fmt.Println("Variable a2 is equal to variable a3")
         
    } else {
        fmt.Println("Variable a2 is not equal to variable a3")
    }
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))
    fmt.Println("Is a2 equal to a3: ", reflect.DeepEqual(a2, a3))
}