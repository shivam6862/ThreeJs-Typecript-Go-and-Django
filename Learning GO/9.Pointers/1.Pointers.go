package main
 
import "fmt"
 
type Employee struct {
    name  string
    empid int
}

func main() {
    x := 0xFF
    y := 0x9C
    fmt.Printf("Type of variable x is %T\n", x)
    fmt.Printf("Value of x in hexadecimal is %X\n", x)
    fmt.Printf("Value of x in decimal is %v\n", x)
    fmt.Printf("Type of variable y is %T\n", y)
    fmt.Printf("Value of y in hexadecimal is %X\n", y)
    fmt.Printf("Value of y in decimal is %v\n", y)   
     
	var x2 int = 5748
	var p *int
	p = &x2
	// var p = &x2
	//  p := &x2
	fmt.Println("Value stored in x = ", x2)
	fmt.Println("Address of x = ", &x2)
	fmt.Println("Value stored in variable p = ", p)

	var s *int
	fmt.Println("s = ", s)

    var y2 = 458
    var p2 = &y2
    fmt.Println("Value stored in y2 = ", y2)
    fmt.Println("Address of y2 = ", &y2)
    fmt.Println("Value stored in pointer variable p = ", p2)
    fmt.Println("Value stored in y2(*p2) = ", *p2)
	*p2 = 500
    fmt.Println("Value stored in y2(*p2) after Changing = ",y2)

	// Go Pointer to Pointer (Double Pointer)
	fmt.Println()
	var V int = 100
    var pt1 *int = &V
    var pt2 **int = &pt1
    fmt.Println("The Value of Variable V is = ", V)
    fmt.Println("Address of variable V is = ", &V)
    fmt.Println("The Value of pt1 is = ", pt1)
    fmt.Println("Address of pt1 is = ", &pt1)
    fmt.Println("The value of pt2 is = ", pt2)
    fmt.Println("Value at the address of pt2 is or *pt2 = ", *pt2)
    fmt.Println("*(Value at the address of pt2 is) or **pt2 = ", **pt2)


	var x3 = 100
    fmt.Printf("The value of x before function call is: %d\n", x3)
    var pa *int = &x3
    ptf(pa)
    fmt.Printf("The value of x after function call is: %d\n", x3)
	x3=200
	ptf(&x3)
    fmt.Printf("The value of x after function call is: %d\n", x3)
  
	n := rpf()
    fmt.Println("Value of n is: ", *n)

	arr := [5]int{78, 89, 45, 56, 14}
    updatearray(&arr)
    fmt.Println(arr)

	s2 := [5]int{78, 89, 45, 56, 14}
    updateslice(s2[:])
    fmt.Println(s2)

	// Pointer to a Struct in Golang
	emp := Employee{"ABC", 19078}
    pts := &emp
    fmt.Println(pts)
    fmt.Println(*pts)
    fmt.Println(pts.name)
    fmt.Println((*pts).name)
	pts.name = "XYZ"
    fmt.Println(pts)

	fmt.Println("Capacity of s2: ", cap(s2))
	fmt.Println("Capacity of s2: ", len(s2))
}
func ptf(a *int) {
    *a = 748
}
func rpf() *int {
    lv := 100
    return &lv
}
func updatearray(funarr *[5]int) {
    (*funarr)[4] = 750
}
func updateslice(funarr []int) {
    funarr[4] = 750
}