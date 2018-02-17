package main

import "fmt"

func main()  {
   //for i:=0; i<10; i++{//most basic for loop i is less than 10 add i
     //fmt.Println(i) //print I}
//option 2 more real world like
//i:=0

//for i<10 {
  //fmt.Println(i)
  //i++}
//option 1
x := 5
for{
  fmt.Println("Do stuff",x) //print this while loop runs
  x+=3 //increment x by 3
  if x > 25{//if x less than 23 then do below or else continue
    break//stop the loop above the if
  }
}

//option 2
for x:=5;x<25;x+=3 {
  fmt.Println("Do stuff",x)
}

//option 3
a:=3 //loop while a is 3
for x:=5; a<25; x+=3{ //do as long as x is 5 and a is less than 25 add 3
  fmt.Println("do stuff",x)//while the above print this
  a+=4//increase a by 4
}


//for index := 0;  < count; ++ {//basic template}

}
