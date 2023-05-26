package main
  
import "fmt"
  
func myfunc(ch chan int) {
    fmt.Println(234 + <-ch)
}
func main() {
    var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel)
    fmt.Printf("Type of the channel: %T ", mychannel)
    mychannel1 := make(chan int)
    fmt.Println("\nValue of the channel1: ", mychannel1)
    fmt.Printf("Type of the channel1: %T ", mychannel1)

	fmt.Println("start Main method")
    ch := make(chan int)
    go myfunc(ch)
    ch <- 23
    fmt.Println("End Main method")

	// Unidirectional Channel in Golang
	mychanl1 := make(<-chan string)
    mychanl2 := make(chan<- string)
    fmt.Printf("%T", mychanl1)
    fmt.Printf("\n%T", mychanl2)

    mychanl := make(chan string)
    go sending(mychanl)
    fmt.Println(<-mychanl)
}
func sending(s chan<- string) {
    s <- "GeeksforGeeks"
}