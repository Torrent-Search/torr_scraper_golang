var express = require("express");
var router = express.Router();
var axios = require("axios");
var cheerio = require("cheerio");
//  Base URL of skytorrents
var BASE_URL = require("./constants").LIMETORRENTS_BASE_URL;

router.get("/limetorrents", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search;
    response = await axios.get(BASE_URL + search);
    var $ = cheerio.load(response.data);
    var jsonResponse = [];

    $("table.table2 tbody tr").each((index, element) => {
        //  File Name
        name = $(element).find("td:nth-child(1)").text();
        //  Seeders
        seeders = $(element).find("td:nth-child(4)").text();
        //  Leechers
        leechers = $(element).find("td:nth-child(5)").text();
        //  Upload Date
        upload_date = $(element).find("td:nth-child(2)").text().split(" - ")[0];
        //  File Size
        file_size = $(element).find("td:nth-child(3)").text();

        //  magnet
        magnet = "";
        url =
            "https://www.limetorrents.info" +
            $(element)
                .find("td.tdleft div.tt-name a:nth-child(2)")
                .attr("href");
        jsonResponse.push({
            name: name,
            torrent_url: url,
            seeders: seeders,
            leechers: leechers,
            upload_date: upload_date,
            size: file_size,
            uploader: "Limetorrents",
            magnet: magnet,
            website: "Limetorrents",
        });
    });
    jsonResponse.shift();
    res.status(200).json({ data: jsonResponse }).end();
});

router.get("/limetorrents_mg", async function (req, res) {
    url = req.query.url;
    response = await axios.get(url);
    $ = cheerio.load(response.data);
    magnet = $(
        "#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(13) > div > p > a"
    ).attr("href");
    res.status(200).json({ magnet: magnet }).end();
});

module.exports = router;
