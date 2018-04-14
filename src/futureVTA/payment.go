//adding payment api code in GO for payment microservice in this file.	

package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/unrolled/render"
    "gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
	
)

func newParent(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {

	//code for database connection with mongo in local
	database := Database{"localhost", "cmpe281", nil}
	data := &database
	Connect(data)

	}
}
