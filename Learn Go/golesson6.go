package main //need to do this all the time

import ("fmt") //fmt import is the format import

type car struct{ //car is a struct and struct is like objects
// add some values associated with it
 gas_pedal uint16 //uint is basically saying anything from min 0 to max 65535
 brake_pedal uint16
 steering_wheel int16// when unit means only positive, int means negative to positive this is -32k to 32k
 top_speed_kmh float64

  }

func main() {//this is where your code runs always need the main part of running the code here
//you can use var a_car car  = to whatver or just a_car :=
a_car := car{gas_pedal : 22341, //this is best practice as its readable and so on
              brake_pedal: 0,
              steering_wheel: 12561,
              top_speed_kmh: 225.0 }//this is referencing the above struct or object we made earlier
// you can also do a_car := car{22341,0,12561,225.0}
fmt.Println(a_car.gas_pedal)//this how you print stuff also when callin a method like .Println the first letter alawys must be captial

}
