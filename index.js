const http = require("http");
const express = require("express");
const app = express();
const PORT_NO = process.env.PORT;

const Scraper_1337x = require("./scrapers/1337x.js");
const skytorrent = require("./scrapers/skytorrents.js");
const rarbg = require("./scrapers/rarbg.js");
const kickass = require("./scrapers/kickass.js");
const limetorrents = require("./scrapers/limetorrents.js");
const torrentgalaxy = require("./scrapers/torrent_galaxy.js");
const torrentdownloads = require("./scrapers/torrentdownloads.js");
const nyaa = require("./scrapers/nyaa.js");
const thepiratebay = require("./scrapers/thepiratebay.js");
const horriblesubs = require("./scrapers/horriblesubs.js");

const server = app.listen(PORT_NO, function () {
    console.log("Listening to Port : ", server.address().port);
    console.log("Listening to Address : ", server.address().address);
});

app.use("/api", Scraper_1337x);
app.use("/api", skytorrent);
app.use("/api", rarbg);
app.use("/api", kickass);
app.use("/api", limetorrents);
app.use("/api", torrentgalaxy);
app.use("/api", torrentdownloads);
app.use("/api", nyaa);
app.use("/api", thepiratebay);
app.use("/api", horriblesubs);

app.get("/", function (req, res) {
    res.status(200).end();
});
