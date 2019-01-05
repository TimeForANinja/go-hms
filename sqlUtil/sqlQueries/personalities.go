package sqlQueries

// QueryCreatePersonalities is the query to create a personalities table
const QueryCreatePersonalities = "CREATE TABLE IF NOT EXISTS \"personalities\" (" +
	// int: unique identifier for a personality
	"  \"id\"      INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the name of a personality
	"  \"name\"    TEXT    NOT NULL UNIQUE," +
	// text: the link to the personality's imdb page if it exists
	"  \"imdbRef\" TEXT             UNIQUE," +
	// int: unix timestamp the user was created
	"  \"created\" INTEGER NOT NULL," +
	// int: the id of a user, who created this user, or null if the creator was removed
	"  \"creator\" INTEGER," +
	// link creator to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"creator\") REFERENCES \"users\"(\"id\") ON UPDATE CASCADE ON DELETE SET NULL" +
	");"

// QueryCreateLinkPersonalityVideo is the query to create a link_personality_video table
const QueryCreateLinkPersonalityVideo = "CREATE TABLE IF NOT EXISTS \"link_personality_video\" (" +
	// int: unique identifier for a link_personality_video
	"  \"id\"          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a personality
	"  \"personality\" INTEGER NOT NULL," +
	// int: id of a video
	"  \"video\"       INTEGER NOT NULL," +
	// link personality to personalities.id, if entry in personalities gets deleted or changed this will be as well
	"  FOREIGN KEY(\"personality\") REFERENCES \"personalities\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\")       REFERENCES \"videos\"(\"id\")        ON UPDATE CASCADE ON DELETE CASCADE" +
	// personality + video combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"personality\", \"video\")" +
	");"

// QueryCreateLinkPersonalityMovie is the query to create a link_personality_movie table
const QueryCreateLinkPersonalityMovie = "CREATE TABLE IF NOT EXISTS \"link_personality_movie\" (" +
	// int: unique identifier for a link_personality_movie
	"  \"id\"          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a personality
	"  \"personality\" INTEGER NOT NULL," +
	// int: id of a movie
	"  \"movie\"       INTEGER NOT NULL," +
	// link personality to personalities.id, if entry in personalities gets deleted or changed this will be as well
	"  FOREIGN KEY(\"personality\") REFERENCES \"personalities\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// link movie to movies.id, if entry in movies gets deleted or changed this will be as well
	"  FOREIGN KEY(\"movie\")       REFERENCES \"movies\"(\"id\")        ON UPDATE CASCADE ON DELETE CASCADE" +
	// personality + movie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"personality\", \"movie\")" +
	");"
