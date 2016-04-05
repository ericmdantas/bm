"use strict";

const restify = require('restify');
const PORT = 3000;

const server = restify.createServer({});

server.get('/', function (req, res, next) {
  res.send("yo!");
  return next();
});

server.listen(PORT, () => {
  console.log(PORT);
});
