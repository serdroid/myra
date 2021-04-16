package meeting

import (
    "fmt"
//    "io"
    "net/http"
)

func greeterHandler(writer http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(writer, "Hello world")
} 

func RunServer() {
    http.ListenAndServe(":5000", http.HandlerFunc(greeterHandler))
}
