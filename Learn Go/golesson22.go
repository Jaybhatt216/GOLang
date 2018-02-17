package main // yes need to type this all the time

import "fmt" //import the format function
//We use Go Channels to connect concurrent goroutines, in order to send and receive values between them, using the channel operator.
func foo(c chan int, someValue int){ //creating the belwo variable and passing the channel created below
  c <- someValue * 5 // c is for the channel operator and <- is used to pass something to it important

}


func main()  { // where your good runs
  fooVal := make(chan int)//making a channel use make then chan for channel then type fooval is a variable we assinged the channel to

go foo(fooVal, 5)
go foo(fooVal, 3)

//option 1
//v1 := <-fooVal
//v2 := <-fooVal
//option2
v1, v2 := <-fooVal, <-fooVal

fmt.Println(v1,v2)

}
