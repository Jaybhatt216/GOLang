package main // yes have to do this all the time
//learing concurrency
import (
  "time"
  "fmt"
  "sync" //sync package
)

var wg sync.WaitGroup//adds one for every go routine

func cleanup(){
  if r := recover(); r!= nil {
    fmt.Println("recovered in cleanup", r)
    }   //learing recover statement
}


func say(s string)  {
  defer wg.Done() // defer statement delyas exeuation of the statement next to it
  defer cleanup()
  for i :=0; i <3; i++ {
  time.Sleep(time.Millisecond*100)
    fmt.Println(s)
    ifi == 2 {//test for panic statement
      panic("0h dear, a 2") //learing the panic statement
    }

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
