var express = require("express");
var bodyParser = require("body-parser");
var cors = require("cors");
var AWS = require("aws-sdk");
AWS.config.update({ region: "us-east-1" });

var app = express();
var db = new AWS.DynamoDB();

app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get("/genres", function(req, res) {
  console.log("genre??");
  db.scan(
    {
      TableName: "Music"
    },
    function(err, data) {
      if (err) {
        console.log("ERR: ", err);
        return res.status(400).send({ message: err.message });
      }
      console.log(`data`, data.Items);
      let allGenreItems = data.Items.map(item => item.Genre.S);

      let genres = [...new Set(allGenreItems)];

      var response = {
        statusCode: 200,
        body: {
          genres
        }
      };
      return res.send(response);
    }
  );
});

app.get("/artists/for/genre", function(req, res) {
  const GEN = "Cinematic";
  console.log("genre", req.query.genre);
  db.query(
    {
      TableName: "Music",
      IndexName: "GENRE",
      KeyConditionExpression: "Genre = :genre",
      ExpressionAttributeValues: {
        ":genre": {
          S: req.query.genre || GEN
        }
      }
    },
    function(err, data) {
      if (err) {
        console.log("ERR: ", err);
        return res.status(400).send({ message: err.message });
      }

      let allArtistItems = data.Items.map(item => item.Artist.S);
      let artists = [...new Set(allArtistItems)];

      console.log("artist: ", artists);

      var response = {
        statusCode: 200,
        body: {
          artists
        }
      };
      return res.send(response);
    }
  );
});

app.get("/albums/for/artist", function(req, res) {
  const ART = "Benjamin Tissot";
  console.log("req.query.artist: ", req.query.artist);
  db.query(
    {
      TableName: "Music",
      IndexName: "ARTIST",
      KeyConditionExpression: "Artist = :artist",
      ExpressionAttributeValues: {
        ":artist": {
          S: req.query.artist || ART
        }
      }
    },
    function(err, data) {
      if (err) {
        console.log("ERR: ", err);
        return res.status(400).send({ message: err.message });
      }

      let allAlbumItems = data.Items.map(item => item.Album.S);
      let albums = [...new Set(allAlbumItems)];

      console.log("albums: ", albums);

      var response = {
        statusCode: 200,
        body: {
          albums
        }
      };
      return res.send(response);
    }
  );
});

app.get("/songs/for/album", function(req, res) {
  const ALB = "Hero";
  db.query(
    {
      TableName: "Music",
      IndexName: "ALBUM",
      KeyConditionExpression: "Album = :album",
      ExpressionAttributeValues: {
        ":album": {
          S: req.query.album || ALB
        }
      }
    },
    function(err, data) {
      if (err) {
        console.log("ERR: ", err);
        return res.status(400).send({ message: err.message });
      }

      let songs = data.Items.map(item => item.Song.S);
      console.log("songs: ", songs);

      var response = {
        statusCode: 200,
        body: {
          records: songs
        }
      };
      return res.send(response);
    }
  );
});

app.get("/song", function(req, res) {
  const SONG = "theduel";
  db.query(
    {
      TableName: "Music",
      IndexName: "SONG",
      KeyConditionExpression: "Song = :song",
      ExpressionAttributeValues: {
        ":song": {
          S: req.query.song || SONG
        }
      }
    },
    function(err, data) {
      if (err) {
        console.log("ERR: ", err);
        return res.status(400).send({ message: err.message });
      }

      console.log("SDATA:", data);

      let songURL =
        data.Items && data.Items.length > 0 ? data.Items[0].url.S : null;

      console.log("URL: ", songURL);

      var response = {
        statusCode: 200,
        body: {
          records: songURL
        }
      };
      return res.send(response);
    }
  );
});

var server = app.listen(8081, function() {
  var host = server.address().address;
  var port = server.address().port;

  console.log("Example app listening at http://%s:%s", host, port);
});
