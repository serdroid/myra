package meeting

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func meetingHandler(writer http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    user, ok := pathParams["user"]
    if ! ok {
        fmt.Fprintf(writer, "Please provide valid user.")
    	return
    }
    meeting := App.deps.meetingResource.findMeeting(user, "today")
    writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    encoder.Encode(meeting)
} 

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

