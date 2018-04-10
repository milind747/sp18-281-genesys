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


