package meeting

import (
    "fmt"
    "io"
    "net/http"
    "github.com/gorilla/mux"
)

func greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s\n", name)
} 

func greeterHandler(writer http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    user := "world"
    if val, ok := pathParams["user"]; ok {
        user = val
    }
    greet(writer, user)
} 

func bye(writer io.Writer) {
    fmt.Fprintf(writer, "Good bye")
} 

func byeHandler(writer http.ResponseWriter, r *http.Request) {
    bye(writer)
} 

func RunServer() {
    router := mux.NewRouter()
    router.HandleFunc("/hello/{user}", greeterHandler).Methods(http.MethodGet)

    router.HandleFunc("/bye", byeHandler)
    http.ListenAndServe(":5000", router)
}
