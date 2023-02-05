const fs = require('fs'); //For file system reading and writing
const https = require("https");//Used by express to create an https server
const express = require('express'); //Library that enables us to create routes
const app = express();//Initializes express
const { urlencoded } = require('body-parser');//Library that enables routes to parse request bodies as json
const process = require('process');//Needed in order to use .env files
const dotenv = require('dotenv');//Needed in order to use .env files
dotenv.config();//Needed in order to use .env files


app.use(express.json());
app.use(urlencoded({ extended: false }));
const nocache = require('nocache'); //Disable browser caching. Prevents certain errors related to caching
app.use(nocache());
app.use(express.static('./')); //Sets the directory where express can serve static files from: EX: index.html
app.disable('etag', false); //Disable etag to help prevent http 304 issues

//If protocol environment var is https start https else start http
//Use key.pem and cert.pem for https. 
if (process.env.PROTOCOL == "https") {
    https.createServer({
        key: fs.readFileSync('key.pem'),
        cert: fs.readFileSync('cert.pem'),
    }, app).listen(process.env.PORT, "0.0.0.0", () => {
        console.log("Listening https on port " + process.env.PORT + "...")
    });
} else if (process.env.PROTOCOL === "http") {
    app.listen(process.env.PORT, "0.0.0.0", () => {
        console.log("Listening http on port " + process.env.PORT + "...")
    });
} else {
    console.log("Invalid protocol! Exiting! ");
    process.exit();
}

//Route for the pi's to send their temperature readings to
app.post('/posttemp', (req, res) => {
    log = new Date().toISOString() + " " + req.body.temp + " " + req.body.humidity + " " + req.body.pressure + "\n"
    fs.appendFile(process.env.LOG_PATH, log, (err) => {
        if (err) {
            console.log(err);
        }
    });
    res.send("Success!")
});
