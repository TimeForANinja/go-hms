package sqlQueries

// QueryCreatePlaylists is the query to create a playlists table
const QueryCreatePlaylists = "CREATE TABLE IF NOT EXISTS \"playlists\" (" +
	// int: unique identifier for a playlist
	"  \"id\"         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: original name of the movie
	"  \"name\"       TEXT," +
	// text: alternative titles formatted as [[<identifier>, <title>], ...]
	"  \"alt_titles\" TEXT," +
	// bool: whether the playlist is public
	"  \"public\"     INTEGER NOT NULL CHECK (public IN (0,1)) DEFAULT 0," +
	// int: unix timestamp the user was created
	"  \"created\"    INTEGER NOT NULL," +
	// int: the id of a user, who created this user, or null if the creator was removed
	"  \"creator\"    INTEGER," +
	// link creator to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"creator\") REFERENCES \"users\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	");"
