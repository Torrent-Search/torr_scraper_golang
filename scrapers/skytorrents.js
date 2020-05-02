const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
//  Base URL of skytorrents
const BASE_URL = require("./constants").SKYTORRENTS_BASE_URL;

router.get("/skytorrents", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search;
    var response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("tr.result");
    if (selector.length > 0) {
        selector.each((index, element) => {
            //  File Name
            name = $(element).find("td:nth-child(1) a:nth-child(1)").text();
            //  Seeders
            seeders = $(element).find("td:nth-child(5)").text();
            //  Leechers
            leechers = $(element).find("td:nth-child(6)").text();
            //  Upload Date
            upload_date = $(element).find("td:nth-child(4)").text();
            //  File Size
            file_size = $(element).find("td:nth-child(2)").text();

            //  magnet
            magnet = $(element)
                .find("td:nth-child(1) a:nth-child(7)")
                .attr("href");
            url =
                "https://www.skytorrents.lol/" +
                $(element).find("td:nth-child(1) a:nth-child(1)").attr("href");
            jsonResponse.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: "Skytorrents",
                magnet: magnet,
                website: "Skytorrents",
            });
        });
        res.status(200).json({ data: jsonResponse }).end();
    } else {
        res.status(204).end();
    }
});

module.exports = router;
