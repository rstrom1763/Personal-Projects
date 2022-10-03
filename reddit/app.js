const fs = require('fs');
const express = require('express');
const app = express();
const { urlencoded } = require('body-parser');
const mongoose = require('mongoose');
const dotenv = require('dotenv')
dotenv.config();

//Connect to mongodb
mongoose.connect('mongodb://' + process.env.MONGO_HOST + '/userdata?retryWrites=true&w=majority', { user: process.env.MONGO_USERNAME, pass: process.env.MONGO_PASSWORD, useNewUrlParser: true, useUnifiedTopology: true }, () => {
    console.log("Connected to MongoDB! ");
});

const commentSchema = new mongoose.Schema({
    author: { type: String, required: true },
    body: { type: String, required: true },
    id: { type: String, unique: true, required: true },
    subreddit: { type: String, required: true },
    subreddit_id: { type: String, required: true }
});
const Comment = mongoose.model('Comment', commentSchema, 'comments');

const authorSchema = new mongoose.Schema({
    username: { type: String, required: true, unique: true },
    comments: { type: Array, required: true }
});
const Author = mongoose.model("Author", authorSchema, "authors");

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
});

//Send the html page for the web gui
app.get('/', (req, res) => {
    res.send(fs.readFileSync('./index.html', 'utf8'));
});

/*
app.post('/uploadComment', (req, res) => {
    Comment.create(
        {
            author: req.body.author,
            body: req.body.body,
            id: req.body.id,
            subreddit: req.body.subreddit,
            subreddit_id: req.body.subreddit_id
        }, (err) => { //Callback Function
            if (err) {
                res.status(208)
                res.send(err)
            } else {
                res.status(200)
                res.send("Succesfully created comment: " + req.body.id)
            }
        }
    )
});
*/

app.post('/uploadComment', (req, res) => {
    Author.create({ //If the author doesn't already exist, create them
        username: req.body.author,
        comments: [req.body.body]
    }, (err) => {
        //If author already exists, add the comment to their comments array
        if (err) {
            Author.findOne({ username: req.body.author }, (author) => {
                console.log(author)
                author.comments.push(req.body.body);
                author.save();
                res.send("Comment added!");
            });
        } else {
            res.send("Author document created!")
        }
    });

});

Author.findOne({ username: "syhanyth" }), (data) => {
    console.log(data)
};