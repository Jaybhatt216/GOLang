package main   //need to do this all the time at the start of go progam

import ("fmt"
         "io/ioutil"
          "log"
           "net/http") // importing packages

type Page struct {//this will be a page in your go website as a struct you can add title and body and image in the Page
  Title string //title of the Page
  Body []byte // a mutable string that can be changed with a slice vs a strign that cannot be changed we want to the bpdy to change but not the title


}

/*This method's signature reads: "This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error."

This method will save the Page's Body to a text file. For simplicity, we will use the Title as the file name.

The save method returns an error value because that is the return type of WriteFile (a standard library function that writes a byte slice to a file). The save method returns the error value, to let the application handle it should anything go wrong while writing the file. If all goes well, Page.save() will return nil (the zero-value for pointers, interfaces, and some other types).

The octal integer literal 0600, passed as the third parameter to WriteFile, indicates that the file should be created with read-write permissions for the current user only. (See the Unix man page open(2) for details.)

In addition to saving pages, we will want to load pages, too:
*/

 func (p *Page) save() error  { //this is for the save method to save presistent storage
   filename := p.Title + ".txt" //filename will be made into a txt file
   return ioutil.WriteFile(filename, p.Body, 0600) // this iwill be saved in the memory location 0600

 }

// the above func saves pages while the below func will load pages
/* The function loadPage constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values.

Functions can return multiple values.
The standard library function io.ReadFile returns
[]byte and error. In loadPage, error isn't being handled yet;
the "blank identifier" represented by the underscore (_)
symbol is used to throw away the error return value
(in essence, assigning the value to nothing).*/


/*func loadPage(title string) *Page {
  filename := title + ".txt"
  body, _ := ioutil.ReadFile(filename)
  return &Page{Title: title, Body: body}
}*/

// below is what what would happen if ReadFile above gets an error good practice to assume and plan for anything failing

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil{ // if error does not equal none
    return nil, err//rerturn the error
    }
    return &Page{Title: title, Body:body}, nil
}

func handler(w http.ResponseWriter, r*http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])



 func main()  {// this where your code will return
   http.HandleFunc("/", handler) //using /view to sue the view handler function
   log.Fatal(http.ListenAndServe(":8000", nil))
   p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")} //page1
   p1.save()
   p2,_:= loadPage("TestPage")
   fmt.Println(string(p2.Body))

 }
