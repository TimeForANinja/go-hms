package sqlQueries

// QueryCreateDownloadQ is the query to create a downloadQ table
const QueryCreateDownloadQ = "CREATE TABLE IF NOT EXISTS \"downloadQ\" (" +
	// int: unique identifier for a q item
	"  \"id\"     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: ref to a video including a downloadRef
	"  \"video\"  INTEGER NOT NULL UNIQUE," +
	// bool: is a worker already downloading?
	"  \"inwork\" INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0," +
	// bool: did a worker fail to download?
	"  \"failed\" INTEGER NOT NULL CHECK (failed IN (0,1)) DEFAULT 0," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\") REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE" +
	");"
