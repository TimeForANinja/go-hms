-- no dependencies
CREATE TABLE IF NOT EXISTS `categories` (
  -- int: unique identifier for a categorie
  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- text: name of the categorie
  `name`            TEXT    NOT NULL,
  -- int: id of a parent categorie
  `parentCategorie` INTEGER,
  -- link categories parent to categorie.id, if parent gets deleted or changed this will be as well
  FOREIGN KEY(`parentCategorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `users` (
  -- int: unique identifier for a user
  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- text: the users name
  `name`            TEXT    NOT NULL,
  -- text: the users password hash
  `pwHash`          TEXT    NOT NULL,
  -- int: permission level (can only edit users/groups with lower permissionLevel)
  `permissionLevel` INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS `groups` (
  -- int: unique identifier for a group
  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- text: name of the group
  `name`            TEXT    NOT NULL,
  -- int: permission level (can only edit users/groups with lower permissionLevel)
  `permissionLevel` INTEGER NOT NULL
);
-- medium dependencies
CREATE TABLE IF NOT EXISTS `videos` (
  -- int: unique identifier for a video
  `id`              INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- text: the title of the video
  `title`           TEXT    NOT NULL,
  -- text: the link the video was downloaded from
  `downloadRef`     TEXT    NOT NULL,
  -- int: id of another video thats equal
  `equalVideo`      INTEGER,
  -- int: how often the video was viewed
  `viewed`          INTEGER NOT NULL DEFAULT 0,
  -- text: location of the local video file (if exists)
  `localFile`       TEXT,
  -- bool: block this video from being downloaded
  `blockDownload`   INTEGER NOT NULL CHECK (blockDownload IN (0,1)) DEFAULT 0,
  -- int: id of the user who issued blockDownload (null if the user doesn't exist anymore)
  `blockDownloadBy` INTEGER,
  -- int: timestamp the video was added
  `added`           INTEGER NOT NULL,
  -- int: id of the user who added the video (null if the user doesn't exist anymore)
  `addedBy`         INTEGER,
  -- link equalVideo to videos.id, if entry in videos gets deleted set to null
  -- TODO: if the "parent" gets deleted and the parent has multiple children they are still the same video
  FOREIGN KEY(`equalVideo`)      REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE SET NULL,
  -- link blockDownloadBy to users.id, if entry in users gets deleted set to null
  -- TODO: may be able to delete video
  FOREIGN KEY(`blockDownloadBy`) REFERENCES `users`(`id`)  ON UPDATE CASCADE ON DELETE SET NULL,
  -- link addedBy to users.id, if entry in users gets deleted set to null
  FOREIGN KEY(`addedBy`)         REFERENCES `users`(`id`)  ON UPDATE CASCADE ON DELETE SET NULL
);
-- lot of dependencies
CREATE TABLE IF NOT EXISTS `downloadQ` (
  -- int: unique identifier for a q item
  `id`     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: ref to a video including a downloadRef
  `video`  INTEGER NOT NULL,
  -- bool: is a worker already downloading?
  `inwork` INTEGER NOT NULL CHECK (inwork IN (0,1)) DEFAULT 0,
  -- bool: did a worker fail to download?
  `failed` INTEGER NOT NULL CHECK (failed IN (0,1)) DEFAULT 0,
  -- link video to videos.id, if entry in videos gets deleted or changed this will be as well
  FOREIGN KEY(`video`) REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `comments` (
  -- int: unique identifier for a comment
  `id`            INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of the user who gave the comment
  `user`          INTEGER NOT NULL,
  -- int: id of the video the comment was given on
  `video`         INTEGER NOT NULL,
  -- int: id of a parent comment this comment refers to
  `parentComment` INTEGER,
  -- text: the comment itself
  `comment`       TEXT    NOT NULL,
  -- int: timestamp the comment was added
  `added`         INTEGER NOT NULL,
  -- link user to users.id, if entry in users gets deleted or changed this will be as well
  FOREIGN KEY(`user`)          REFERENCES `users`(`id`)    ON UPDATE CASCADE ON DELETE CASCADE,
  -- link video to videos.id, if entry in videos gets deleted or changed this will be as well
  FOREIGN KEY(`video`)         REFERENCES `videos`(`id`)   ON UPDATE CASCADE ON DELETE CASCADE,
  -- link parentComment to comments.id, if entry in comments gets deleted or changed this will be as well
  FOREIGN KEY(`parentComment`) REFERENCES `comments`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `ratings` (
  -- int: unique identifier for a rating
  `id`     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of the user who gave the rating
  `user`   INTEGER NOT NULL,
  -- int: id of the video the rating was given on
  `video`  INTEGER NOT NULL,
  -- int: the rating itself, n element N and 0<=n<=5
  `rating` INTEGER NOT NULL CHECK (rating IN (1,2,3,4,5)),
  -- int: timestamp the rating was given
  `added`  INTEGER NOT NULL,
  -- link user to users.id, if entry in users gets deleted or changed this will be as well
  FOREIGN KEY(`user`)          REFERENCES `users`(`id`)    ON UPDATE CASCADE ON DELETE CASCADE,
  -- link video to videos.id, if entry in videos gets deleted or changed this will be as well
  FOREIGN KEY(`video`)         REFERENCES `videos`(`id`)   ON UPDATE CASCADE ON DELETE CASCADE
);
-- basicly only dependencies
CREATE TABLE IF NOT EXISTS `user_group_links` (
  -- int: unique identifier for a user_group_link
  `id`    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of a user
  `user`  INTEGER NOT NULL,
  -- int: id of a group
  `group` INTEGER NOT NULL,
  -- link user to users.id, if entry in users gets deleted or changed this will be as well
  FOREIGN KEY(`user`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  -- link group to groups.id, if entry in groups gets deleted or changed this will be as well
  FOREIGN KEY(`group`) REFERENCES `groups`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `video_categorie_links` (
  -- int: unique identifier for a video_categorie_link
  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of a video
  `video`     INTEGER NOT NULL,
  -- int: id of a categorie
  `categorie` INTEGER NOT NULL,
  -- link video to videos.id, if entry in videos gets deleted or changed this will be as well
  FOREIGN KEY(`video`) REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  -- link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `categorie_perm_groups` (
  -- int: unique identifier for a categorie_readable_group
  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of a user
  `group`     INTEGER NOT NULL,
  -- int: id of a group
  `categorie` INTEGER NOT NULL,
  -- bool: whether a group can read/see this categorie
  `canRead` INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0,
  -- bool: whether a group can write/edit/assign this categorie
  `canWrite` INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0,
  -- link group to groups.id, if entry in groups gets deleted or changed this will be as well
  FOREIGN KEY(`group`) REFERENCES `groups`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  -- link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS `categorie_perm_users` (
  -- int: unique identifier for a categorie_readable_user
  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  -- int: id of a user
  `user`     INTEGER NOT NULL,
  -- int: id of a group
  `categorie` INTEGER NOT NULL,
  -- bool: whether a user can read/see this categorie
  `canRead` INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0,
  -- bool: whether a user can write/edit/assign this categorie
  `canWrite` INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0,
  -- link user to users.id, if entry in users gets deleted or changed this will be as well
  FOREIGN KEY(`user`) REFERENCES `user`(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  -- link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
