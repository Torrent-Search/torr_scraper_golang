const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
//  Base URL of skytorrents
const BASE_URL = require("./constants").LIMETORRENTS_BASE_URL;
const isMagnet = require("./utils/misc_utils.js").isMagnet;

router.get("/limetorrents", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search.trim();
    response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("table.table2 tbody tr");
    if (selector.length > 1) {
        selector.each((index, element) => {
            //  File Name
            name = $(element).find("td:nth-child(1)").text();
            //  Seeders
            seeders = $(element).find("td:nth-child(4)").text();
            //  Leechers
            leechers = $(element).find("td:nth-child(5)").text();
            //  Upload Date
            upload_date = $(element)
                .find("td:nth-child(2)")
                .text()
                .split(" - ")[0];
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
                uploader: "--",
                magnet: magnet,
                website: "Limetorrents",
            });
        });
        jsonResponse.shift();
        res.status(200).json({ data: jsonResponse }).end();
    } else {
        res.status(204).end();
    }
});

router.get("/limetorrents_mg", async function (req, res) {
    url = req.query.url;
    response = await axios.get(url).catch((err) => {
        console.log(err);
    });
    $ = cheerio.load(response.data);
    magnet = $(
        "#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(13) > div > p > a"
    ).attr("href");
    if (isMagnet(magnet)) {
        res.status(200).json({ magnet: magnet }).end();
    } else {
        res.status(204).end();
    }
});

module.exports = router;
