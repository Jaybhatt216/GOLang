package main //always have to do this

//import "fmt" // fmnt is importing format also need this always
//import "math/rand"//this is how you import math
import ("fmt"
        "math/rand") //right way to do multiple imports, /rand is how you add a function in the import

func main(){
  fmt.Println("A number from 1-100",rand.Intn(100))//how to print stuff, the math.sqrt is using the math import to do square root which is a part of the math import
  //need to capitalize the first letter of a funtion in example Sof sqrt and P of Println
  }  //how to create function in go also where codeing is done
