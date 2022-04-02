package config

import (
	"todo/src/util"
)

func GetDBUser() string {
	return util.OrElse(getenv("user"), "user")
}

func GetDBPassword() string {
	return util.OrElse(getenv("DB_PASSWORD"), "password")
}

func GetDBAddress() string {
	return util.OrElse(getenv("DB_ADDRESS"), "127.0.0.1")
}

func GetDBPort() string {
	return util.OrElse(getenv("DB_PORT"), "3306")
}

func GetDBName() string {
	return util.OrElse(getenv("DB_NAME"), "todo-db")
}
