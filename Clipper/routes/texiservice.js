var express = require('express');
var router = express.Router();
var ejs = require("ejs");


router.get('/', function(req, res) {
  console.log("Inside the texiservice");
  ejs.renderFile('./views/contact.ejs',function(err,result){
    if (!err) {
     
        res.status(200).send(result);
    }
  
});
});

module.exports = router;