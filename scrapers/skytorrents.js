var express = require("express");
var router = express.Router();
var axios = require("axios");
var cheerio = require("cheerio");
//  Base URL of skytorrents
var BASE_URL = require("./constants").SKYTORRENTS_BASE_URL;

router.get("/skytorrents", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search;
    response = await axios.get(BASE_URL + search);
    var $ = cheerio.load(response.data);
    var jsonResponse = [];

    $("tr.result").each((index, element) => {
        //  File Name
        name = $(element).find("td:nth-child(1) a:nth-child(1)").text()
        //  Seeders
        seeders = $(element).find("td:nth-child(5)").text()
        //  Leechers
        leechers = $(element).find("td:nth-child(6)").text()
        //  Upload Date
        upload_date = $(element).find("td:nth-child(4)").text()
        //  File Size
        file_size = $(element).find("td:nth-child(2)").text()

        //  magnet
        magnet = $(element).find("td:nth-child(1) a:nth-child(7)").attr("href");
        url = "https://www.skytorrents.lol/" + $(element)
            .find("td:nth-child(1) a:nth-child(1)")
            .attr("href");
        jsonResponse.push({
            name: name,
            torrent_url : url,
            seeders: seeders,
            leechers: leechers,
            upload_date: upload_date,
            size: file_size,
            uploader : "Skytorrents",
            magnet: magnet,
            website: "Skytorrents",
        });
    });
    res.status(200).json({"data":jsonResponse}).end();
});

module.exports = router;
