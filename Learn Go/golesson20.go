/*package main

import "fmt"

func foo() {
  defer fmt.Println("done!") //defer statement puts off the execution of something until the surrounding dunction is done
  defer fmt.Println("are we done")
  for i :=0; i < 5;i++ {
    defer fmt.Println(i)
  }
  fmt.Println("doing stuff who knows what")
}

func main(){
  foo()
}*/

package main // yes have to do this all the time
//learing concurrency
import (
  "time"
  "fmt"
  "sync" //sync package
)

var wg sync.WaitGroup//adds one for every go routine

func say(s string)  {
  defer wg.Done()
  for i :=0; i <3; i++ {
    fmt.Println(s)
    time.Sleep(time.Millisecond*100)
  }
   //wg.Done()//the above is done do below
}

func main()  {
  wg.Add(1)//add 1 for go rountine
   go say("Hey") //making go routine by just putting go infront of say
   wg.Add(1)//add 1 for go routine
   go say("there")
   wg.Wait()//wait for go routine


}
