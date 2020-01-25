var express = require('express');
var router = express.Router();
var request = require('request');
var cheerio = require('cheerio');

//  Base URL of 1337x
var base_url = 'https://1337x.to/search/';

router.get("/1337x",function(req,res){
    //  Get the String to be Searched from URL
    var search = req.query.search;

    var jsonRes = [];
    request(base_url+search+'/1/',function(error,req,html){
        
        var $ = cheerio.load(html);

        $("tr").each((i,element)=>{    
            //  File Name
            name = $(element)
                .children()
                .eq(0) //select all the children
                .children()
                .eq(1)
                .text();
            //  Seeders
            seeders = $(element).children().eq(1).text();
            //  Leechers
            leechers = $(element).children().eq(2).text();
            //  Upload Date
            upload_date = $(element).children().eq(3).text();
            //  File Size
            file_size = $(element)
                            .children()
                            .eq(4) //select all the children
                            .clone()    //clone the element
                            .children()
                            .remove()   //remove all the children
                            .end()  //again go back to selected element
                            .text();
            //  Uploader Name
             uploader_name = $(element).children().eq(5).text();

            //  url
            url = "https://1337x.to"+$(element)
                .children()
                .eq(0) //select all the children
                .children()
                .eq(1)
                .attr("href");

             jsonRes.push({"name":name,
                "seeders":seeders,
                "leechers":leechers,
                "upload_date":upload_date,
                "size":file_size,
                "uploader":uploader_name,
                "url":url,
                "website":"1337x"});  
        })
        jsonRes.shift();
        res.json(jsonRes);
        res.end();
    })
    
})

router.get("/1337x_getMagnet",function(req,res){
    url = req.query.url;
    request("https://1337x.to"+url,function(err,req,html){
        $ = cheerio.load(html);
        magnet = $("a.l46fd9d65ce2147030dc0271928d13abca87d1c5b.l2bce15be5dd5ee93bd47c1d4dc73d144e6da9f0c.l2132091ab3ba4e64c89010c70542dcd20599bc7a").first().attr("href");
        res.json({"magnet":magnet});
        res.end();
    })
})

module.exports = router;
