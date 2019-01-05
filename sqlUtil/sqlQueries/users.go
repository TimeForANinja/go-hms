package sqlQueries

// QueryCreateUsers is the query to create a users table
const QueryCreateUsers = "CREATE TABLE IF NOT EXISTS \"users\" (" +
	// int: unique identifier for a user
	"  \"id\"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the users name
	"  \"name\"            TEXT    NOT NULL UNIQUE," +
	// text: the users password hash
	"  \"pwHash\"          TEXT    NOT NULL," +
	// int: permission level (can only edit users/groups with lower permissionLevel)
	"  \"permissionLevel\" INTEGER NOT NULL," +
	// int: unix timestamp the user was created
	"  \"created\"         INTEGER NOT NULL," +
	// int: the id of a user, who created this user, or null if the creator was removed
	"  \"creator\"         INTEGER," +
	// link creator to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"creator\") REFERENCES \"users\"(\"id\") ON UPDATE CASCADE ON DELETE SET NULL" +
	");"

// QueryCreateLinkUserGroup is the query to create a link_user_group table
const QueryCreateLinkUserGroup = "CREATE TABLE IF NOT EXISTS \"link_user_group\" (" +
	// int: unique identifier for a link_user_group
	"  \"id\"    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a user
	"  \"user\"  INTEGER NOT NULL," +
	// int: id of a group
	"  \"group\" INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"user\")  REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE CASCADE," +
	// link group to groups.id, if entry in groups gets deleted or changed this will be as well
	"  FOREIGN KEY(\"group\") REFERENCES \"groups\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	// user + categorie combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"user\", \"group\")" +
	");"

// QueryCreateComments is the query to create a comments table
const QueryCreateComments = "CREATE TABLE IF NOT EXISTS \"comments\" (" +
	// int: unique identifier for a comment
	"  \"id\"            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of the user who gave the comment
	"  \"user\"          INTEGER NOT NULL," +
	// int: id of the video the comment was given on
	"  \"video\"         INTEGER NOT NULL," +
	// int: id of a parent comment this comment refers to
	"  \"parentComment\" INTEGER," +
	// text: the comment itself
	"  \"comment\"       TEXT    NOT NULL," +
	// int: timestamp the comment was added
	"  \"added\"         INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"user\")  REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\") REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// link parentComment to comments.id, if entry in comments gets deleted or changed this will be as well
	// TODO: might want to change this so that child comments keep existing when parent gets deleted
	"  FOREIGN KEY(\"parentComment\") REFERENCES \"comments\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateRatings is the query to create a ratings table
const QueryCreateRatings = "CREATE TABLE IF NOT EXISTS \"ratings\" (" +
	// int: unique identifier for a rating
	"  \"id\"     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of the user who gave the rating
	"  \"user\"   INTEGER NOT NULL," +
	// int: id of the video the rating was given on
	"  \"video\"  INTEGER NOT NULL," +
	// int: the rating itself, n element N and 0<=n<=10
	"  \"rating\" INTEGER NOT NULL CHECK (rating IN (1,2,3,4,5,6,7,8,9,10))," +
	// int: timestamp the rating was given
	"  \"added\"  INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"user\")  REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\") REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	");"
