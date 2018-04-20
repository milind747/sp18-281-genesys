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


		var u UserDetails


		u.UID = req.FormValue("uid")
		u.USERNAME = req.FormValue("username")
		err = c.Find(bson.M{"uid":u.UID}).One(&u)

		//err2 = c.Find(bson.M{"parentid":j}).One(&u)	
		//fmt.Println(i)	
	
		if err != nil{
			u._id = bson.NewObjectId()
			err = c.Insert(&u)
		}
		formatter.JSON(w, http.StatusOK, u)	
		
		}
		
}

func getAllUsersofParent(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {

		err := req.ParseForm()	
		//var user User
		//_ = json.NewDecoder(req.Body).Decode(&user)
		
		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails


		u.UID = req.FormValue("uid")
		u.USERNAME = req.FormValue("username")
		err = c.Find(bson.M{"uid":u.UID}).One(&u)

		//err2 = c.Find(bson.M{"parentid":j}).One(&u)	
		//fmt.Println(i)	
	
		if err != nil{
			u._id = bson.NewObjectId()
			err = c.Insert(&u)
		}
		formatter.JSON(w, http.StatusOK, u)	
		
		}
}


func postUserInfo(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {
	
		err := req.ParseForm()	
		//var user User
		//_ = json.NewDecoder(req.Body).Decode(&user)
		
		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, ust)
			
	}
}

func postUpdateInfo(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {

		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails
}
}

