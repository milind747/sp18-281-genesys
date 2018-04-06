var express = require('express');
var ejs = require("ejs");
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res) {
  ejs.renderFile('./views/services.ejs',function(err,result){
    if (!err) {
             res.end(result);
    }
  
});
});

module.exports = router;
