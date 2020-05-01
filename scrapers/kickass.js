const express = require("express");
const router = express.Router();
const axios = require("axios");
const cheerio = require("cheerio");
const BASE_URL = require("./constants").KICKASS_BASE_URL;

async function getKickassSearchResult(search) {
    const response = await axios.get(BASE_URL + search);
    const $ = cheerio.load(response.data);
    const odd_rows = $("tr.odd , tr.even");
    const jsonRes = [];
    odd_rows.each((i, element) => {
        name = $(element).find(".cellMainLink").text().replace("\n", "");
        uploader_name = $(element)
            .find("td:nth-child(3)")
            .text()
            .replace("\n", "");
        file_size = $(element).find("td:nth-child(2)").text().replace("\n", "");
        upload_date = $(element)
            .find("td:nth-child(4)")
            .text()
            .replace("\n", "");
        //  Seeders
        seeders = $(element).find("td:nth-child(5)").text().replace("\n", "");
        //  Leechers
        leechers = $(element).find("td:nth-child(6)").text().replace("\n", "");
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
            magnet: "",
            website: "Kickass",
        });
    });

    return jsonRes;
}

router.get("/kickass_mg", async (req, res) => {
    var url = req.query.url;
    var response = await axios.get(url);
    var $ = cheerio.load(response.data);
    magnet = $(".kaGiantButton ").attr("href");
    res.status(200).json({"magnet":magnet}).end();
});

router.get("/kickass", async function (req, res) {
    console.log(BASE_URL);
    //  Get the String to be Searched from URL
    var search = req.query.search;
    jsonResult = await getKickassSearchResult(search);
    res.status(200).json({ data: jsonResult }).end();
});

module.exports = router;
