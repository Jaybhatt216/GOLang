package main // need to do this all the time

import (//importing format and net
        "fmt"
        "net/http"
        "html/template"//html import
  )

type NewsAggPage struct {
  Title string
  News string

}

  func newsAggHandler(w http.ResponseWriter, r*http.Request)  {//using http we make a writer and request listening for out server
  p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}//page on the webpage
  t, _ := template.ParseFiles("basictemplating.html")//making the template
  t.Execute(w,p)
  }

  func indexHandler(w http.ResponseWriter, r*http.Request)  {//using http we make a writer and request listening for out server
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>") //my header on webpage
  }

  func main(){//where the code runs
    http.HandleFunc("/", indexHandler)//read request and write request from above
    http.HandleFunc("/agg/", newsAggHandler)
    http.ListenAndServe(":8000", nil)//local server opening port and nil means none

  }
