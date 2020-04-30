const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
const BASE_URL = require("./constants").KICKASS_BASE_URL;

async function getKickassSearchResult(search) {
    const response = await axios.get(BASE_URL + search);
    const $ = cheerio.load(response.data);
    const odd_rows = $(".odd");
    const even_rows = $(".even");
    const jsonRes = [];
    if (odd_rows.length > 0) {
        odd_rows.each((i, element) => {
            name = $(element).find(".cellMainLink").text().replace("\n", "");
            // uploader_name = $(element).find(".plain").text().substring(2);
            uploader_name = $(
                "table > tbody > tr:nth-child(2) > td:nth-child(1) > div.torrentname > div > span > a"
            ).text();
            file_size = $(element).children().eq(1).text().replace("\n", "");
            upload_date = $(element).children().eq(3).text().replace("\n", "");
            //  Seeders
            seeders = $(element).children().eq(4).text().replace("\n", "");
            //  Leechers
            leechers = $(element).children().eq(5).text().replace("\n", "");
            url =
                "https://kickasstorrents.to" +
                $(element).find(".cellMainLink").attr("href");
            jsonRes.push({
                name: name,
                torrent_url: url,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: uploader_name,
                magnet : "",
                website: "Kickass",
            });
        });
    }
    if (even_rows.length > 0) {
        even_rows.each((i, element) => {
            name = $(element).find(".cellMainLink").text().replace("\n", "");
            uploader_name = $(element).find(".plain").text().substring(2);
            file_size = $(element).children().eq(1).text().replace("\n", "");
            upload_date = $(element).children().eq(3).text().replace("\n", "");
            //  Seeders
            seeders = $(element).children().eq(4).text().replace("\n", "");
            //  Leechers
            leechers = $(element).children().eq(5).text().replace("\n", "");
            url =
                "https://kickasstorrents.to" +
                $(element).find(".cellMainLink").attr("href");
            jsonRes.push({
                name: name,
                seeders: seeders,
                leechers: leechers,
                upload_date: upload_date,
                size: file_size,
                uploader: uploader_name,
                url: url,
                website: "Kickass",
            });
        });
    }
    return jsonRes;
}

router.get("/kickassMagnet", async (req, res) => {
    let url = req.query.url;
    let response = await axios.get(url);
    let $ = cheerio.load(response.data);
    magnet = $(".kaGiantButton ").attr("href");
    res.send(magnet);
    res.end();
});

router.get("/kickass", async function (req, res) {
    console.log(BASE_URL);
    //  Get the String to be Searched from URL
    let search = req.query.search;
    jsonResult = await getKickassSearchResult(search);
    res.json({"data":jsonResult});
    res.end();
});

module.exports = router;
