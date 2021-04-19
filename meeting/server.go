package meeting

import (
    "fmt"
    "io"
    "net/http"
    "github.com/gorilla/mux"
)

type App struct {
    router *mux.Router
    meetingResource *MeetingResource
}

var Application App;

func (a *App) Initialize() {
    a.router = mux.NewRouter()
    a.meetingResource = initializeMeetingResource()
}

func (a *App) Run() {
    bindHandlers(a.router)
    http.ListenAndServe(":5000", a.router)
}

func bindHandlers(router *mux.Router) {
    router.HandleFunc("/hello/{user}", greeterHandler).Methods(http.MethodGet)
    router.HandleFunc("/bye", byeHandler)
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

