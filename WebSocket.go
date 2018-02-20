package main //need to do thsi always

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
  "time"//the time package
)

var upgrader = websocket.upgrader{}//going to upgrade http Request to a websocket

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)) {
    http.serverFile(w, r, "index.html")
  }

  http.HandleFunc("/v1/w2", func(w http.ResponseWriter, r *http.Request)){
    var Conn, _ := upgrader.upgrade(w ,r, nil) //upgrade the Request to websocket
    go func (Conn *websocket.Conn) {
      for {
        mType, msg, _ := Conn.ReadMessage() // read the message coming from the client
       Conn.WriteMessage(mtype, msg)//write to the ReadMessage

      }
    }
  }(Conn)
  http.HandleFunc("/v2/ws" , func (w http.ResponseWriter, r*http.Request)  {
    var Conn, _ := upgrader.upgrade(w ,r, nil)
    go func (Conn *websocket.Conn)  {
      for {
        _, msg, _ := Conn.ReadMessage()
        println(string(msg)) //client asking for update or whatever
      }
      http.HandleFunc("/v3/w2", func(w http.ResponseWriter, r *http.Request)){
        var Conn, _ := upgrader.upgrade(w ,r, nil) //upgrade the Request to websocket
        go func (Conn *websocket.Conn) {
          ch := time.Tick(5 *time.Second //making a channel to send the info to the server every 5 sec
            for range ch {
              Conn.WriteJSON(myStruct){
                username: "J!N!"
                Firstname: "Jay"
                Lastname: "Jay"


              }

            }

        }
      }(Conn)
    }
  })
  http.ListenAndServe(":3000",nil)
}


type myStruct struct {
  username string 'json:"username"'
  Firstname string 'json:"firstname"'
  lastname string 'json:"lastname"'

}

http.HandleFunc("/v4/w2", func(w http.ResponseWriter, r *http.Request)){
  var Conn, _ := upgrader.upgrade(w ,r, nil) //upgrade the Request to websocket
  go func(Conn *websocket.Conn){
    for {
      mType, msg, err := Conn.ReadMessage //with gorilla if client sends a close then its an error
     if err != nil {
       Conn.Close()
     }
    }
  }
  go func (Conn *websocket.Conn) {
    ch := time.Tick(5 *time.Second //making a channel to send the info to the server every 5 sec
      for range ch {
        Conn.WriteJSON(myStruct){
          username: "J!N!"
          Firstname: "Jay"
          Lastname: "Jay"


        }

      }

  }
}(Conn)
