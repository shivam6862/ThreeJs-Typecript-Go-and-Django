package main
import "fmt"
 
func main() {
   // taking a local variable
   var v1 int = 700
   // checking the condition
   if v1 == 100 {
      // if condition is true then
      // display the following */
      fmt.Printf("Value of v1 is 100\n")
   } else if v1 == 200 {
      fmt.Printf("Value of a is 20\n")
   } else if v1 == 300 {
      fmt.Printf("Value of a is 300\n")
   } else {
      // if none of the conditions is true
      fmt.Printf("None of the values is matching\n")
   }

	// post statement is i++
	for i := 0; i < 4; i++{
	fmt.Printf("work\n")  
	}

	// // Infinite loop
	// for {
	// 	fmt.Printf("work\n")  
	// 	}

	// .for loop as  a while loop
	i:= 0
    for i < 3 {
       i += 2
    }
  	fmt.Println(i) 


	// Here rvariable is a array
	rvariable:= []string{"wo", "works", "work"} 
	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j:= range rvariable {
		fmt.Println(i, j) 
	}

	// using maps
	mmap := map[int]string{
		22:"works",
		33:"wo",
		44:"work",
	}
	for key, value:= range mmap {
		fmt.Println(key, value) 
	}


	// using channel
	chnl := make(chan int)
	go func(){
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i:= range chnl {
		fmt.Println(i) 
	}


	for i:=0; i<5; i++{
		fmt.Println(i)
		// For loop breaks when the value of i = 3
		if i == 3{
			  break;
	   	}
	}


	var x int = 0
	// for loop work as a while loop
   Lable1: for x < 8 {
	   if x == 5 {
		  // using goto statement
		  x = x + 1;
		  goto Lable1
	   }
	   fmt.Printf("value is: %d\n", x);
	   x++;     
	} 


	   // for loop work as a while loop
	   for x < 8 {
		if x == 5 {
			 
		   // skip two iterations
		   x = x + 2;
		   continue;
		}
		fmt.Printf("value is: %d\n", x);
		x++;     
	 }
	 
	 // Switch statement with both 
    // optional statement, i.e, day:=4
    // and expression, i.e, day
    switch day:=4; day{
	case 1:
	fmt.Println("Monday")
	case 2:
	fmt.Println("Tuesday")
	case 3:
	fmt.Println("Wednesday")
	case 4:
	fmt.Println("Thursday")
	case 5:
	fmt.Println("Friday")
	case 6:
	fmt.Println("Saturday")
	case 7:
	fmt.Println("Sunday")
	default: 
	fmt.Println("Invalid")
	}


	var value int = 2
      
    // Switch statement without an     
    // optional statement and 
    // expression
   switch {
       case value == 1:
       fmt.Println("Hello")
       case value == 2:
       fmt.Println("Bonjour")
       case value == 3:
       fmt.Println("Namstay")
       default: 
       fmt.Println("Invalid")
   }


   var value1 string = "five"
   // Switch statement without default statement
   // Multiple values in case statement
  switch value1 {
	  case "one":
	  fmt.Println("C#")
	  case "two", "three":
	  fmt.Println("Go")
	  case "four", "five", "six":
	  fmt.Println("Java")
  }


  var value11 interface{}
  switch q:= value11.(type) {
	 case bool:
	 fmt.Println("value is of boolean type")
	 case float64:
	 fmt.Println("value is of float64 type")
	 case int:
	 fmt.Println("value is of int type")
	 default:
	 fmt.Printf("value is of type: %T", q)
	   
 }
}