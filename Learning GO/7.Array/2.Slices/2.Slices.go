package main
 
import (
	"fmt"
	"sort"
	"bytes"
)
 
func main() {
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4]
    fmt.Println("Array: ", array)
    fmt.Println("Slice: ", slice)

    slice = append(slice, 4, 5, 6)
    fmt.Println("Slice: ", slice)

	arr := [7]string{"This", "is", "the", "tutorial", "of", "Go", "language"}
    fmt.Println("Array:", arr)
    myslice := arr[1:6]
    fmt.Println("Slice:", myslice)
    fmt.Printf("Length of the slice: %d", len(myslice))
    fmt.Printf("\nCapacity of the slice: %d", cap(myslice))
	fmt.Println()

	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)

	var my_slice_11 = make([]int, 4, 7)
    fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n", my_slice_11, len(my_slice_11), cap(my_slice_11))
    var my_slice_21 = make([]int, 7)
    fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n", my_slice_21, len(my_slice_21), cap(my_slice_21))

	myslice21 := []string{"This", "is", "the", "tutorial", "of", "Go", "language"}
    for e := 0; e < len(myslice21); e++ {
        fmt.Println(myslice21[e])
    }
	for index, ele := range myslice {
        fmt.Printf("Index = %d and element = %s\n", index, ele)
    }

	arr123 := [6]uint{55, 66, 77, 88, 99, 22}
	slc := arr123[0:4]
	fmt.Println("Original_Array: ", arr123)
	fmt.Println("Original_Slice: ", slc)
	slc[0] = 100
	slc[1] = 1
	slc[2] = 10
	fmt.Println("\nNew_Array: ", arr123)
	fmt.Println("New_Slice: ", slc)

	s1 := []int{12, 34, 56}
    var s2 []int
    fmt.Println(s1 == nil)
    fmt.Println(s2 == nil)

	s11 := [][]int{{12, 34},
        {56, 47},
        {29, 40},
        {46, 78},
    }
    fmt.Println("Slice 1 : ", s11)

    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}
    fmt.Println("Before sorting:")
    fmt.Println("Slice 1: ", slc1)
    fmt.Println("Slice 2: ", slc2)
    sort.Strings(slc1)
    sort.Ints(slc2)
    fmt.Println("\nAfter sorting:")
    fmt.Println("Slice 1: ", slc1)
    fmt.Println("Slice 2: ", slc2)
	fmt.Println("Is Slice 1 is sorted?: ", sort.StringsAreSorted(slc1))
	fmt.Println("Is Slice 1 is sorted?: ", sort.IntsAreSorted(slc2))

	banana := []string{"apple", "banana", "cherry"}
    fmt.Println("Slice: ", banana)

	slice_1_1:= [] string { "shi", "shiv", "shiva", "shivam" }
    slice_2_1 := make([] string, 3)
    fmt.Println("Slice_1: ", slice_1_1)
    fmt.Println("Slice_2: ", slice_2_1)
    Copy_1 := copy(slice_2_1, slice_1_1)
	fmt.Println("\nSlice_1: ", slice_1_1)
    fmt.Println("Slice_2: ", slice_2_1)
    fmt.Println("Number of elements copied: ",Copy_1)
	Copy_2:= copy(slice_1_1, [] string { "123geeks", "gfg" })
	fmt.Println("\nSlice_1 : ", slice_1_1)
	fmt.Println( "Number of elements copied:", Copy_2)

	source := []int{1, 2, 3, 4, 5}
    destination := make([]int, len(source))
    copy(destination, source)
    fmt.Println("Source: ", source)
    fmt.Println("Destination: ", destination)

	slcmyfun := []string{"C#", "Python", "C", "Perl"}
    fmt.Println("Initial slice: ", slcmyfun)
    myfun(slcmyfun)
    fmt.Println("Final slice:", slcmyfun)

	slice1 := []byte{1, 2, 3, 4, 5}
    slice2 := []byte{1, 2, 3, 4, 5}
    slice3 := []byte{5, 4, 3, 2, 1}
    fmt.Println(bytes.Equal(slice1, slice2))
    fmt.Println(bytes.Equal(slice1, slice3))

	fmt.Println("\nResult 1:",  bytes.Compare(slice1, slice2))
	fmt.Println("\nResult 2:",  bytes.Compare(slice2, slice3))

	slice_1Trim := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
	fmt.Printf("\nSlice 1: %s", bytes.Trim(slice_1Trim, "!#"))
	fmt.Printf("\nSlice 2: %s", bytes.Trim(slice_1Trim, "e!"))
	fmt.Printf("\nSlice 3: %s", bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*"))
	fmt.Println()
	fmt.Println(string( bytes.TrimSpace([]byte("  Hello, World!  "))))
	fmt.Printf("\nSlice 4: %s",  bytes.Split([]byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}, []byte("eek")))
	fmt.Printf("\nSlice 4: %s", bytes.Split([]byte("GeeksforGeeks, Geek"), []byte("")))
}

func myfun(element []string) {
    element[2] = "Java"
    fmt.Println("Modified slice: ", element)

	element = append(element, "Java")
    fmt.Println("Modified slice append: ", element)
}