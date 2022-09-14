const fs = require('fs');
const express = require('express');
const app = express();
const { urlencoded } = require('body-parser');
const mongoose = require('mongoose');
const dotenv = require('dotenv')
dotenv.config();

app.use(express.json());
app.use(urlencoded({ extended: false }));
const nocache = require('nocache'); //Disable browser caching
app.use(nocache());
app.use(express.static('./'));
app.disable('etag', false); //Disable etag to help prevent http 304 issues
port = process.env.PORT

//Start express server
app.listen(port, (error) => {
    if (error) {
        return error
    } else {
        console.log("Listening on port " + port + "... ")
    }
})

//Send the html page for the web gui
app.get('/', (req, res) => {
    res.send(fs.readFileSync('./index.html', 'utf8'));
});
