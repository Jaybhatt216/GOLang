package main //you need to do this all the time

import ("fmt" // import format and net for http stuff and  I/O utility functions.
  "net/http"
  "io/ioutil")


func main() { //this is where all code runs
  //making a request the request returns a response and error
  //use _ with a variable yuo dont want to use
  //telling code to get the source code from web page below
    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //the Get function is part of http must be captial because its part of http this is a get request
    bytes, _ := ioutil.ReadAll(resp.Body)//read everything from the resp.body the above going to be in bytes
    string_body := string(bytes) // going to convert the above bytes into string
    fmt.Println(string_body)
    resp.Body.Close()//free the resoucres that mad


}
