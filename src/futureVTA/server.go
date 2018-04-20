package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"log"
	"net/http"
)


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



// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/generate",generateQRCode(formatter)).Methods("GET")
	mx.HandleFunc("/scanned", postScannedQRCode Detail(formatter)).Methods("POST")
	mx.HandleFunc("/newParent",newParent(formatter)).Methods("POST")
	mx.HandleFunc("/getPAmount",readBal(formatter)).Methods("GET")
	mx.HandleFunc("/updatePAmount",updateBal(formatter)).Methods("POST")
	mx.HandleFunc("/addAmount",addBal(formatter)).Methods("POST")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


