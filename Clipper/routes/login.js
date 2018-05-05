var express = require('express');
var ejs = require("ejs");
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res) {
    console.log("Neha")
    var data=req.query.No;
    if (!req.query.No)
    {
        data="1"
    }
    else
    {
        data=req.query.No;
    }
    console.log("data:", data)
    ejs.renderFile('./views/login.ejs',{data:data},function(err,result){
        console.log(err)
        if (!err) {
            res.end(result);
        }

    });
});

module.exports = router;
