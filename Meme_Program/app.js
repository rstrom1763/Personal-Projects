const fs = require("fs");
const express = require("express");
const app = express();
const MongoClient = require("mongodb").MongoClient;

port = 8081
app.use(express.json());
app.listen(port);
console.log('Listening on port ' + port + '... ')

url = "mongodb://localhost:27017/";

async function sendText(accountSid, authToken, message, fromNum, toNum) {

    const client = require('twilio')(accountSid, authToken);

    client.messages
        .create({
            body: message,
            from: fromNum,
            to: toNum
        })
        .then(message => console.log(message.sid));
}

async function get(dataBase, coll) {

    MongoClient.connect(url, { useUnifiedTopology: true }, (err, db) => {
        if (err) throw err;
        var dbo = db.db(dataBase);

        dbo.collection(coll).find({}, {}).toArray((err, result) => {
            if (err) throw err;
            console.log(result);
            db.close();
        });

    });

}

async function insert(database, coll, myobj) {

    MongoClient.connect(url, { useUnifiedTopology: true }, (err, db) => {
        if (err) throw err;
        var dbo = db.db(database);
        dbo.collection(coll).insertOne(myobj, { unique: true }, (err, res) => {
            if (err) throw err;;
            console.log("Document Inserted");
            db.close();
        });
    });
};

app.post('/', (req, res) => {
    console.log(req.body);
    res.send("Thank you)");
});

//insert('mydb', 'Customers', { UserName: 'Rstrom1763', FirstName: 'Ryan', LastName: 'Strom', Email: 'Rstrom1763@gmail.com' });
//get('mydb','Customers');
sendText("ACd1eccbf09f918b263f0e6aaf133127d9", "7da7a17e19f658f1951ca71f337dfa8c", "Hey Babe I love You♥", 15098225903, 8168359580)