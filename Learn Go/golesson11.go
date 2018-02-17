package main //you need to do this all the time

import ("fmt" // import format and net for http stuff and  I/O utility functions. and xml package for xml parsing
  "net/http"
  "io/ioutil"
  "encoding/xml")


//[number]type is an array this is fixed size for example [5 5]int
//[no number]type is a slice example []int



type SitemapIndex struct{ //creating an object in go its called a struct and this is how to do it
  Locations []Location 'xml:"sitemap"' //creating an arry type Location must captialize these values
}

type Location struct{ //creating a struct this is how to do it struct is an object in Go
  Loc string 'xml:"loc"' //this will look through the xml code and parse out or cut out all info with the loc tag

}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc) //return results in a string Sprintf

}

func main() { //this is where all code runs
  //making a request the request returns a response and error
  //use _ with a variable yuo dont want to use
  //telling code to get the source code from web page below
    resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") //the Get function is part of http must be captial because its part of http this is a get request
    bytes, _ := ioutil.ReadAll(resp.Body)//read everything from the resp.body the above going to be in bytes
    //string_body := string(bytes) // going to convert the above bytes into string
    //fmt.Println(string_body)
    resp.Body.Close()//free the resoucres that mad

    var s SitemapIndex
    xml.Unmarshal(bytes, &s) //unmarhsal bytes in the memory address of s that is done via &

fmt.Println(s.Locations)

}
