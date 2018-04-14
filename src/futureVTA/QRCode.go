package main

//QRCode related codes
import qrcode "github.com/skip2/go-qrcode"
import (
	//"encoding/json"
	"net/http"
	"log"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2/bson"
	"time"

	)

var hostname string = "localhost"
var databaseName string = "cmpe281"

type QRCodeStruct struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	UID string `bson:"uid" json:"uid"`
	PARENTID string `bson:"parentid" json:"parentid"`
	QRDATA []byte `bson:"qrdata" json:"qrdata"`
	GENERATEDTIME time.Time `bson:"time" json:"time"`
	USETIMES []*QRCodeUse `bson:"usetimes" json:"usetimes"`
}

type QRCodeUse struct {
	TIME time.Time `bson:"time" json:"time"`
}


func checkIfQRCodeGenerated(uid string,parentid string,result *QRCodeStruct){
	database := Database{hostname, databaseName, nil}
	Connect(&database);
	qrdataC :=  (&database).db.C("qrcode")

	current := time.Now()
	twoHourBack := current.Add(time.Hour * -2)
	
	query := bson.M{"uid": uid, "parentid":parentid,"time":bson.M{"$gte":twoHourBack}}
	//result1 := QRCodeStruct{}
	err := qrdataC.Find(query).One(result)
	if(err!=nil){
		log.Println(err)
	}
}
