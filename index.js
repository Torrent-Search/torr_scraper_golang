const http = require('http');
const express = require('express')
const app = express()
const PORT_NO = process.env.PORT_NO;

const Scraper_1337x = require("./scrapers/1337x.js");
const skytorrent = require("./scrapers/skytorrents.js");
const thepiratebay = require("./scrapers/thepiratebay.js");

const server = app.listen(PORT_NO || 8080, function()
{
        console.log("Listening to Port : ",server.address().port);
        console.log("Listening to Address : ",server.address().address);
});

app.use("/api",Scraper_1337x);
app.use("/api",skytorrent);
app.use("/api",thepiratebay);

app.get("/test",function(req,res){
    res.send("<h1>Tejas<h1>");
    res.end();
})  
