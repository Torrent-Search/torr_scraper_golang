const express = require("express");
const router = express.Router();
// const cheerio = require("cheerio");
// const axios = require("axios");
//  Base URL of Rargb
// const BASE_URL = require("./constants").RARGB_BASE_URL;
const rargbapi = require("rarbg-api");
const filesize = require("filesize");

router.get("/rarbg", async function (req, res) {
    var search = req.query.search.trim();
    var jsonResult = [];
    rargbapi
        .search(search, { sort: "seeders" })
        .then((response) => {
            if (response.length > 0) {
                json = JSON.parse(JSON.stringify(response));
                // json = JSON.stringify(response)
                // console.log(json);
                for (let i = 0; i < response.length; i++) {
                    jsonResult.push({
                        name: json[i].title,
                        torrent_url: "--",
                        seeders: String(json[i].seeders),
                        leechers: String(json[i].leechers),
                        upload_date: json[i].pubdate.split(" ")[0],
                        size: filesize(json[i].size),
                        uploader: "--",
                        magnet: json[i].download,
                        website: "rargb",
                    });
                }
                // console.log(json[0])
                res.json(jsonResult).status(200).end();
            } else {
                res.status(204).end();
            }
        })
        .catch((err) => {
            res.status(204).end();
        });


});

module.exports = router;
