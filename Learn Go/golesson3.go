package main //you will need to this always

import ("fmt") //import the format function

func add(x float64,y float64) float64{//aslo needs to specify that the function will return something//when you choose variables you need to delcare types like float int etc  { //this is the add function which adds 2 variables
         //another way of typing the above is x,y float64
   return x+y
}

func multiple(a,b string) (string,string) { // two different things even if they are same types
  return a,b
}

func main(){//all go languange runs in the main function so this needs doing
  //once the function add is made in the main you can just add
  // different ways to create variables
  //var name type as you can see its var then name of it and the type like float or number
  var num1 float64 = 5.6
  var num2 float64 = 9.5
//another way to do type the above is var num1,num2 float64 = 5.6,9.5
//another way is var num1.num2 := 5.6,9.5 the colon : tells go that look above its a float64 only works while in the same function
  w1, w2 := "Hey","There"
  fmt.Println(multiple(w1,w2))

  fmt.Println(add(num1,num2)) //printing the adding
}

/* convert one type to another below example
var a int = 62
var b float64 = float64(a)
so now a is a float64 not an int
x := a x will be type int
 */
