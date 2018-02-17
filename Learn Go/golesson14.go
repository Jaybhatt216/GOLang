package main

import "fmt"

func main() { //learning maps
  //option 1
  //var grades map[string]float32 //grades map
  grades := make(map[string]float32) //map is refernce type has no values use the make to give it values

  grades["Timmy"] = 42
  grades["Jess"] = 92
  grades["Sam"] = 67

fmt.Println(grades)

TimsGrade := grades["Timmy"]
fmt.Println(TimsGrade)
delete(grades,"Timmy")//delete something from maps
fmt.Println(grades)

for k, v := range grades {
  fmt.Println(k,":",v)
}


}
