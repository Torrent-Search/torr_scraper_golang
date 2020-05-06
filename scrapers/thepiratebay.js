const express = require("express");
const cherrio = require("cheerio");
const router = express.Router();
// const puppeteer = require("puppeteer");
const Browser = require("zombie");
const BASE_URL = require("./constants").THEPIRATEBAY_BASE_URL;

router.get("/thepiratebay", function (req, res) {
    search = req.query.search;
    var browser = new Browser();
    browser.visit(BASE_URL + search)//.catch(err=>console.log(err));
    content = browser.html()//.catch(err=>console.log(err));
    if (content != undefined) {
        var $ = cherrio.load(content);
        var jsonResponse = [];
        var selector = $("#st");
        if (selector.length > 0) {
            selector.each((index, element) => {
                file_name = $(element)
                    .find("span.list-item.item-name.item-title")
                    .text();
                seeders = $(element).find("span.list-item.item-seed").text();
                leechers = $(element).find("span.list-item.item-leech").text();
                upload_date = $(element)
                    .find("span.list-item.item-uploaded")
                    .text();
                size = $(element).find("span.list-item.item-size").text();
                uploader_name = $(element)
                    .find("span.list-item.item-user")
                    .text();
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
            try{
            res.status(200).json({ data: jsonResponse }).end();}
            catch(e){
                console.log(e)
            }
        } else {
            res.status(204).end();
        }
        browser.tabs.closeAll();
    } else {
        res.status(204).end();
        browser.tabs.closeAll();
    }
});

module.exports = router;
