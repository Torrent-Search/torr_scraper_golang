const express = require("express");
const router = express.Router();
const cheerio = require("cheerio");
const axios = require("axios");
const BASE_URL = require("./constants").BASE_URL_1337X;

async function get1337xSearchResult(search) {
    var response = await axios.get(BASE_URL + search + "/1/");
    // console.log(response)
    $ = cheerio.load(response.data);
    var jsonRes = [];
    $("tr").each((i, element) => {
        //  File Name
        name = $(element).find("td.coll-1.name").text();
        //  Seeders
        seeders = $(element).find("td.coll-2.seeds").text();
        //  Leechers
        leechers = $(element).find("td.coll-3.leeches").text();
        //  Upload Date
        upload_date = $(element).find("td.coll-date").text();
        //  File Size
        file_size = $(element)
            .find("td:nth-child(5)")
            .clone()
            .children()
            .remove()
            .end()
            .text();
        //  Uploader Name
        uploader_name = $(element).find("td:nth-child(6)").text();

        //  url
        url =
            "https://1337x.to" +
            $(element).find("td.coll-1.name > a:nth-child(2)").attr("href");

        jsonRes.push({
            name: name,
            torrent_url: url,
            seeders: seeders,
            leechers: leechers,
            upload_date: upload_date,
            size: file_size,
            uploader: uploader_name,
            magnet: "",
            website: "1337x",
        });
    });
    return jsonRes;
}

router.get("/1337x", async function (req, res) {
    //  Get the String to be Searched from URL
    var search = req.query.search;
    var jsonResult = await get1337xSearchResult(search);
    jsonResult.shift();
    res.status(200).json({ data: jsonResult }).end();
});

router.get("/1337x_mg", async function (req, res) {
    url = req.query.url;
    response = await axios.get(url);
    $ = cheerio.load(response.data);
    magnet = $(".clearfix ul li a").attr("href");
    res.status(200).json({ magnet: magnet }).end();
});

module.exports = router;
