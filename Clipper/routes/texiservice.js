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

    var payload = {
        "Id": 5,
        "Accountstype":"Uber",
        "Clippercarnumber":0
    }
    axios.post('http://localhost:3001/getUserData',payload)
        .then(function(response){
            console.log("response.data"+response.data)
            ejs.renderFile('./views/contact.ejs',{data:req.query.No},function(err,result){
            if (!err) {
                    res.status(200).send(result);
                    }
    
  });

})
  }
});

module.exports = router;