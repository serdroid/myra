package meeting

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type Application struct {
    router *mux.Router
    deps *dependencies
}

var App Application;

type DependencyWeaver interface {
    weave() *dependencies
}

type dependencies struct {
    store *dataStore
    meetingResource *MeetingResource   
}

type AppWire struct {}

func (a AppWire) weave() *dependencies {
    store := NewHardCodedDataStore()
    meetingResource := NewMeetingResource(store)
    return &dependencies{store, meetingResource}
}

func (a *Application) Initialize() {
    var wr DependencyWeaver = &AppWire{}
    App.initDeps(&wr)
}


func (a *Application) initDeps(wr *DependencyWeaver) {
    fmt.Println("Initializing application")
    a.router = mux.NewRouter()
    bindHandlers(a.router)
    w := *wr
    a.deps = w.weave()
}

func (a *Application) Run() {
    http.ListenAndServe(":5000", a.router)
}

func bindHandlers(router *mux.Router) {
    router.HandleFunc("/hello/{user}", greeterHandler).Methods(http.MethodGet)
    router.HandleFunc("/bye", byeHandler).Methods(http.MethodGet)
    router.HandleFunc("/meeting/{user}", meetingHandler).Methods(http.MethodGet)
}

