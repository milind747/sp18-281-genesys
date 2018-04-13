package main

type User struct {
	UID   string `json:"uid" bson:"uid"`
	ISPARENT bool `json:"isparent" bson:"isparent"`
	PARENTID string `json:"parentid" bson:"parentid"`
	USERNAME string `json:"username" bson:"username"`
}

