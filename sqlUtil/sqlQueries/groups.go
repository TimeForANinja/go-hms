package sqlQueries

// QueryCreateGroups is the query to create a groups table
const QueryCreateGroups = "CREATE TABLE IF NOT EXISTS \"groups\" (" +
	// int: unique identifier for a group
	"  \"id\"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the group
	"  \"name\"            TEXT    NOT NULL UNIQUE," +
	// int: permission level (can only edit users/groups with lower permissionLevel)
	"  \"permissionLevel\" INTEGER NOT NULL," +
	// int: unix timestamp the user was created
	"  \"created\"         INTEGER NOT NULL," +
	// int: the id of a user, who created this user, or null if the creator was removed
	"  \"creator\"         INTEGER," +
	// link creator to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(\"creator\") REFERENCES \"users\"(\"id\") ON UPDATE CASCADE ON DELETE SET NULL" +
	");"
