package conf

// Host ...
var Host = "localhost"

// DefaultPort uses if custom port isn't set
var DefaultPort = ":8080"

// UsersFile uses if custom port isn't set
var UsersFile = "resources/users.json"

// DefaultStorage uses if custom port isn't set
var DefaultStorage = "cache"


/////////////////////////// DATABASE \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

// DatabaseName store name of your database
// If database does not exist, it will be created one.
var DatabaseName = "quests"

// DatabaseUser - user name to connect sql server
var DatabaseUser = "demian"

// DatabasePsw - password to connect sql server
var DatabasePsw = "VeryHardPassword123456"

// DatabasePort ...
var DatabasePort = ":3306"