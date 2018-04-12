package main

import(
        "gopkg.in/mgo.v2"
        "fmt"
        "time"
)
// MongoDB Config
var mongodb_server = ""
var Mongodb_database = ""
var Mongodb_collection = ""
var MongoSession *mgo.Session

func init(){

        mongodb_server = "mongodb://10.0.1.151:27017,10.0.1.50:27017,10.0.2.12:27017,10.0.2.83:27017,10.0.2.148:27017/?replicaSet=njcloud"
        Mongodb_database="clipper"
        Mongodb_collection="login"
}


