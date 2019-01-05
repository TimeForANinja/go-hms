package sqlQueries

// QueryCreateVideos is the query to create a videos table
const QueryCreateVideos = "CREATE TABLE IF NOT EXISTS \"videos\" (" +
	// int: unique identifier for a video
	"  \"id\"              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// text: the title of the video
	"  \"title\"           TEXT    NOT NULL," +
	// text: the link the video was downloaded from
	"  \"downloadRef\"     TEXT    NOT NULL UNIQUE," +
	// int: how often the video was viewed
	"  \"views\"           INTEGER NOT NULL DEFAULT 0," +
	// text: location of the local video file (if exists)
	"  \"localFile\"       TEXT             UNIQUE," +
	// bool: block this video from being downloaded
	"  \"blockDownload\"   INTEGER NOT NULL CHECK (blockDownload IN (0,1)) DEFAULT 0," +
	// int: id of the user who issued blockDownload (null if the user doesn't exist anymore)
	// TODO: if a video is blocked and another user tries to download it that might give away that it was blocked
	"  \"blockDownloadBy\" INTEGER," +
	// int: timestamp the video was added
	"  \"added\"           INTEGER NOT NULL," +
	// int: id of the user who added the video (null if the user doesn't exist anymore)
	"  \"addedBy\"         INTEGER," +
	// int: the length of the video in milliseconds
	"  \"length\"          INTEGER NOT NULL," +
	// text: the max resolution of the video in <pixel>x<pixel>
	"  \"resolution\"      TEXT    NOT NULL," +
	// text: the has of the downloaded file (if any)
	"  \"fileHash\"        TEXT             UNIQUE," +
	// int: the file size in bytes (if any)
	"  \"fileSize\"        INTEGER," +
	// link blockDownloadBy to users.id, if entry in users gets deleted set to null
	// TODO: may be able to delete video if theres also no one blocking it anymore?
	"  FOREIGN KEY(\"blockDownloadBy\") REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE SET NULL," +
	// link addedBy to users.id, if entry in users gets deleted set to null
	// TODO: rly want to delete the video if the user who downloaded it gets removed?
	"  FOREIGN KEY(\"addedBy\")         REFERENCES \"users\"(\"id\")  ON UPDATE CASCADE ON DELETE SET NULL" +
	");"

const QueryCreateLinkIdenticalVideos = "CREATE TABLE IF NOT EXISTS \"link_identical_videos\" (" +
	// int: unique identifier for a link_identical_videos
	"  \"id\"      INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a video
	"  \"video_a\" INTEGER NOT NULL," +
	// int: id of another video
	"  \"video_b\" INTEGER NOT NULL," +
	// link video_a to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video_a\") REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// link video_b to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video_b\") REFERENCES \"videos\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// video's combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"video_a\", \"video_b\")" +
	");"

const QueryCreateLinkVideoPlaylist = "CREATE TABLE IF NOT EXISTS \"link_video_playlist\" (" +
	// int: unique identifier for a link_video_playlist
	"  \"id\"       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a video
	"  \"video\"    INTEGER NOT NULL," +
	// int: id of a playlist
	"  \"playlist\" INTEGER NOT NULL," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(\"video\")    REFERENCES \"videos\"(\"id\")    ON UPDATE CASCADE ON DELETE CASCADE," +
	// link playlist to playlists.id, if entry in playlists gets deleted or changed this will be as well
	"  FOREIGN KEY(\"playlist\") REFERENCES \"playlists\"(\"id\") ON UPDATE CASCADE ON DELETE CASCADE," +
	// video's combination only allowed once
	"  CONSTRAINT unq UNIQUE (\"video\", \"playlist\")" +
	");"
