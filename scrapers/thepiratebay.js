const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
//  Base URL of The Pirate Bay
const BASE_URL = require("./constants").THEPIRATEBAY_BASE_URL;

router.get("/thepiratebay", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search.trim();
    var response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("tr");
    if (selector.length > 1) {
        selector.each((index, element) => {
            //  File Name
            name = $(element).find("td:nth-child(2) div a").text();
            //  Seeders
            seeders = $(element).find("td:nth-child(3)").text();	
            leechers =  $(element).find("td:nth-child(4)").text();
            file_info = $(element).find("font.detDesc").text();	

            upload_date_temp = String(file_info.split(",")[0]).replace("Uploaded ","");	

            upload_date = upload_date_temp.replace(/\s/g,"-");	

            file_size = String(file_info.split(",")[1])
                                    .replace("Size ","")	
                                    .slice(1);	

            uploader_name = String(file_info.split(",")[2])
                                    .replace("ULed by ","")	
                                    .slice(1);

            //  magnet
            magnet = $(element)
                .find("td:nth-child(2) > a:nth-child(2)")
                .attr("href");
            url = $(element).find("td:nth-child(2) div a").attr("href");
            jsonResponse.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date.replace("Uploaded ","").replace(" ","-"),
                size: file_size,
                uploader: uploader_name.replace("ULed by ",""),
                magnet: magnet,
                website: "The Pirate Bay",
            });
        });
        jsonResponse.shift();
        res.status(200).json({ data: jsonResponse }).end();
    } else {
        res.status(204).end();
    }
});

module.exports = router;
