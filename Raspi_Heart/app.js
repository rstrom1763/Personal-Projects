const fs = require('fs');
const express = require('express');
const app = express();

app.use(express.json());
app.listen(8081);

datafile = 'test.txt';
htmlfile = 'button.html';

app.get('/button', (req, res) => {
    status = "";
    fs.readFile(datafile, 'utf8', (err, data) => {
        if (err) { err };
        if (data === '1') {
            fs.writeFile(datafile, '0', (err) => err);
            status = 'Off';
        } else {
            fs.writeFile(datafile, '1', (err) => err);
            status = 'On';
        }
        res.send(status);
    });

});

app.get('/check', (req, res) => {

    fs.readFile(datafile, 'utf8', (err, data) => {

        if (err) {
            res.send(err);
        } else {
            res.send(data);
        }

    });

});

app.get('/', (req, res) => {

    fs.readFile(htmlfile, 'utf8', (err, data) => {

        if (err) {
            res.send(err);
        } else {
            res.send(data.body);
        }

    });

});

app.post('/', (req, res) => {

    fs.writeFile(datafile, req.body.data, (err) => err);
    res.send(req.body.data);

});
