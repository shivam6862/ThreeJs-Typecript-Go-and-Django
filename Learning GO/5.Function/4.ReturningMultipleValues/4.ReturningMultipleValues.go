package main
  
import "fmt"
  
// myfunc return 3 values of int type
func myfunc(p , q int)(int, int, int ){
    return p - q, p * q, p + q 
}

func myfunc2(p, q int)( rectangle int, square int ){
	fmt.Println(p*q)
    rectangle = p*q
    square = p*p
    return  
}

// Main Method
func main() {
	var myvar1, myvar2, myvar3 = myfunc(4, 2)
	fmt.Printf("Result is: %d", myvar1 )
	fmt.Printf("\nResult is: %d", myvar2)
	fmt.Printf("\nResult is: %d", myvar3)
	fmt.Println()
	
	// LIFO(Last-In, First-Out) 
	defer fmt.Println("End1")
	defer fmt.Println("End2")
	defer myfunc2(23, 56)
	defer fmt.Println("End3")
	// Named Return Values
   var area1, area2 = myfunc2(2, 4)
   // Display the values
   fmt.Printf("Area of the rectangle is: %d", area1 )
   fmt.Printf("\nArea of the square is: %d", area2)
   fmt.Println()
}