var express = require('express');
var router = express.Router();
var ejs = require("ejs");


var axios = require ('axios');

var path ='http://localhost:3003/'

/* GET home page. */
router.get('/', function(req, res) {
    ejs.renderFile('./views/addBal.ejs',function(err,result){
        
        console.log(err)
        if (!err) {
            res.end(result);
        }

    });
});

// router.post('/addBalance', function(req, res) {
//     ejs.renderFile('./views/paymentSuccess.ejs',function(err,result){
        
//         console.log(err)
//         if (!err) {
//             res.end(result);
//         }

//     });
// });

router.post('/addBalance', function(req, res) {
    console.log("Inside the login service");

    console.log("Request Body : ", req.body)

    // API Call for the get user info
    //console.log("req.query.username:"+req.body.username)
    //console.log("req.query.password:"+req.body.password)


    axios.post(path+'/addAmount?parentid='+req.body.parentid+'&amt='+req.body.amt)
        .then(function(response){
            console.log("response.data" , response.data)

            ejs.renderFile('./views/paymentSuccess.ejs',function(err,result){
                console.log("error ",err)
                if (!err) {
                    res.status(200).send(result);
                }

            })});
});

module.exports = router;


//window.location = "/texiservice?No="+"2&userid="+document.getElementById('userid').value+"&clippercard="+document.getElementById('clippercard').value+"&username="+document.getElementById('username').value+"&password="+document.getElementById('password').value;