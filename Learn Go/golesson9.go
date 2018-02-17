package main //always have to do this

import "fmt"
import "net/http"//import fmt and webserver

func index_handler(w http.ResponseWriter, r *http.Request){ //w means write r means reader, basicall read the web request and write to it
/multiline print
fmt.Fprintf(w'<h1>hey there</h1>
      <p>Go is fast!</p>
      <p>and simple</P>
     ')

fmt.Fprintf(w, "<h1>hey there</h1>")//print on the webpage we can also add in html tags too
 fmt.Fprintf(w, "<p>Go is fast!</p>")
 fmt.Fprintf(w, "<p>and simple</p>")
 fmt.Fprintf(w, "<p>you %s even add %s</p>","can","<strong>variables</strong>")
}

func main(){
  http.HandleFunc("/", index_handler) //the calling the index handler
  http.ListenAndServe(":8000",nil) //local host and nil is none

}
