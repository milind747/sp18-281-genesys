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


func generateQRCode(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm();
		if  err != nil {
		        log.Fatal("form")
		        log.Fatal(err)
		}
    		//data, err := json.Marshal(req.Form)
    		//if err != nil {
		//        log.Fatal("marshal")
		//        log.Fatal(err)
    		//}
		//log.Println(data)

		user := new(User)
	       	//if err = json.Unmarshal(data, user); err != nil {
        	//	log.Fatal("unmarshal")
    		//}

		user.UID = req.FormValue("uid")
		user.PARENTID = req.FormValue("parentid")

		database := Database{hostname,databaseName, nil}
		Connect(&database);
		qrdataC :=  (&database).db.C("qrcode")
		current_time := time.Now()
		var png []byte
		previous := QRCodeStruct{}
		checkIfQRCodeGenerated(user.UID,user.PARENTID,&previous)
		if(len(previous.USETIMES) == 0){
			var useArry []*QRCodeUse
			png, err = qrcode.Encode("https://example.org?uid="+user.UID+"&parentid="+user.PARENTID, qrcode.Medium, 256)
			if err != nil {
				log.Fatal(err)
			}
			previous = QRCodeStruct{bson.NewObjectId(), user.UID,user.PARENTID,png,current_time,useArry}
			err = qrdataC.Insert(&previous)
			if err != nil {
				log.Fatal(err)
			}

		}

		formatter.JSON(w, http.StatusOK, previous)

	}
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
