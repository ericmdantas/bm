'use strict';

const Hapi = require('hapi');
const server = new Hapi.Server();

const PORT = 3000;

server.connection({
    host: 'localhost',
    port: PORT
});

server.route({
    method: 'GET',
    path: '/',
    handler: function (request, reply) {
        return reply('yo!');
    }
});

server.start((err) => {
    console.log(PORT);
});
