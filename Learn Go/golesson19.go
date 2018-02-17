package main // yes have to do this all the time
//learing concurrency
import (
  "time"
  "fmt"
  "sync" //sync package
)

var wg sync.waitGroup//adds one for every go routine

func say(s string)  {
  for i :=0; i <3; i++ {
    fmt.Println(s)
    time.Sleep(time.Millisecond*100)
  }
   wg.Done()//the above is done do below
}

func main()  {
  wg.Add(1)//add 1 for go rountine
   go say("Hey") //making go routine by just putting go infront of say
   wg.Add(1)//add 1 for go routine
   go say("there")
   wg.Wait()//wait for go routine


}
