var express = require("express");
var router = express.Router();
var request = require("request");
var cheerio = require("cheerio");
//  Base URL of skytorrents
var BASE_URL = "https://www.skytorrents.lol/?query="

router.get("/skytorrents",function(req,res){
    //  Get the Item to be searched from Query
    var search = req.query.search;

    request(BASE_URL+search , function(err , request , body){
        var $ = cheerio.load(body);
        var jsonRespons = [];

        $("tr.result").each((index,element)=>{
            //  File Name
            name = $(element)
                .children()
                .eq(0) //select all the children
                .children()
                .eq(0)
                .text();
            //  Seeders
            seeders = $(element).children().eq(4).text();
            //  Leechers
            leechers = $(element).children().eq(5).text();
            //  Upload Date
            upload_date = $(element).children().eq(3).text();
            //  File Size
            file_size = $(element)
                            .children()
                            .eq(1) //select all the children
                            .text();

            //  url
            magnet = $(element)
                        .children()
                        .eq(0)
                        .find("a")
                        .eq(2)
                        .attr("href");

            jsonRespons.push({"name":name,
                "seeders":seeders,
                "leechers":leechers,
                "upload_date":upload_date,
                "size":file_size,
                "magnet":magnet,
                "website":"Skytorrents"});  
        })
        res.json(jsonRespons);
        res.end()
    })
})

function checkRegex(str1)
{
    return str1.match(/magnet:\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i)
}
module.exports = router;