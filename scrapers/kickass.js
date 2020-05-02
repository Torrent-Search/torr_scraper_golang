const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
const BASE_URL = require("./constants").KICKASS_BASE_URL;
const isMagnet = require("./utils/misc_utils.js").isMagnet

router.get("/kickass_mg", async (req, res) => {
    var url = req.query.url;
    var response = await axios.get(url).catch(err=>{console.log(err)});
    var $ = cheerio.load(response.data);
    magnet = $("a.kaGiantButton").attr("href");
    if (isMagnet(magnet)) {
        res.status(200).json({ magnet: magnet }).end();
    } else {
        res.status(204).end();
    }
});

router.get("/kickass", async function (req, res) {
    //  Get the String to be Searched from URL
    const search = req.query.search;
    const html_response = true;
    const response = await axios.get(BASE_URL + search).catch((err) => {
        console.log(err);
        html_response = false;
    });
    const $ = cheerio.load(response.data);
    const selector = $("tr.odd , tr.even");
    const jsonResult = [];
    const content_avail = $("span[itemprop=name]").length;
    if (html_response) {
        if (content_avail == 0) {
            selector.each((i, element) => {
                name = $(element)
                    .find(".cellMainLink")
                    .text()
                    .replace("\n", "");
                uploader_name = $(element)
                    .find("td:nth-child(3)")
                    .text()
                    .replace("\n", "");
                file_size = $(element)
                    .find("td:nth-child(2)")
                    .text()
                    .replace("\n", "");
                upload_date = $(element)
                    .find("td:nth-child(4)")
                    .text()
                    .replace("\n", "");
                //  Seeders
                seeders = $(element)
                    .find("td:nth-child(5)")
                    .text()
                    .replace("\n", "");
                //  Leechers
                leechers = $(element)
                    .find("td:nth-child(6)")
                    .text()
                    .replace("\n", "");
                url =
                    "https://kickasstorrents.to" +
                    $(element).find(".cellMainLink").attr("href");
                jsonResult.push({
                    name: name,
                    torrent_url: url,
                    seeders: seeders,
                    leechers: leechers,
                    upload_date: upload_date,
                    size: file_size,
                    uploader: uploader_name,
                    magnet: "",
                    website: "Kickass",
                });
            });
            res.status(200).json({ data: jsonResult }).end();
        } else {
            res.status(204).end();
        }
    } else {
        res.status(204).end();
    }
});


module.exports = router;
