package sqlQueries

// QueryCreateMovies is the query to create a movies table
const QueryCreateMovies = "CREATE TABLE IF NOT EXISTS \"movies\" (" +
	// int: unique identifier for a video
	"  \"id\"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the title of the video
	"  \"title\"           TEXT    NOT NULL," +
	// text: alternative titles formatted as [[<identifier>, <title>], ...]
	"  \"alt_titles\" TEXT," +
	// text: the mobie's imdb rating
	"  \"imdbRef\"         TEXT    NOT NULL UNIQUE," +
	// real: the movie's imdb rating
	"  \"imdbRating\"      REAL    NOT NULL," +
	// int: the movie's trailer in video
	"  \"trailer\"         INT," +
	// text: the movie's cover photo file location
	"  \"coverFile\"       TEXT," +
	// text: location of the local video file (if exists)
	"  \"localFile\"       TEXT             UNIQUE," +
	// bool: block this video from being downloaded
	"  \"blockDownload\"   INTEGER NOT NULL CHECK (blockDownload IN (0,1)) DEFAULT 0," +
	// int: id of the user who issued blockDownload (null if the user doesn't exist anymore)
	// TODO: if a video is blocked and another user tries to download it that might give away that it was blocked
	"  \"blockDownloadBy\" INTEGER," +
	// int: timestamp the video was added
	"  \"added\"           INTEGER NOT NULL," +
	// int: id of the user who added the video (null if the user doesn't exist anymore)
	"  \"addedBy\"         INTEGER," +
	// int: the length of the video in milliseconds
	"  \"length\"          INTEGER NOT NULL," +
	// text: the has of the downloaded file (if any)
	"  \"fileHash\"        TEXT             UNIQUE," +
	// int: the file size in bytes (if any)
	"  \"fileSize\"        INTEGER," +
	// link trailer to videos.id, if entry in videos gets deleted set null
	"  FOREIGN KEY(\"trailer\")         REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE SET NULL," +
	// link blockDownloadBy to users.id, if entry in users gets deleted set to null
	// TODO: may be able to delete video if theres also no one blocking it anymore?
	"  FOREIGN KEY(\"blockDownloadBy\") REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE SET NULL," +
	// link addedBy to users.id, if entry in users gets deleted set to null
	// TODO: rly want to delete the video if the user who downloaded it gets removed?
	"  FOREIGN KEY(\"addedBy\")         REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE SET NULL" +
	");"
