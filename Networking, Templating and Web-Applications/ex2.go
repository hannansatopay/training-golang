package main
import (
"fmt"
"net/http"
"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Inside HelloServer handler")
  fmt.Fprint(w, "Hello, " + req.URL.Path[1:])
}

func Spy(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, "James Bond")
}


func main() {
  http.HandleFunc("/",HelloServer)
  http.HandleFunc("/spy",Spy)
  err := http.ListenAndServe("0.0.0.0:3000", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err.Error())
  }
}