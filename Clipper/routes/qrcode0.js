var express = require('express');
var router = express.Router();
var ejs = require("ejs");

var axios = require ('axios');

var server = "http://localhost:3000/"


router.get('/getQrCode', function(req, res) {
  console.log("Inside getQrCode");
        var qrdata = {
            "uid": 1,
            "parentid":1,
	          "qrid":"5aeac3289047b4536baec6b3",
            "isparent":"false",
            "username":"Divyang"
        }
   axios.get(server+'generate?uid=1&parentid=1',qrdata)
        .then(function(response){
          console.log(response);
	      ejs.renderFile('./views/qrcode.ejs',{data:response.data},function(err,result){
    		    if (!err) {
       		     res.status(200).send(result);
    		    }else{
               console.log(err)
               res.status(500).send(result);
            }

	      });
        console.log("qr use updated")

    })
});

router.post('/scanQrCode', function(req, res) {

    var user_details = {
        "uid": "1",
        "parentid":"1",
	       "qrid":"5aeac3289047b4536baec6b3"

        }
    axios.post(server+'scanned?qrid='+req.body.qrid,user_details)
        .then(function(response){
            console.log("qr use updated")
    })
});

router.get('/getHistory', function(req, res) {
    var user_details = {
        "uid": "1",
        "parentid":"1",
	      "qrid":"5aeac3289047b4536baec6b3"
    }
    axios.post(server+'generate',user_details)
        .then(function(response){
            console.log("hostory got as "+response)
            ejs.renderFile('./views/qrcodeHistory.ejs',{data:response},function(err,result){
        		    if (!err) {
           		     res.status(200).send(result);
        		    }
    	      });
    })
});

module.exports = router;


//window.location = "/texiservice?No="+"2&userid="+document.getElementById('userid').value+"&clippercard="+document.getElementById('clippercard').value+"&username="+document.getElementById('username').value+"&password="+document.getElementById('password').value;
