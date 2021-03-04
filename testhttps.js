var https = require('https');
var fs = require('fs');
const options = {
  key: fs.readFileSync('key.pem'),
  cert: fs.readFileSync('cert.pem')
};

https.createServer(options, function (req, res) {
  var body = "";
  req.on('data', function (chunk) {
    body += chunk;
  });
  req.on('end', function () {
	  
	var fs = require('fs');
	body = body.replace(/(\r\n|\n|\r)/gm,"");
    console.log('POSTed: ' + body);
	var stream = fs.createWriteStream('test.txt',{'flags':'a'});

	stream.write(body + "\n");
	
	stream.end();
    res.writeHead(200);
    res.end("end");
  });
}).listen(8081);