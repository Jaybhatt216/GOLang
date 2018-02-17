package main //need to do this always

import "fmt" //importing format
//import "time"
import "sync"

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
  defer wg.Done()
  c <- someValue * 5
}

func main()  {
  fooVal := make(chan int, 10)//making the channel and the 10 is a buffer, buffer lets us  we can send these values into the channel without a corresponding concurrent receive.
  for i := 0; i < 10; i++ {//loop for the channe;
      wg.Add(1)//need to do this for each iteration since we added a WaitGroup
    go foo(fooVal, i)//printing the channel
  }
  wg.Wait()//waitig for all wait group to end
 close(fooVal)//close the channe;

  for item := range fooVal {//looping over the channel
    fmt.Println(item)
  }
}
