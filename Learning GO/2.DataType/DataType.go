package main

import (
	"fmt"
)

func main (){
	    // Using 8-bit unsigned int
		var X uint8 = 225
		fmt.Println(X, X-3)
		// Using 16-bit signed int
		var Y int16 = 32767
		fmt.Println(Y+2, Y-2)

		var x int16 = 170
		var y int16 = 83
		//Addition
		fmt.Printf(" addition :  %d + %d = %d\n ", x, y, x+y)
		//Subtraction
		fmt.Printf("subtraction : %d - %d = %d\n", x, y, x-y)
		//Multiplication
		fmt.Printf(" multiplication : %d * %d = %d\n", x, y, x*y)
		//Division
		fmt.Printf(" division : %d / %d = %d\n", x, y, x/y)
		//Modulus
		fmt.Printf(" remainder : %d %% %d = %d\n", x, y, x%y)

		a := 20.45
		b := 34.89
		// Subtraction of two
		// floating-point number
		c := b-a
		// Display the result
		fmt.Printf("Result is: %f", c)
		// Display the type of c variable
		fmt.Printf("\nThe type of c is : %T", c) 


		var aa complex128 = complex(6, 2)
		var bb complex64 = complex(9, 2)
		fmt.Println(aa)
		fmt.Println(bb)
		// Display the type
	   fmt.Printf("The type of a is %T and "+
				 "the type of b is %T", aa, bb)


		comp1 := complex(10, 11)
		// complex number init syntax
		comp2 := 13 + 33i
		fmt.Println("Complex number 1 is :", comp1)
		fmt.Println("Complex number 1 is :", comp2)
		// get real part
		realNum := real(comp1)
		fmt.Println("Real part of complex number 1:", realNum)
		// get imaginary part
		imaginary := imag(comp2)
		fmt.Println("Imaginary part of complex number 2:", imaginary)


		str1 := "GeeksforGeeks"
		str2:= "geeksForgeeks"
		str3:= "GeeksforGeeks"
		result1:= str1 == str2
		result2:= str1 == str3
		// Display the result
		fmt.Println( result1)
		fmt.Println( result2)
		// Display the type of
		// result1 and result2
		fmt.Printf("The type of result1 is %T and "+ "the type of result2 is %T", result1, result2)


		// str variable which stores strings
		str := "GeeksforGeeks"
		// Display the length of the string
		fmt.Printf("Length of the string is:%d", len(str))
		// Display the string
		fmt.Printf("\nString is: %s", str)
		// Display the type of str variable
		fmt.Printf("\nType of str is: %T", str)
		var str11 string = "STRING_"
		var str22 string = "Concatenation"
		// Concatenating strings using + operator
		fmt.Println("New string : ", str11+str22)

		p:= 34
		q:= 20
		  
		// & (bitwise AND)
		result11:= p & q
		fmt.Printf("Result of p & q = %d", result11)
		  
		// | (bitwise OR)
		result21:= p | q
		fmt.Printf("\nResult of p | q = %d", result21)
		// ^ (bitwise XOR)
		result31:= p ^ q
		fmt.Printf("\nResult of p ^ q = %d", result31)
		// << (left shift)
		result41:= p << 1
		fmt.Printf("\nResult of p << 1 = %d", result41)
		// >> (right shift)
		result51:= p >> 1
		fmt.Printf("\nResult of p >> 1 = %d", result51)
		// &^ (AND NOT)
		result61:= p &^ q
		fmt.Printf("\nResult of p &^ q = %d", result61)


		a1 := 4
		// Using address of operator(&) and
		// pointer indirection(*) operator
		b1 := &a1
		fmt.Println(*b1)
		*b1 = 7
		fmt.Println(a1)
}