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

