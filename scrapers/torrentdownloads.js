const express = require("express");
const router = express.Router();
const cheerio = require("cheerio");
const axios = require("axios");
const BASE_URL = require("./constants").TORRENTDOWNLOADS_BASE_URL;
const isMagnet = require("./utils/misc_utils.js").isMagnet;

router.get("/torrentdownloads", async function (req, res) {
    //  Get the Item to be searched from Query
    var search = req.query.search.trim();
    response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
    });
    var $ = cheerio.load(response.data);
    var jsonResponse = [];
    var selector = $("body div table:nth-child(8) tr");
    if (selector.length > 1) {
        selector.each((index, element) => {
            //  File Name
            name = $(element).find("td.tdleft").find("a").text();
            //  Seeders
            seeders = $(element).find("td:nth-child(4)").text();
            //  Leechers
            leechers = $(element).find("td:nth-child(5)").text();
            //  Upload Date
            upload_date = $(element).find("td:nth-child(2)").text();
            //  File Size
            file_size = $(element).find("td:nth-child(3)").text();

            //  magnet
            magnet = "";
            url =
                "https://www.torrentdownload.info" +
                $(element).find("td.tdleft").find("a").attr("href");
            jsonResponse.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: "--",
                magnet: magnet,
                website: "torrentdownloads",
            });
        });
        jsonResponse.shift();
        if(jsonResponse[0].name.includes('Download Anonymously')){
        res.status(204).end();
        }
        else{
        res.status(200).json({ data: jsonResponse }).end();
        }
    } else {
        res.status(204).end();
    }
});

router.get("/torrentdownloads_mg", async function (req, res) {
    url = req.query.url;
    response = await axios.get(url).catch((err) => {
        console.log(err);
        res.status(204).end();
    });
    if (response != undefined) {
        $ = cheerio.load(response.data);
        magnet = $("tbody tr:nth-child(5) td span a").attr("href");
        if (isMagnet(magnet)) {
            res.status(200).json({ magnet: magnet }).end();
        } else {
            res.status(204).end();
        }
    } else {
        res.status(204).end();
    }
});

module.exports = router;
