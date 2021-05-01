package meeting

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func getMeetingHandler(writer http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    user, ok := pathParams["user"]
    if ! ok {
        fmt.Fprintf(writer, "Please provide valid user.")
    	return
    }
    ds := *App.deps.store
    meeting := ds.findMeeting(user, "today")
    writer.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(writer)
    encoder.Encode(meeting)
} 


func createMeetingHandler(writer http.ResponseWriter, req *http.Request) {
    var meet Meeting
    decoder := json.NewDecoder(req.Body)
    defer req.Body.Close()
    err := decoder.Decode(&meet)
    if err != nil {
        fmt.Fprintf(writer, "Please provide valid meeting.")
    	return
    }
    ds := *App.deps.store
    err = ds.createMeeting(&meet)
    if err != nil {
        fmt.Fprintf(writer, "Error while creating meeting : %s\n", err)
    	return
    }
    writer.Header().Set("Content-Type", "application/json")
    res := map[string]string {"id": meet.ID}
    encoder := json.NewEncoder(writer)
    encoder.Encode(res)
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

