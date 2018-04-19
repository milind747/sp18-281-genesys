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

// Signup API
func write(formatter *render.Render) http.HandlerFunc {
        return func(w http.ResponseWriter, req *http.Request) {
                // queue_send(uuid.String())
                session, collection, err := getMongoConnection()
        if err != nil {
                panic(err)
        }
                defer session.Close()
                fmt.Println("Signup :" )
//              uuid := uuid.NewV4()
//              fmt.Println(uuid.String())
                fmt.Println(req.Body)

//              var user User
                var user bson.M
                var result bson.M
                _ = json.NewDecoder(req.Body).Decode(&user)
                fmt.Println("User data:", user )
//              user.Id = bson.NewObjectId()
//              var counter = 100
//              var result bson.M
//              query := bson.M{"_id": bson.ObjectIdHex("5ad657af38a068950fba94a3")}
//              change := bson.M{"id": "nehaj2"}
//               _, err = collection.Upsert(query, change)
//              query := bson.M{"username":"Neha1","password":"123"}
                err = collection.Find(bson.M{"username": user["username"],"password":user["password"]}).One(&result)
                if err!=nil{
                        if err==mgo.ErrNotFound{
                                        err = collection.Insert(user)
                                        if err != nil {
                                            fmt.Println("Cannot Signup")
                                           log.Fatal(err)
                                        }else{
                                                fmt.Println("Signed Up")
                                        }
                        }
                }
				if result != nil {
                                fmt.Println("Cannot Signup. User already exists")
				}
                formatter.JSON(w, http.StatusOK, user)
        }
        }

// LOGIN API
func login(formatter *render.Render) http.HandlerFunc{
        return func(w http.ResponseWriter, req *http.Request){
                session, collection, err := getMongoConnection()
        if err != nil {
                panic(err)
        }
                defer session.Close()
                fmt.Println("Login :" )
                fmt.Println(req.Body)
                var result bson.M

//              var user User
                var user bson.M
                _ = json.NewDecoder(req.Body).Decode(&user)
                fmt.Println("User data in login:", user )
               err = collection.Find(bson.M{"username": user["username"],"password":user["password"]}).One(&result)
//               err = collection.Find(bson.M{"username":"neha").One(&result)
                if err != nil {

                                fmt.Println("Inside login error")
                        if err==mgo.ErrNotFound{
                                fmt.Println("Invalid Username/Password")
}else{
        log.Fatal(err)
}          
       }
                fmt.Println("Response",result)
                formatter.JSON(w, http.StatusOK, result)
        }
        }

// change password API
func changepwd(formatter *render.Render) http.HandlerFunc{
        return func(w http.ResponseWriter, req *http.Request){
                session, collection, err := getMongoConnection()
                if err != nil {
                        panic(err)
                }

                defer session.Close()
                fmt.Println("change password :" )
                fmt.Println(req.Body)
                
        }
