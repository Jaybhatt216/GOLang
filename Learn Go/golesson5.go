package main //need to dp this everytime

import ("fmt")
import ("net/http") //importing the format function and net/http for web server

func index_handler(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "whoa, Go is neat!")//format what you specify the W
  }//w is a writer and http.ResponseWriter is the type making index_handler a functions

  func about_handler(w http.ResponseWriter, r *http.Request){
       fmt.Fprintf(w, "Expert web Design by Jay thanks to sentdex")//format what you specify the W
    }

func main()  {//the main function is where your code goes type the code there
  http.HandleFunc("/", index_handler) //this will let the http handle functions
  http.HandleFunc("/about/", about_handler) // about page basically
  http.ListenAndServe(":8000", nil)//nil is like none)//this is for the server listen for the request and give first parameter is the port

}
