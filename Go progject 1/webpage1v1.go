package main   //need to do this all the time at the start of go progam

import ("fmt"
         "io/ioutil"
          "log"
           "net/http")

           type Page struct {//this will be a page in your go website as a struct you can add title and body and image in the Page
             Title string //title of the Page
             Body []byte // a mutable string that can be changed with a slice vs a strign that cannot be changed we want to the bpdy to change but not the title


           }


            func (p *Page) save() error  { //this is for the save method to save presistent storage
              filename := p.Title + ".txt" //filename will be made into a txt file
              return ioutil.WriteFile(filename, p.Body, 0600) // this iwill be saved in the memory location 0600

            }

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


// going to make a view handler to allow users to view a page
//for now it will only handle urls prefixed with /view/

//creating the fucntion
/*First, this function extracts the page title from r.URL.Path,
the path component of the request URL.
The Path is re-sliced with [len("/view/"):]
to drop the leading "/view/" component of the request path.
This is because the path will invariably begin with "/view/",
which is not part of the page's title.*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)//The function then loads the page data, formats the page with a string of simple HTML, and writes it to w, the http.ResponseWriter.
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
  //To use this handler, we rewrite our main function to initialize http using the viewHandler to handle any requests under the path /view/.
}

func main()  {// this where your code will return
  http.HandleFunc("/view/", viewHandler) //using /view to sue the view handler function
  log.Fatal(http.ListenAndServe(":8000", nil))
  p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")} //page1
  p1.save()
  p2,_:= loadPage("TestPage")
  fmt.Println(string(p2.Body))

}
