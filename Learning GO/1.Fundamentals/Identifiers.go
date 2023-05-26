package main

import (
	"strconv"
	"fmt"
)

var global uint32 =98
var global_f float32 =98.6
var (
	Name string = "Shyam"
	Enroll uint64 = 21117125
	Branch string = "Mech. eng."
	Batch string = "Q5"

)

func main() {
	var num uint8 = 9 //variablee of interger uint8
	fmt.Println(num)
	i:=19  //variable of unknown data type 
	fmt.Println(i)
	fmt.Printf("%v, %T\n", global, global) 
	fmt.Printf("%v, %T\n", global_f, global_f) 
	fmt.Printf("%v, %T\n", Enroll, Enroll) 
	// 
	// 
	var ii int =42
	fmt.Printf("%v, %T\n",ii,ii)
	var jj float32
	jj=float32(ii)
	fmt.Printf("%v, %T\n",jj,jj)
	// 
	var jjj float32 = 65.6
	fmt.Printf("%v, %T\n",jjj,jjj)
	var iii int =int(jjj)
	fmt.Printf("%v, %T\n",iii,iii)
	// 
	var ss string =strconv.Itoa(ii)
	fmt.Printf("%v, %T\n",ss,ss)
}

// Go is a procedural programming language