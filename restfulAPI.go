package main //do this always

import "github.com/gorilla/mux" //mux router import
import "net/http"
import "log"
import "encoding/json"




type Person struct {
  ID string 'json:"id,omitempty"'
  Firstname string 'json:"firstname,omitempty"'
  Lastname string 'json:"lastname"omitempty'
  Address *Address 'json:"address",omitempty'

}

type Address struct {
  City string 'json:"city",omitempty'
  State string 'json:"state",omitempty'

}

var people []Person



func GetPersonEndpoint(w httpResponseWriter, req *http.Request)  {

}
func GetPeopleEndpoint(w httpResponseWriter, req *http.Request)  {
  json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w httpResponseWriter, req *http.Request)  {

}

func DeletePersonEndpoint(w httpResponseWriter, req *http.Request)  {

}


func main()  {
  router := mux.NewRouter()
  people = append(people, Person{ID: "1", Firstname: "Jay", Lastname:"Bhatt", Address: &Address{City:"New York",State:"NY"}})
  people = append(people, Person{ID: "2", Firstname: "John", Lastname:"Jay", Address: &Address{City:"New York",State:"NY"}})
  people = append(people, Person{ID: "3", Firstname: "Jake", Lastname:"Jhon", Address: &Address{City:"New York",State:"NY"}})
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpointPeopleEndpoint).Methods("POST")
  router.HandleFunc("/people{id}", DeletePersonEndpoint).Methods("DELETE")
  log.Fatal(http.ListenandServe(":8000",router))


}
