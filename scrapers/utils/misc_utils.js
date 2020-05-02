function isMagnet(str) {	
    return str.match(/magnet:\?xt=urn:[a-z0-9]+:[a-z0-9]{32}/i);	
}

module.exports={isMagnet}