package sqlUtil

// QueryCreateCategories is the query to create a categories table
const QueryCreateCategories = "CREATE TABLE IF NOT EXISTS `categories` (" +
	// int: unique identifier for a categorie" +
	"  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the categorie" +
	"  `name`            TEXT    NOT NULL," +
	// int: id of a parent categorie" +
	"  `parentCategorie` INTEGER," +
	// link categories parent to categorie.id, if parent gets deleted or changed this will be as well" +
	"  FOREIGN KEY(`parentCategorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateUsers is the query to create a users table
const QueryCreateUsers = "CREATE TABLE IF NOT EXISTS `users` (" +
	// int: unique identifier for a user
	"  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the users name
	"  `name`            TEXT    NOT NULL," +
	// text: the users password hash
	"  `pwHash`          TEXT    NOT NULL," +
	// int: permission level (can only edit users/groups with lower permissionLevel)
	"  `permissionLevel` INTEGER NOT NULL" +
	");"

// QueryCreateGroups is the query to create a groups table
const QueryCreateGroups = "CREATE TABLE IF NOT EXISTS `groups` (" +
	// int: unique identifier for a group
	"  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the group
	"  `name`            TEXT    NOT NULL," +
	// int: permission level (can only edit users/groups with lower permissionLevel)
	"  `permissionLevel` INTEGER NOT NULL" +
	");"
