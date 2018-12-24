package sqlUtil

// QueryCreateDownloadQ is the query to create a downloadQ table
const QueryCreateDownloadQ = "CREATE TABLE IF NOT EXISTS `downloadQ` (" +
	// int: unique identifier for a q item
	"  `id`     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: ref to a video including a downloadRef
	"  `video`  INTEGER NOT NULL," +
	// bool: is a worker already downloading?
	"  `inwork` INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0," +
	// bool: did a worker fail to download?
	"  `failed` INTEGER NOT NULL CHECK (failed IN (0,1)) DEFAULT 0," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(`video`) REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateComments is the query to create a comments table
const QueryCreateComments = "CREATE TABLE IF NOT EXISTS `comments` (" +
	// int: unique identifier for a comment
	"  `id`            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of the user who gave the comment
	"  `user`          INTEGER NOT NULL," +
	// int: id of the video the comment was given on
	"  `video`         INTEGER NOT NULL," +
	// int: id of a parent comment this comment refers to
	"  `parentComment` INTEGER," +
	// text: the comment itself
	"  `comment`       TEXT    NOT NULL," +
	// int: timestamp the comment was added
	"  `added`         INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(`user`)          REFERENCES `users`(`id`)    ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(`video`)         REFERENCES `videos`(`id`)   ON UPDATE CASCADE ON DELETE CASCADE," +
	// link parentComment to comments.id, if entry in comments gets deleted or changed this will be as well
	"  FOREIGN KEY(`parentComment`) REFERENCES `comments`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateRatings is the query to create a ratings table
const QueryCreateRatings = "CREATE TABLE IF NOT EXISTS `ratings` (" +
	// int: unique identifier for a rating
	"  `id`     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of the user who gave the rating
	"  `user`   INTEGER NOT NULL," +
	// int: id of the video the rating was given on
	"  `video`  INTEGER NOT NULL," +
	// int: the rating itself, n element N and 0<=n<=5
	"  `rating` INTEGER NOT NULL CHECK (rating IN (1,2,3,4,5))," +
	// int: timestamp the rating was given
	"  `added`  INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(`user`)          REFERENCES `users`(`id`)    ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(`video`)         REFERENCES `videos`(`id`)   ON UPDATE CASCADE ON DELETE CASCADE" +
	");"
