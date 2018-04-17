package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	)


type UserDetails struct {
	_id 		bson.ObjectId 	`bson:"_id" json:"id"`
	UID		string		`bson:"uid" json:"uid"`
	USERNAME	string		`bson:"username" json:"username"`
	PARENTID	string		`bson:"parentid" json:"parentid"`
	PHONE		string		`bson:"phone" json:"phone"`
	EMAIL		string		`bson:"email" json:"email"`
}


func getUserInfo(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {

		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		}
}

