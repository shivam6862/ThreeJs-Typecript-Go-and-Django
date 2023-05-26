package main
import "fmt"

func area(length, width int)int{
    Ar := length* width
    return Ar
}

func swap(a, b int)int{
    var o int
    o= a
    a=b
    b=o
	return o
}

// function which swap values
func swap1(a, b *int)int{
	var o int
    o = *a
    *a = *b
    *b = o
   return o
}
  

func main() {
   // with method calling
   fmt.Printf("Area of rectangle is : %d", area(12, 10))

   // call by values
    var p int = 10
	var q int = 20
	fmt.Println()
	fmt.Printf("p = %d and q = %d", p, q)
	swap(p, q)
	fmt.Printf("\np = %d and q = %d",p, q)
	
	
	// call by reference
	var p1 int = 10
	var q1 int = 20
	fmt.Println()
   	fmt.Printf("p = %d and q = %d", p1, q1)
   	swap1(&p1, &q1)
	fmt.Printf("\np = %d and q = %d",p1, q1)
}