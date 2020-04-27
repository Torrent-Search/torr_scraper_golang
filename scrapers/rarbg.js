const express = require("express");
const router = express.Router();
const cheerio = require("cheerio");
const axios = require("axios");
//  Base URL of Rargb
const BASE_URL = require("./constants").RARGB_BASE_URL;

router.get("/rarbg", async function (req, res) {
    var baseUrl = BASE_URL + req.query.search;
    var jsonRes = [];
    response = await axios.get(baseUrl);
    var $ = cheerio.load(response.data);
    console.log(html);
    $("tr.lista2").each((index, element) => {
        var name = $(element).children().eq(1).children().eq(0).text();
        var upload_date = $(element).children().eq(2).text();
        var size = $(element).children().eq(3).text();
        var seeders = $(element).children().eq(4).text();
        var leechers = $(element).children().eq(5).text();
        var uploader_name = $(element).children().eq(7).text();
        var url = $(element).children().eq(1).children().eq(0).attr("href");

        jsonRes.push({
            name: name,
            seeders: seeders,
            leechers: leechers,
            upload_date: upload_date.replace("Uploaded ", "").replace(" ", "-"),
            size: size.replace("Size ", ""),
            uploader: uploader_name.replace("ULed by ", ""),
            url: url,
            website: "thepiratebay",
        });
    });
    jsonRes.shift();
    res.json(jsonRes);
    res.end();
});

module.exports = router;
