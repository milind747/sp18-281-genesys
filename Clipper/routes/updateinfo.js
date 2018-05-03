var express = require('express');
var router = express.Router();
var ejs = require("ejs");

var axios = require ('axios');


router.get('/', function(req, res) {
  console.log("Inside the update user details");

  
  ejs.renderFile('./views/updateinfo.ejs',function(err,result){
    console.log(err)
    if (!err) {
        res.status(200).send(result);
    }
  
    });
});
//your api
router.post('/', function(req, res) {
    console.log("Inside the port of uber service");


    // API Call for the get user info
    //console.log("req.query.username:"+req.body.username)
    //console.log("req.query.password:"+req.body.password)

    var user_details = {
        "username": req.body.username,
        "phone": req.body.phone,
        "email":req.body.email

        
    }


    axios.post('http://localhost:3000/postuserinfo',user_details)
        .then(function(response){
            console.log("response.data" , response)
            console.log("typeof response.data ",typeof response.data)
            
            console.log("condition",response.data=='false\nnull\n')
            
                console.log("invalid login")
                ejs.renderFile('./views/updateinfo.ejs',{data:"3"},function(err,result){
                    console.log(err)
                    if (!err) {
                        res.end(result);
                    }

                });
                console.log("successful login")
                localStorage.setItem("user_details",user_logged_in)
                console.log(" local storage data ",JSON.stringify(localStorage.getItem("user_details")))
                console.log(" ")
                console.log("logged in")

                ejs.renderFile('./views/updateinfo.ejs',function(err,result){
                    console.log("error ",err)
                    if (!err) {
                        res.status(200).send(result);
                    }

                })
            }
        );


});

module.exports = router;

