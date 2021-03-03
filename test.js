var http = require('http');




http.createServer(function (req, res) {
  var body = "";
  req.on('data', function (chunk) {
    body += chunk;
  });
  req.on('end', function () {
	  
	var fs = require('fs');
    console.log('POSTed: ' + body);
	var stream = fs.createWriteStream('test.txt',{'flags':'a'});
	
	stream.write(body + "\n");
	
	stream.end();
    res.writeHead(200);
    res.end("end");
  });
}).listen(8081);
