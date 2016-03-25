"use strict";

const express = require('express');
const app = express();

const PORT = 3000;

console.log(PORT);

app.get('/', (req, res) => {
  res.send('yo!');
});

app.listen(PORT);
