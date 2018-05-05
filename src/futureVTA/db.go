package main

import(
	mgo "gopkg.in/mgo.v2"
	"log"
	
)

// struct to define the Database fields
type Database struct {
	host   string
	Database string
	db *mgo.Database
}


func Connect(m *Database) {
	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatal(err)
	}
	
	m.db =  session.DB(m.Database)
}
