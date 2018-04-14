package main

import (
        "fmt"
        "log"
        "net/http"
        "github.com/codegangsta/negroni"
        "github.com/gorilla/mux"
        "github.com/unrolled/render"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "encoding/json" 
        "time"
        "os"
)


// MongoDB Configuration
var mongodb_database = "clipper"
var mongodb_collection = "login"


// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
        formatter := render.New(render.Options{
                IndentJSON: true,
        })
        n := negroni.Classic()
        mx := mux.NewRouter()
        initRoutes(mx, formatter)
        n.UseHandler(mx)
        return n
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
                formatter.JSON(w, http.StatusOK, struct{ Test string }{"API login alive!"})
        }
}
