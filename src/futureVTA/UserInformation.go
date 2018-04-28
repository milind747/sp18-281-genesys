package main

import (
	//"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	//"strconv"
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

		err := req.ParseForm()	
		
		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails


		u.UID = req.FormValue("uid")
		u.USERNAME = req.FormValue("username")
		err = c.Find(bson.M{"uid":u.UID}).One(&u)

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
		
		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails
		var ust []UserDetails
		u.UID = req.FormValue("uid")
		err = c.Find(bson.M{"parentid":u.UID}).All(&ust)
		
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, ust)
	}
}


func postUserInfo(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {
		
		req.ParseForm()
		
		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails
		
		u._id = bson.NewObjectId()
		u.UID = req.FormValue("uid")
		u.USERNAME = req.FormValue("username")		
		u.PARENTID = req.FormValue("parentid")
		u.PHONE = req.FormValue("phone")
		u.EMAIL = req.FormValue("email")
		/*err = c.Find(bson.M{"uid":u.UID}).One(&u)
		
		if err != nil {
			log.Fatal("ERRRORR")
		        log.Fatal(err)
		}
	*/
		ins := UserDetails{u._id, u.UID, u.USERNAME, u.PARENTID, u.PHONE, u.EMAIL}
		err := c.Insert(ins)

		//fmt.Println("User Data inserted")	
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("data inserted")
		formatter.JSON(w, http.StatusOK, ins)
			
	}
}

func postUpdateInfo(formatter *render.Render) http.HandlerFunc{
		return func(w http.ResponseWriter, req *http.Request) {
		
		req.ParseForm()
		//var user User
		//_ = json.NewDecoder(req.Body).Decode(&user)

		database := Database{"localhost", "cmpe281", nil}
		data := &database
		Connect(data)
		c := data.db.C("users")

		var u UserDetails

		u._id = bson.NewObjectId()
		u.UID = req.FormValue("uid")
		u.USERNAME = req.FormValue("username")		
		u.PARENTID = req.FormValue("parentid")
		u.PHONE = req.FormValue("phone")
		u.EMAIL = req.FormValue("email")

		/*getid := bson.M{"uid": u.UID}
		change := bson.M{"$set": bson.M{"username": u.USERNAME, "phone": u.PHONE, "email": u.EMAIL}}
		err := c.Update(getid, change)
		*/
		//fmt.Println("User Data inserted")			
		
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("data updated")
		formatter.JSON(w, http.StatusOK, change)

	}
}

