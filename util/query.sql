CREATE TABLE IF NOT EXISTS `downloadQ` (
  -- int: unique identifier for a q item
  `id`     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: ref to a video including a downloadRef
  `video`  INTEGER NOT NULL,
  -- bool: is a worker already downloading?
  `inwork` INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0,
  -- bool: did a worker fail to download?
  `failed` INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0,
  -- link video to videos.id, if entry in videos gets deleted this will be as well
  FOREIGN KEY(video) REFERENCES videos(id) ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS `videos` (
  -- int: unique identifier for a video
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- text: the title of the video
  `title`
  -- text: the link the video was downloaded from
  `downloadRef`
  -- text: comma seperated alternative links
  `altRef` => swap to `sameAs` linking to another video
  -- int: how often the video was viewed
  `viewed`
  -- text: location of the local video file
  `localFile`
  -- bool: block this video from being downloaded
  `blockDownload` INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0,
  -- int: id of the user who issued blockDownload (null if the user doesn't exist anymore)
  `blockDownloadBy` => link to user, `ON UPDATE SET NULL`
  -- text: comma seperated list of categorie id's
  `categories`
  -- int: timestamp the video was added
  `added`
  -- int: id of the user who added the video
  `addedBy`
);
CREATE TABLE IF NOT EXISTS `categories` (
  -- int: unique identifier for a categorie
  `id` INTEGER PRIMARY KEY
  -- int: id of a parent categorie
  `parentCategorie`
  -- text: name of the categorie
  `name`
  -- text: comma seperated list of group id's who can see this categorie
  `readableGroups`
  -- text: comma seperated list of user id's who can see this categorie
  `readableUsers`
);
CREATE TABLE IF NOT EXISTS `comments` (
  -- int: unique identifier for a comment
  `id` INTEGER PRIMARY KEY
  -- int: id of the user who gave the comment
  `user`
  -- int: id of the video the comment was given on
  `video`
  -- int: id of a parent comment this comment refers to
  `parentComment`
  -- text: the comment itself
  `comment`
  -- int: timestamp the comment was added
  `added`
);
CREATE TABLE IF NOT EXISTS `ratings` (
  -- int: unique identifier for a rating
  `id` INTEGER PRIMARY KEY
  -- int: id of the user who gave the rating
  `user`
  -- int: id of the video the rating was given on
  `video`
  -- int: the rating itself, n element N and 0<=n<=5
  `rating`
  -- int: timestamp the rating was given
  `added`
);
CREATE TABLE IF NOT EXISTS `users` (
  -- int: unique identifier for a user
  `id` INTEGER PRIMARY KEY
  -- text: the users name
  `name`
  -- text: the users password hash
  `pwHash`
  -- text: the users password salt
  `salt`
  -- text: comma seperated list of group id's
  `groups`
  -- int: permission level
  `permissionLvl`
);
CREATE TABLE IF NOT EXISTS `groups` (
  -- int: unique identifier for a group
  `id` INTEGER PRIMARY KEY
  -- text: name of the group
  `name`
  -- int: permission level
  `permissionLvl`
);

CREATE TABLE IF NOT EXISTS `user_group_links` ()
CREATE TABLE IF NOT EXISTS `video_categorie_links` ()
