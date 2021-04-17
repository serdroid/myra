package meeting

import (
    "fmt"
    "io"
    "net/http"
)

func greet(writer io.Writer) {
    fmt.Fprintf(writer, "Hello world")
} 

func greeterHandler(writer http.ResponseWriter, r *http.Request) {
    greet(writer)
} 

func bye(writer io.Writer) {
    fmt.Fprintf(writer, "Good bye")
} 

func byeHandler(writer http.ResponseWriter, r *http.Request) {
    bye(writer)
} 

func RunServer() {
    http.HandleFunc("/hello", greeterHandler)
    http.HandleFunc("/bye", byeHandler)
    http.ListenAndServe(":5000", nil)
}
