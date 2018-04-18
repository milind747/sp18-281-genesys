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

type Payment struct {
	_id bson.ObjectId `bson:"_id" json:"_id"`
	id int `bson:"id" json:"id"`
	amount int  `bson:"amount" json:"amount"`
}

func newParent(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {

	//parsing json file	
	req.ParseForm()
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)
	//code for database connection with mongo in local
	database := Database{"localhost", "cmpe281", nil}
	data := &database
	Connect(data)
	//collection
	c := data.db.C("payment")
	
	//need parent id for my api which i took from query string
	var p Payment 
	i, err := strconv.Atoi(req.FormValue("parentid"))
	err = c.Find(bson.M{"id":i}).One(&p) 
		 
	if err==nil{
		//handle error
		fmt.Println("Already present in db")
		formatter.JSON(w, http.StatusOK, false)	
			

	} else {
	
		//handle error
		p := c.Insert(bson.M{"id":i,"amount":0})
		d.Insert(bson.M{"id":i,"tamt":0,"bal":0,"timestamp":time.Now()})
		fmt.Println("output:", p)
		formatter.JSON(w, http.StatusOK, true)

	}
	//below code does not give output in struct as needed	
	//for unmarshalling data into JSON format and storing in struct	
	//data1, err := json.Marshal(req.Form)
    	//if err != nil {
         //Handle error
    	//}
    	//var user User
    	//if err = json.Unmarshal(data1, user); err != nil {
        // Handle error
    	//}
   	//fmt.Printf("%+v", user)

	}
}

//function to read Balance of the parent user
func readBal(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	database := Database{"localhost", "cmpe281", nil}
	data := &database
	Connect(data)
	c := data.db.C("payment")
	fmt.Println("collection:", c)
        var result Payment
	i, err := strconv.Atoi(req.FormValue("parentid"))
	//fmt.Println("USER.PARENTID:", i)
	//var re = Payment{}
        err = c.Find(bson.M{"id":i}).One(&result)
        if err != nil {
		log.Fatal("ERRRORR")
                log.Fatal(err)
		//x := c.Insert(bson.M{"id":i,"amount":0})
	//if x != nil {
		//panic(err)
	}
	fmt.Println("Amount:", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}
