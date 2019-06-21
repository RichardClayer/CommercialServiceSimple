package routes

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name       string
    Method     string
    Pattern    string
    HandlerFun http.HandlerFunc
}

type Routes = []Route

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range PayRoutes {
        router.Methods(route.Method).Name(route.Name).Path(route.Pattern).HandlerFunc(route.HandlerFun)
    }

    return router
}