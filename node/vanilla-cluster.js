"use strict";

const cluster = require('cluster');
const http = require('http');
const numCPUs = 4;
const PORT = 3000;

console.log(`${PORT} - pid: ${process.pid}`);

if (cluster.isMaster) {
    for (var i = 0; i < numCPUs; i++) {
        cluster.fork();
    }
} 
else {
    http.createServer(function(req, res) {
        res.end('yo!');
    }).listen(PORT);
}