package sqlQueries

// QueryCreateConfig is the query to create a config table
const QueryCreateConfig = "CREATE TABLE IF NOT EXISTS \"config\" (" +
	// int: unique identifier for a config entry
	"  \"id\"    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the property
	"  \"key\"   TEXT    NOT NULL UNIQUE," +
	// int: value of the property
	"  \"value\" TEXT    NOT NULL UNIQUE" +
	");"
