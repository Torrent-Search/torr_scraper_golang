const express = require("express");
const cherrio = require("cheerio");
const router = express.Router();
const puppeteer = require("puppeteer");
BASE_URL = require("./constants").THEPIRATEBAY_BASE_URL;

router.get("/thepiratebay", async function (req, res) {
    search = req.query.search;

    const browser = await puppeteer.launch({
        headless: true,
    });
    const page = await browser.newPage();
    await page.goto(BASE_URL + search);
    const content = await page.content();
    if (content != undefined) {
        let $ = cherrio.load(content);
        let jsonResponse = [];
        $("#st").each((index, element) => {
            file_name = $(element)
                .find("span.list-item.item-name.item-title")
                .text();
            seeders = $(element).find("span.list-item.item-seed").text();
            leechers = $(element).find("span.list-item.item-leech").text();
            upload_date = $(element)
                .find("span.list-item.item-uploaded")
                .text();
            size = $(element).find("span.list-item.item-size").text();
            uploader_name = $(element).find("span.list-item.item-user").text();
            magnet_link = $(element).find("span.item-icons a").attr("href");
            url =
                "https://thepiratebay.org" +
                $(element)
                    .find("span.list-item.item-name.item-title a")
                    .attr("href");
            jsonResponse.push({
                name: file_name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: size.replace("Size ", ""),
                uploader: uploader_name.replace("ULed by ", ""),
                magnet: magnet_link,
                website: "thepiratebay",
            });
        });
        res.json({ data: jsonResponse });
        res.end();
    }
});

module.exports = router;
