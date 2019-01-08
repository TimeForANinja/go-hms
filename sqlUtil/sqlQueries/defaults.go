package sqlQueries

// add the initial root user with pw "root"
const QueryDefaultAdminUser = "INSERT INTO \"users\" (\"name\", \"pwHash\", \"permissionLevel\", \"created\", \"creator\") VALUES('root', '$2a$10$vyRPiPsOLdY7xOW4ctEa1emX16B6G1f7nOEZomJkdytWOwVLCS.UC', 9999, 0, 1)"

// add @everyone group
const QueryDefaultEveryoneGroup = "INSERT INTO \"groups\" (\"name\", \"permissionLevel\", \"created\", \"creator\") VALUES('@everyone', 0, 0, 1)"

// add admin to @everyone group
const QueryDefaultAdminUserGroups = "INSERT INTO \"link_user_group\" (\"user\", \"group\") VALUES(1, 1)"

// save @everyone group as default group
const QueryDefaultGroup = "INSERT INTO \"config\" (\"key\", \"value\") VALUES('DEFAULT_GROUP', '1')"
