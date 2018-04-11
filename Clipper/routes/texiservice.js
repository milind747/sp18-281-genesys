var express = require('express');
var router = express.Router();
var ejs = require("ejs");


router.get('/', function(req, res) {
  console.log("Inside the texiservice");

  ejs.renderFile('./views/contact.ejs',{data:req.query.No},function(err,result){
    if (!err) {
     
        res.status(200).send(result);
    }
  
});
});

module.exports = router;