package sqlQueries

// QueryCreateCategories is the query to create a categories table
const QueryCreateCategories = "CREATE TABLE IF NOT EXISTS \"categories\" (" +
	// int: unique identifier for a categorie
	"  \"id\"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the categorie
	"  \"name\"            TEXT    NOT NULL UNIQUE," +
	// int: id of a parent categorie or null
	"  \"parentCategorie\" INTEGER," +
	// link categories parent to categorie.id, if parent gets deleted or changed this will be as well
	"  FOREIGN KEY(\"parentCategorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateLinkCategorieUserPermission is the query to create a link_categorie_user_permission table
const QueryCreateLinkCategorieUserPermission = "CREATE TABLE IF NOT EXISTS \"link_categorie_user_permission\" (" +
	// int: unique identifier for a link_categorie_user_permission
	"  \"id\"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a user
	"  \"user\"      INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\" INTEGER NOT NULL," +
	// bool: whether a user can read/see this categorie
	"  \"canRead\"   INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0," +
	// bool: whether a user can write/edit/assign this categorie
	"  \"canWrite\"  INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"user\")      REFERENCES \"users\"(\"id\")      ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// user + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"user\", \"categorie\")" +
	");"

// QueryCreateLinkCategoriePersonality is the query to create a link_categorie_personality table
const QueryCreateLinkCategoriePersonality = "CREATE TABLE IF NOT EXISTS \"link_categorie_personality\" (" +
	// int: unique identifier for a link_categorie_personality
	"  \"id\"          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a personality
	"  \"personality\" INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\"   INTEGER NOT NULL," +
	// link personality to personalities.id, if entry in personalities gets deleted or changed this will be as well
	"  FOREIGN KEY(\"personality\") REFERENCES \"personalities\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\")   REFERENCES \"categories\"(\"id\")    ON UPDATE CASCADE ON DELETE CASCADE," +
	// personality + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"personality\", \"categorie\")" +
	");"

// QueryCreateLinkCategorieVideo is the query to create a link_categorie_video table
const QueryCreateLinkCategorieVideo = "CREATE TABLE IF NOT EXISTS \"link_categorie_video\" (" +
	// int: unique identifier for a link_categorie_video
	"  \"id\"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a video
	"  \"video\"     INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\" INTEGER NOT NULL," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\")     REFERENCES \"videos\"(\"id\")     ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// video + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"video\", \"categorie\")" +
	");"

// QueryCreateLinkCategorieMovie is the query to create a link_categorie_movie table
const QueryCreateLinkCategorieMovie = "CREATE TABLE IF NOT EXISTS \"link_categorie_movie\" (" +
	// int: unique identifier for a link_categorie_movie
	"  \"id\"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a movie
	"  \"movie\"     INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\" INTEGER NOT NULL," +
	// link movie to movies.id, if entry in movies gets deleted or changed this will be as well
	"  FOREIGN KEY(\"movie\")     REFERENCES \"movies\"(\"id\")     ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// video + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"movie\", \"categorie\")" +
	");"

// QueryCreateLinkCategorieGroupPermission is the query to create a link_categorie_group_permission table
const QueryCreateLinkCategorieGroupPermission = "CREATE TABLE IF NOT EXISTS \"link_categorie_group_permission\" (" +
	// int: unique identifier for a link_categorie_group_permission
	"  \"id\"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a group
	"  \"group\"     INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\" INTEGER NOT NULL," +
	// bool: whether a group can read/see this categorie
	"  \"canRead\"   INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0," +
	// bool: whether a group can write/edit/assign this categorie
	"  \"canWrite\"  INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0," +
	// link group to groups.id, if entry in groups gets deleted or changed this will be as well
	"  FOREIGN KEY(\"group\")     REFERENCES \"groups\"(\"id\")     ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// group + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"group\", \"categorie\")" +
	");"

// QueryCreateLinkCategoriePlaylist is the query to create a link_categorie_playlist table
const QueryCreateLinkCategoriePlaylist = "CREATE TABLE IF NOT EXISTS \"link_categorie_playlist\" (" +
	// int: unique identifier for a link_categorie_playlist
	"  \"id\"        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a playlist
	"  \"playlist\"  INTEGER NOT NULL," +
	// int: id of a categorie
	"  \"categorie\" INTEGER NOT NULL," +
	// link playlist to playlists.id, if entry in playlists gets deleted or changed this will be as well
	"  FOREIGN KEY(\"playlist\")  REFERENCES \"playlists\"(\"id\")  ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets gets deleted or changed this will be as well
	"  FOREIGN KEY(\"categorie\") REFERENCES \"categories\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// playlist + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"playlist\", \"categorie\")" +
	");"
