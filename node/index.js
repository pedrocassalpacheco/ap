require('@instana/collector')();
const express = require('express')
const https = require('http');
const app = express()
const port = 3000

var mysql = require('mysql')
var connection = mysql.createConnection({
  host: 'localhost',
  user: 'root',
  password: 'password',
  database: 'people'
})
app.get('/home', function (req, res) {
//connection.connect()

connection.query('SELECT * FROM employee', function (err, rows, fields) {
  if (err) throw err
  
  console.log(rows)
})

//connection.end()
	res.send("Hello World!")
})

app.get('/step1', function (req, res) {
  https.get('http://localhost:8080/step2', (resp) => {
    let data = '';

    resp.on('data', (chunk) => {
      data += chunk;
    });

    // The whole response has been received. Print out the result.
    resp.on('end', () => {
      //console.log(JSON.parse(data));
      console.log(data);
    });

  }).on("error", (err) => {
    console.log("Error: " + err.message);
  });

})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
