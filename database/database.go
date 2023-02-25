package database

var connection string

func init() { // init is constructor, will be executed first time when we import this package / file
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
