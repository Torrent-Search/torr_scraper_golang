var express = require('express');
var cherrio = require('cheerio');
var request = require('request');
var router = express.Router();

router.get("/thepiratebay",function(req,res){
    search = req.query.search;
    baseUrl = `https://thepiratebay.org/search/${search}/0/99/0`;
    request(baseUrl,function(err,request,html){
        $ = cherrio.load(html);
        jsonRes = [];

        $("tr").each((index,element)=>{
            file_name = $(element).find("a.detLink").text();
            seeders = $(element).children().eq(2).text();
            leechers =  $(element).children().eq(3).text();
            file_info = $(element).find("font.detDesc").text();

            upload_date_temp = String(file_info.split(",")[0])
                                    .replace("Uploaded ","");

            upload_date = upload_date_temp
                                    .replace(/\s/g,"-");

            size = String(file_info.split(",")[1])
                                    .replace("Size ","")
                                    .slice(1);

            uploader_name = String(file_info.split(",")[2])
                                    .replace("ULed by ","")
                                    .slice(1);

            magnet_link = $(element)
                                    .children()
                                    .eq(1)
                                    .children()
                                    .eq(1)
                                    .attr("href");
        
            
            jsonRes.push({"name":file_name,
                "seeders":seeders,
                "leechers":leechers,
                "upload_date":upload_date.replace("Uploaded ","").replace(" ","-"),
                "size":size.replace("Size ",""),
                "uploader":uploader_name.replace("ULed by ",""),
                "magnet":magnet_link,
                "website":"thepiratebay"}); 
        })
        jsonRes.shift();
        res.json(jsonRes);
        res.end();
    })
    
})

module.exports = router;