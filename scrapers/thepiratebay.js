const express = require("express");
const cherrio = require("cheerio");
const router = express.Router();
const puppeteer = require("puppeteer");
BASE_URL = require("./constants").THEPIRATEBAY_BASE_URL;

router.get("/thepiratebay", async function (req, res) {
    search = req.query.search;

    const browser = await puppeteer.launch({
        headless: true
    });
    const page = await browser.newPage();
    await page.goto(BASE_URL + search + "/0/99/0");
    const content = await page.content();
    if (content != undefined) {
        let $ = cherrio.load(content);
        let jsonRes = [];
        $("#st").each((index, element) => {
            file_name = $(element).children().eq(1).text();
            seeders = $(element).children().eq(4).text();
            leechers = $(element).children().eq(6).text();
            upload_date = $(element).children().eq(2).text();
            size = $(element).children().eq(5).text();
            uploader_name = $(element).children().eq(7).text();
            magnet_link = $(element)
                .find(".item-icons")
                .children()
                .eq(0)
                .attr("href");

            jsonRes.push({
                name: file_name,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: size.replace("Size ", ""),
                uploader: uploader_name.replace("ULed by ", ""),
                magnet: magnet_link,
                website: "thepiratebay",
            });
        });
        jsonRes.shift();
        res.json(jsonRes);
        res.end();
    }
});

module.exports = router;
