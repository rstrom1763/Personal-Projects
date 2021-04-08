const fs = require("fs");
const express = require("express");
const app = express();
const MongoClient = require("mongodb").MongoClient;

//app.use(express.json());
//app.listen(8081);

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

//insert("mydb", 'customers', { userName: 'Rstrom1763', firstName: 'Ryan', lastName: 'Strom', phone: '816-787-7716', address: '53398 Lawrence Ct McConnell AFB KS 67221', email: 'rstrom1763@gmail.com', favoriteColor: 'Red' });
//get('mydb', 'customers');
sendText('ACd1eccbf09f918b263f0e6aaf133127d9', 'f22582e73ef9724117c62a1d52546845', 'This is the ship that made the Kessel Run in fourteen parsecs?', '+15098225903', '(816)-787-7716')