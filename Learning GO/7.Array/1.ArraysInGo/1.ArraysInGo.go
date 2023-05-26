package main
  
import "fmt"
  
func main() {
	arr:= [4]string{"shi", "shiv", "shiva", "shivam"}
	fmt.Println("Elements of the array:")
	for i:= 0; i < len(arr); i++{
	fmt.Println(arr[i])
	}

	arr2 := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
        {"C++", "Go", "HTML"}}
    fmt.Println("Elements of Array 2")
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
  
            fmt.Println(arr2[x][y])
        }
	}

	var arr1 [2][2]int
    arr1[0][0] = 100
    arr1[0][1] = 200
    arr1[1][0] = 300
    arr1[1][1] = 400
    fmt.Println("Elements of array 2")
    for p := 0; p < 2; p++ {
        for q := 0; q < 2; q++ {
            fmt.Println(arr1[p][q])
        }
    }

	var myarr2[2] int 
	fmt.Println("Elements of the Array: ", myarr2)

	myarray3:= [...]string{"GFG", "gfg", "geeks", "GeeksforGeeks", "GEEK"}
	fmt.Println("Elements of the array: ", myarray3)
	fmt.Println("Length of the array is:", len(myarray3))

	// an array is of value type not of reference type
	my_array:= [...]int{100, 200, 300, 400, 500}
	fmt.Println("Original array(Before):", my_array)
	new_array := my_array 
	fmt.Println("New array(before):", new_array)
	new_array[0] = 500
	fmt.Println("New array(After):", new_array)
	fmt.Println("Original array(After):", my_array)

	arr11:= [3]int{9,7,6}
	arr21:= [...]int{9,7,6}
	arr31:= [3]int{9,5,3}
	fmt.Println(arr11==arr21)
	fmt.Println(arr21==arr31)
	fmt.Println(arr11==arr31)

	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}
    my_arr2 := my_arr1
    fmt.Println("Array_1: ", my_arr1)
    fmt.Println("Array_2:", my_arr2)
    my_arr1[0] = "C++"
    fmt.Println("\nArray_1: ", my_arr1)
    fmt.Println("Array_2: ", my_arr2)

	my_arr11:= [6] int { 12, 456, 67, 65, 34, 34 }
    my_arr21 := &my_arr11
	fmt.Println("Array_1: ", my_arr11)
    fmt.Println("Array_2:", *my_arr21)
	my_arr11[5]= 1000000
    fmt.Println("\nArray_1: ", my_arr11)
    fmt.Println("Array_2:", *my_arr21)

	originalArray := []int{1, 2, 3, 4, 5}
    copyArray := make([]int, len(originalArray))
    // for i, value := range originalArray {
    //     copyArray[i] = value
    // }
	// copy(copyArray, originalArray)
	copyArray = append(copyArray, originalArray...)
    fmt.Println("Original Array: ", originalArray)
    fmt.Println("Copy Array: ", copyArray)

	var send = [6]int{67, 59, 29, 35, 4, 34}
	var res int
    res = myfun(send, 6)
    fmt.Printf("Final result is: %d ", res)

	array := []int{1, 2, 3, 4, 5}
    printArray(array)
}

func myfun(a [6]int, size int) int {
    var k, val, r int
    for k = 0; k < size; k++ {
        val += a[k]
    }
    r = val / size
    return r
}
func printArray(array []int) {
    for _, value := range array {
        fmt.Println(value)
    }
}
