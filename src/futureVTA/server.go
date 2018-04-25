package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	
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
	mx.HandleFunc("/scanned", addQRCodeUseDetail(formatter)).Methods("POST")
	mx.HandleFunc("/newParent",newParent(formatter)).Methods("POST")
	mx.HandleFunc("/getPAmount",readBal(formatter)).Methods("GET")
	mx.HandleFunc("/updatePAmount",updateBal(formatter)).Methods("POST")
	mx.HandleFunc("/addAmount",addBal(formatter)).Methods("POST")
	mx.HandleFunc("/usageHistory",userHistory(formatter)).Methods("GET")
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/read", read(formatter)).Methods("GET")
	mx.HandleFunc("/signup", write(formatter)).Methods("POST")
    mx.HandleFunc("/login", login(formatter)).Methods("POST")
	mx.HandleFunc("/change_pass", changepwd(formatter)).Methods("PUT")
	mx.HandleFunc("/getuserinfo", getUserInfo(formatter)).Methods("GET")
	mx.HandleFunc("/getusers", getAllUsersofParent(formatter)).Methods("GET")
	mx.HandleFunc("/postuserinfo", postUserInfo(formatter)).Methods("POST")
	mx.HandleFunc("/updateuserinfo", postUpdateInfo(formatter)).Methods("POST")
	mx.HandleFunc("/getQrCodeHistory",getQRCodeUseHistory(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


