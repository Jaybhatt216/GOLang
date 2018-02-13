package main //you have to do this always it is how the code starts

import "fmt" //importing format

func main(){ //main function is where the code runs
     x := 15
     a := &x /* the & lets you point to something and
             this will be the memory address*/
      fmt.Println(a) // this will be print the memory address hwere x is being stored
      fmt.Println(*a)// the * read through the memory address it will read the value stored in memory address
      *a = 5
      fmt.Println(x) //x will be 5 now because it changed
      *a = *a**a
      fmt.Println(x)
      fmt.Println(*a)

      /* & will prin t the memory address
  * will print the value of the variable stored in the memory addrress*/
         

  }
