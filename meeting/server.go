package meeting

import (
    "fmt"
//    "io"
    "net/http"
)

func greeterHandler(writer http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(writer, "Hello world")
} 

func byeHandler(writer http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(writer, "Good bye")
} 

func RunServer() {
    http.HandleFunc("/", greeterHandler)
    http.HandleFunc("/bye", byeHandler)
    http.ListenAndServe(":5000", nil)
}
