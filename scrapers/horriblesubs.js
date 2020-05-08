const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
//  Base URL of skytorrents
const BASE_URL = require("./constants").HORRIBLESUBS_BASE_URL;

router.get("/horriblesubs", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search.trim();
    var response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("tr");
    if (selector.length > 0) {
        selector.each((index, element) => {
            //  File Name
            if ($(element).find("td:nth-child(2) a").length == 2) {
                name = $(element).find("td:nth-child(2) a").eq(1).text();
            } else {
                name = $(element).find("td:nth-child(2) a").text();
            }

            //  Seeders
            seeders = $(element).find("td:nth-child(6)").text();
            //  Leechers
            leechers = $(element).find("td:nth-child(7)").text();
            //  Upload Date
            upload_date = $(element)
                .find("td:nth-child(5)")
                .text()
                .split(" ")[0];
            //  File Size
            file_size = $(element).find("td:nth-child(4)").text();

            //  magnet
            magnet = $(element)
                .find("td:nth-child(3) a:nth-child(2)")
                .attr("href");
            url =
                "https://nyaa.si/" +
                $(element).find("td:nth-child(2) a").attr("href");
            jsonResponse.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: "HorribleSubs",
                magnet: magnet,
                website: "Horriblesubs",
            });
        });
        jsonResponse.shift();
        res.status(200).json({ data: jsonResponse }).end();
    } else {
        res.status(204).end();
    }
});

module.exports = router;
