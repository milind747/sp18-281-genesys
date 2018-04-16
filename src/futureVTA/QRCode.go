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

