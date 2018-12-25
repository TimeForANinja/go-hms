package sqlQueries

// QueryCreateCategories is the query to create a categories table
const QueryCreateCategories = "CREATE TABLE IF NOT EXISTS `categories` (" +
	// int: unique identifier for a categorie" +
	"  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: name of the categorie" +
	"  `name`            TEXT    NOT NULL UNIQUE," +
	// int: id of a parent categorie" +
	"  `parentCategorie` INTEGER," +
	// link categories parent to categorie.id, if parent gets deleted or changed this will be as well" +
	"  FOREIGN KEY(`parentCategorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateVideos is the query to create a videos table
const QueryCreateVideos = "CREATE TABLE IF NOT EXISTS `videos` (" +
	// int: unique identifier for a video
	"  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the title of the video
	"  `title`           TEXT    NOT NULL," +
	// text: the link the video was downloaded from
	"  `downloadRef`     TEXT    NOT NULL UNIQUE," +
	// int: id of another video thats equal
	"  `equalVideo`      INTEGER," +
	// int: how often the video was viewed
	"  `views`          INTEGER NOT NULL DEFAULT 0," +
	// text: location of the local video file (if exists)
	"  `localFile`       TEXT UNIQUE," +
	// bool: block this video from being downloaded
	"  `blockDownload`   INTEGER NOT NULL CHECK (blockDownload IN (0,1)) DEFAULT 0," +
	// int: id of the user who issued blockDownload (null if the user doesn't exist anymore)
	"  `blockDownloadBy` INTEGER," +
	// int: timestamp the video was added
	"  `added`           INTEGER NOT NULL," +
	// int: id of the user who added the video (null if the user doesn't exist anymore)
	"  `addedBy`         INTEGER," +
	// link equalVideo to videos.id, if entry in videos gets deleted set to null
	// TODO: if the "parent" gets deleted and the parent has multiple children they are still the same video
	"  FOREIGN KEY(`equalVideo`)      REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE SET NULL," +
	// link blockDownloadBy to users.id, if entry in users gets deleted set to null
	// TODO: may be able to delete video
	"  FOREIGN KEY(`blockDownloadBy`) REFERENCES `users`(`id`)  ON UPDATE CASCADE ON DELETE SET NULL," +
	// link addedBy to users.id, if entry in users gets deleted set to null
	"  FOREIGN KEY(`addedBy`)         REFERENCES `users`(`id`)  ON UPDATE CASCADE ON DELETE SET NULL" +
	// TODO: might add length, resolution, filesize, filehash...
	");"
