const express = require("express");
const router = express.Router();
const request = require("request");
const cheerio = require("cheerio");
const axios = require("axios");
const BASE_URL = require("./constants").BASE_URL_1337X;

async function get1337xSearchResult(search) {
    var html;
    var response = await axios.get(BASE_URL + search + "/1/");
    // console.log(response)
    $ = cheerio.load(response.data);
    var jsonRes = [];
    $("tr").each((i, element) => {
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
            .clone() //clone the element
            .children()
            .remove() //remove all the children
            .end() //again go back to selected element
            .text();
        //  Uploader Name
        uploader_name = $(element).children().eq(5).text();

        //  url
        url =
            "https://1337x.to" +
            $(element)
                .children()
                .eq(0) //select all the children
                .children()
                .eq(1)
                .attr("href");

        jsonRes.push({
            name: name,
            torrent_url: url,
            seeders: seeders,
            leechers: leechers,
            upload_date: upload_date,
            size: file_size,
            uploader: uploader_name,
            magnet : "",
            website: "1337x",
        });
    });
    return jsonRes;
}

router.get("/1337x",async function (req, res) {
    //  Get the String to be Searched from URL
    var search = req.query.search;
    var jsonResult = await get1337xSearchResult(search);
    jsonResult.shift();
    res.json({"data":jsonResult});
    res.end();
});

router.get("/1337x_getMagnet", function (req, res) {
    url = req.query.url;
    request("https://1337x.to" + url, function (err, req, html) {
        $ = cheerio.load(html);
        magnet = $(
            "a.l46fd9d65ce2147030dc0271928d13abca87d1c5b.l2bce15be5dd5ee93bd47c1d4dc73d144e6da9f0c.l2132091ab3ba4e64c89010c70542dcd20599bc7a"
        )
            .first()
            .attr("href");
        res.json({ magnet: magnet });
        res.end();
    });
});

module.exports = router;
