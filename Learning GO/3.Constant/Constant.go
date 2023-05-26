package main
 
import (
	"fmt"
	"math"
)
 
const PI = 3.14
 
func main(){
    const GFG = "GeeksforGeeks"
    fmt.Println("Hello", GFG)
    fmt.Println("Happy", PI, "Day")
    const Correct= true
    fmt.Println("Go rules?", Correct)

	const A = "GFG"
    var B = "GeeksforGeeks"
    // Concat strings.
    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld)
    // Compare strings.
    fmt.Println(A == "GFG")  
    fmt.Println(B < A)

	const trueConst = true
    // Type definition using type keyword
    type myBool bool   
    var defaultBool = trueConst // allowed
    var customBool myBool = trueConst // allowed
    //  defaultBool = customBool // not allowed
    fmt.Println(defaultBool)
    fmt.Println(customBool) 

	const s string = "GeeksForGeeks"
	fmt.Println(s)
    const n = 5
    const d = 3e10 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}