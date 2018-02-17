package main // yes have to do this all the time
//learing concurrency
import (
  "time"
  "fmt"
)

func say(s string)  {
  for i :=0; i <3; i++ {
    fmt.Println(s)
    time.Sleep(time.Millisecond*100)
  }

}

func main()  {
   go say("Hey") //making go routine by just putting go infront of say
    say("there")
}
