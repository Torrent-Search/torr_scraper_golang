const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
//  Base URL of Torrent Galaxy
const BASE_URL = require("./constants").TORRENTGALAXY_BASE_URL;

router.get("/tgx", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search.trim();
    var response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("div.tgxtablerow");
    if (selector.length > 0) {
        selector.each((index, element) => {
            //  File Name
            name = $(element).find("div:nth-child(4)").text();
            //  Seeders
            seeders = $(element).find(
                "div:nth-child(11) span font:nth-child(1)"
            ).text();
            //  Leechers
            leechers = $(element)
                .find("div:nth-child(11) span font:nth-child(2)")
                .text();
            //  Upload Date
            upload_date = $(element).find("div:nth-child(12)").text().split(' ')[0];
            //  File Size
            file_size = $(element).find("div:nth-child(8)").text();
            //  Uploader
            uploader = $(element).find("div:nth-child(7)").text();
            //  magnet
            magnet = $(element)
                .find("div:nth-child(5) a:nth-child(2)")
                .attr("href");
            url =
                "https://torrentgalaxy.to" +
                $(element).find("div:nth-child(4) a").attr("href");

            jsonResponse.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: uploader,
                magnet: magnet,
                website: "Torrent Galaxy",
            });
        });
        res.status(200).json({ data: jsonResponse }).end();
    } else {
        res.status(204).end();
    }
});

module.exports = router;
