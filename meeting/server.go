package meeting

import (
    "fmt"
    "io"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type App struct {
    router *mux.Router
    meetingResource *MeetingResource
}

var Application App;

func (a *App) Initialize() {
    fmt.Println("Initializing application")
    a.router = mux.NewRouter()
    bindHandlers(a.router)
    a.meetingResource = initializeMeetingResource()
}

func (a *App) Run() {
    http.ListenAndServe(":5000", a.router)
}

func bindHandlers(router *mux.Router) {
    router.HandleFunc("/hello/{user}", greeterHandler).Methods(http.MethodGet)
    router.HandleFunc("/bye", byeHandler).Methods(http.MethodGet)
    router.HandleFunc("/meeting/{user}", meetingHandler).Methods(http.MethodGet)
}

func meetingHandler(writer http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    user, ok := pathParams["user"]
    if ! ok {
        fmt.Fprintf(writer, "Please provide valid user.")
    	return
    }
    meeting := Application.meetingResource.findMeeting(user, "today")
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

