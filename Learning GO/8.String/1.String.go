package main
 
import (
    "fmt"
    "unicode/utf8"
	"strings"
	"bytes"
)
 
func main() {
    My_value_1 := "Welcome to Shivam"
    var My_value_2 string
    My_value_2 = "Shivam"
    fmt.Println("String 1: ", My_value_1)
    fmt.Println("String 2: ", My_value_2)

	for index, s := range My_value_2 {
        fmt.Printf("The index number of %c is %d\n", s, index)
    }

	for c := 0; c < len(My_value_2); c++ {
		fmt.Printf("\nCharacter = %c Bytes = %v", My_value_2[c], My_value_2[c])
	}

    myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
    mystring1 := string(myslice1)
	fmt.Println()
    fmt.Println("String 1: ", mystring1)
    myslice2 := []rune{0x0047, 0x0065, 0x0065,0x006b, 0x0073}
    mystring2 := string(myslice2)
    fmt.Println("String 2: ", mystring2)

	fmt.Println("Length 2:",  utf8.RuneCountInString(mystring2))
	fmt.Println("Compare: ", "aa"=="aa")
	fmt.Println("Compare: ", "aa"!="aa")
	fmt.Println("Compare: ", "aa" > "aa")
	fmt.Println(strings.Compare("Roorkee","Roorkee"))
	fmt.Println(strings.Compare("Gee","Roorkee"))
	fmt.Println("Gee" + "Roorkee")

	var b bytes.Buffer
    b.WriteString("s")
    b.WriteString("h")
    b.WriteString("i")
    b.WriteString("v")
    b.WriteString("a")
    b.WriteString("m")
    fmt.Println("String: ", b.String())

	result :=fmt.Sprintf("%s%s%s%s", "str1","str2", "str3", "str4")
	fmt.Println(result)

	myslice21 := []string{"Welcome", "To","Roorkee", "Portal"}
	result11 := strings.Join(myslice21, "-")
	fmt.Println(result11)

	var str strings.Builder
    str.WriteString("Welcome")
    fmt.Println("String: ", str.String())
    str.WriteString(" to")
    str.WriteString(" Roorkee!")
    fmt.Println("String: ", str.String())

	fmt.Println("Answer: ", strings.Trim("!!Welcome to Roorkee !!", "!"))
	fmt.Println("Answer: ", strings.TrimLeft("!!Welcome to Roorkee !!", "!"))
	fmt.Println("Answer: ", strings.TrimRight("!!Welcome to Roorkee !!", "!"))
	fmt.Println("Answer: ", strings.TrimSpace("              !!Welcome to Roorkee !!"))
	fmt.Println("Answer: ", strings.TrimSuffix("!!Welcome to Roorkee !!", "!!"))
	fmt.Println("Answer: ", strings.TrimPrefix("!!Welcome to Roorkee !!", "!!Welcome"))
	fmt.Println("Answer: ", strings.Split("!!Welcome to Roorkee !!", ""))
	fmt.Println("Answer: ", strings.SplitAfter("!!Welcome to Roorkee !!", "!"))
	fmt.Println("Answer: ", strings.SplitAfterN("!!Welcome to Roorkee!!", "!",3))
	fmt.Println("Answer: ", strings.Contains("!!Welcome to Roorkee!!", "!"))
	fmt.Println("Answer: ", strings.ContainsAny("!!Welcome to Roorkee!!", "es"))
	fmt.Println("Answer: ", strings.Repeat("!!Welcome to Roorkee!!", 2))
	fmt.Println("Answer: ", strings.Index("Hello, !!Welcome to Roorkee!!", "!"))
	fmt.Println("Answer: ", strings.IndexAny("Hello, !!Welcome to Roorkee!!", "e!"))
	fmt.Println("Answer: ", strings.IndexByte("Hello, !!Welcome to Roorkee!!", 'k'))
	fmt.Println("Answer: ", strings.Count("Hello, !!Welcome to Roorkee!!", "o"))
}