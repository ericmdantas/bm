"use strict";

const koa = require('koa');
const app = koa();

const PORT = 3000;

console.log(PORT);

app.use(function*() {
  this.body = 'yo!';
});

app.listen(PORT);
