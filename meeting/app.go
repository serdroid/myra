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

// global Application variable
var App Application;

// DependencyWeaver interface provides polymorphism to dependencies.
// Different implementations can provide different dependency trees.
// We use this interface to provide mock implementations for tests.
type DependencyWeaver interface {
    weave() *dependencies
}

// dependency tree structure
type dependencies struct {
    store *dataStore
    meetingResource *MeetingResource   
}

// Production implementation of the DependencyWeaver interface
type AppWire struct {}

func (a AppWire) weave() *dependencies {
    store := NewPostgresDataStore()
    meetingResource := NewMeetingResource(store)
    return &dependencies{store, meetingResource}
}

// Initializes application. Creates router and dependency tree.
// Binds handler functions to router.
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

