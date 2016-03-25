"use strict";

const http = require('http');
const PORT = 3000;

console.log('running on port', PORT)

http.createServer((req, res) => {
  res.end('yo!');
}).listen(PORT);
