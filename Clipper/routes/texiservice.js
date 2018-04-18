var express = require('express');
var router = express.Router();
var ejs = require("ejs");

var axios = require ('axios');


router.get('/', function(req, res) {
  console.log("Inside the texiservice");

  console.log("req.query.No:"+req.query.No)
  
  if(req.query.No == "1" || req.query.No == "0"){
  ejs.renderFile('./views/contact.ejs',{data:req.query.No},function(err,result){
    if (!err) {
     
        res.status(200).send(result);
    }
  
});
}

else if(req.query.No == "2"){

    // API Call for the get user info 
    console.log("req.query.userid:"+req.query.userid)
    console.log("req.query.clippercard:"+req.query.clippercard)
    console.log("req.query.username:"+req.query.username)
    console.log("req.query.password:"+req.query.password)
    
    if (req.query.direct != "T"){
    var user_details = {
        "Id": 11,
        "Accounts":{
                "Type": "Uber",
                "Username": req.query.username,
                "Clippercarnumber": parseInt(req.query.clippercard),
                "Password": req.query.password,
                "Payment": 0,
                "Noofrides":0
        }
    }

    axios.post('http://localhost:3001/updateUserData',user_details)
        .then(function(response){
            console.log("User Data Updated")

            var payload = {
                "Id": 11,
                "Accountstype":"Uber",
                "Clippercarnumber":parseInt(req.query.clippercard)
            }
            axios.post('http://localhost:3001/getUserData',payload)
                .then(function(response){
                    console.log("response.data"+response.data)
                    ejs.renderFile('./views/contact.ejs',{data:req.query.No,details:response.data},function(err,result){
                    if (!err) {
                            res.status(200).send(result);
                            }
            
          });
        
        })
            })

        }

        else{
            var payload = {
                "Id": 11,
                "Accountstype":"Uber",
                "Clippercarnumber":parseInt(178)
            }
            axios.post('http://localhost:3001/getUserData',payload)
                .then(function(response){
                    console.log("response.data"+response.data)
                    ejs.renderFile('./views/contact.ejs',{data:req.query.No,details:response.data},function(err,result){
                    if (!err) {
                            res.status(200).send(result);
                            }
            
          });
        
        })

        }

    
  }

    else if(req.query.No == "3" && req.query.data_int=="3_D"){
        console.log("Inside 3_D")
    var payload = {
        "Id": 11,
        "Accountstype":"Uber",
        "Payment":parseInt(101)
    }

    axios.post('http://localhost:3001/updateRides',payload)
                .then(function(response){
                    console.log("response.data"+response.data)
                    ejs.renderFile('./views/contact.ejs',{data:req.query.No,data_int:req.query.data_int},function(err,result){
                    if (!err) {
                            res.status(200).send(result);
                            }
            
          });
        
        })
  }

  else if(req.query.No == "3" ){
    ejs.renderFile('./views/contact.ejs',{data:req.query.No},function(err,result){
      if (!err) {
       
          res.status(200).send(result);
      }
    
  });
  }
});

module.exports = router;


//window.location = "/texiservice?No="+"2&userid="+document.getElementById('userid').value+"&clippercard="+document.getElementById('clippercard').value+"&username="+document.getElementById('username').value+"&password="+document.getElementById('password').value;