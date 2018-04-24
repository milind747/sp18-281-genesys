var express = require('express');
var router = express.Router();
var ejs = require("ejs");
var LocalStorage = require('node-localstorage').LocalStorage;
localStorage = new LocalStorage('./scratch');
var axios = require ('axios');

var path ='http://18.218.174.63:3001/'

router.post('/', function(req, res) {
    console.log("Inside the login service");

    console.log("Request Body : ", req.body)

    // API Call for the get user info
    console.log("req.query.username:"+req.body.username)
    console.log("req.query.password:"+req.body.password)

    var user_details = {
        "username": req.body.username,
        "password": req.body.password
    }


    axios.post('http://18.218.174.63:3001/login',user_details)
        .then(function(response){
            console.log("response.data" , response)
            console.log("typeof response.data ",typeof response.data)
            var user_logged_in = {
                "UID" : response.data.id,
                "USERNAME" : response.data.username,
                "PARENT_ID" : "",
                "IS_PARENT" : "",

            }
            console.log("condition",response.data=='false\nnull\n')
            if(response.data=='false\nnull\n')
            {
                console.log("invalid login")
                ejs.renderFile('./views/login.ejs',{data:"3"},function(err,result){
                    console.log(err)
                    if (!err) {
                        res.end(result);
                    }

                });
            }
            else{
                console.log("successful login")
                localStorage.setItem("user_details",user_logged_in)
                console.log(" local storage data ",JSON.stringify(localStorage.getItem("user_details")))
                console.log(" ")
                console.log("logged in")

                ejs.renderFile('./views/services.ejs',{data:user_logged_in},function(err,result){
                    console.log("error ",err)
                    if (!err) {
                        res.status(200).send(result);
                    }

                })
            }
        });


});

module.exports = router;


//window.location = "/texiservice?No="+"2&userid="+document.getElementById('userid').value+"&clippercard="+document.getElementById('clippercard').value+"&username="+document.getElementById('username').value+"&password="+document.getElementById('password').value;