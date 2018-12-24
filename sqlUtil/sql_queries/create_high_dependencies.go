package sqlUtil

// QueryCreateLinksUserGroup is the query to create a links_user_group table
const QueryCreateLinksUserGroup = "CREATE TABLE IF NOT EXISTS `links_user_group` (" +
	// int: unique identifier for a user_group_link
	"  `id`    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a user
	"  `user`  INTEGER NOT NULL," +
	// int: id of a group
	"  `group` INTEGER NOT NULL," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(`user`) REFERENCES `users`(`id`) ON UPDATE CASCADE ON DELETE CASCADE," +
	// link group to groups.id, if entry in groups gets deleted or changed this will be as well
	"  FOREIGN KEY(`group`) REFERENCES `groups`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateLinksVideoCategorie is the query to create a links_video_categorie table
const QueryCreateLinksVideoCategorie = "CREATE TABLE IF NOT EXISTS `links_video_categorie` (" +
	// int: unique identifier for a video_categorie_link
	"  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a video
	"  `video`     INTEGER NOT NULL," +
	// int: id of a categorie
	"  `categorie` INTEGER NOT NULL," +
	// link video to videos.id, if entry in videos gets deleted or changed this will be as well
	"  FOREIGN KEY(`video`) REFERENCES `videos`(`id`) ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateLinksCategorieGroupPermission is the query to create a link_categorie_group_permission table
const QueryCreateLinksCategorieGroupPermission = "CREATE TABLE IF NOT EXISTS `link_categorie_group_permission` (" +
	// int: unique identifier for a categorie_readable_group
	"  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a user
	"  `group`     INTEGER NOT NULL," +
	// int: id of a group
	"  `categorie` INTEGER NOT NULL," +
	// bool: whether a group can read/see this categorie
	"  `canRead` INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0," +
	// bool: whether a group can write/edit/assign this categorie
	"  `canWrite` INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0," +
	// link group to groups.id, if entry in groups gets deleted or changed this will be as well
	"  FOREIGN KEY(`group`) REFERENCES `groups`(`id`) ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"

// QueryCreateLinksCategorieUserPermission is the query to create a link_categorie_user_permission table
const QueryCreateLinksCategorieUserPermission = "CREATE TABLE IF NOT EXISTS `link_categorie_user_permission` (" +
	// int: unique identifier for a categorie_readable_user
	"  `id`        INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
	// int: id of a user
	"  `user`     INTEGER NOT NULL," +
	// int: id of a group
	"  `categorie` INTEGER NOT NULL," +
	// bool: whether a user can read/see this categorie
	"  `canRead` INTEGER NOT NULL CHECK (canRead IN (0,1)) DEFAULT 0," +
	// bool: whether a user can write/edit/assign this categorie
	"  `canWrite` INTEGER NOT NULL CHECK (canWrite IN (0,1)) DEFAULT 0," +
	// link user to users.id, if entry in users gets deleted or changed this will be as well
	"  FOREIGN KEY(`user`) REFERENCES `user`(`id`) ON UPDATE CASCADE ON DELETE CASCADE," +
	// link categorie to categories.id, if entry in categories gets deleted or changed this will be as well
	"  FOREIGN KEY(`categorie`) REFERENCES `categories`(`id`) ON UPDATE CASCADE ON DELETE CASCADE" +
	");"
