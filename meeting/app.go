package meeting

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type App struct {
    router *mux.Router
    deps *dependencies
}

type wireDependencies interface {
    wire() *dependencies
}

type dependencies struct {
    store *dataStore
    meetingResource *MeetingResource   
}

type AppWire struct {}

func (a AppWire) wire() *dependencies {
    store := NewHardCodedDataStore()
    meetingResource := NewMeetingResource(store)
    return &dependencies{store, meetingResource}
}

var Application App;

func (a *App) Initialize(wir *wireDependencies) {
    fmt.Println("Initializing application")
    a.router = mux.NewRouter()
    bindHandlers(a.router)
    w := *wir
    a.deps = w.wire()
}

func (a *App) Run() {
    http.ListenAndServe(":5000", a.router)
}

func bindHandlers(router *mux.Router) {
    router.HandleFunc("/hello/{user}", greeterHandler).Methods(http.MethodGet)
    router.HandleFunc("/bye", byeHandler).Methods(http.MethodGet)
    router.HandleFunc("/meeting/{user}", meetingHandler).Methods(http.MethodGet)
}

