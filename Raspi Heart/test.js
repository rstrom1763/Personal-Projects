var http = require('http');

http.createServer(function (req, res) {
  var body = "";
  req.on('data', function (chunk) {
    body += chunk;
  });
  req.on('end', function () {

    if(body !== ""){
    var fs = require('fs');
    body = body.replace(/(\r\n|\n|\r)/gm, "");
    console.log('POSTed: ' + body);
    var stream = fs.createWriteStream('test.txt', { 'flags': 'a' });

    stream.write(body + "\n");

    stream.end();
    res.writeHead(200);
    res.end("end");
    }
    else{

      res.writeHead(200);
      var fs = require('fs');
      var stream = fs.createReadStream('test.txt');
      res.end("steam")
      stream.end();

    }
  });
}).listen(8081);
