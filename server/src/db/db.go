package db

import (
	"fmt"
	"log"
	"time"
	"todo/src/config"
	"todo/src/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var lastConnection util.Optional[util.Result[*sqlx.DB]] = util.None[util.Result[*sqlx.DB]]()

func ConnectDB() util.Result[*sqlx.DB] {
	if lastConnection.Has {
		return lastConnection.Value
	}
	user := config.GetDBUser()
	password := config.GetDBPassword()
	address := config.GetDBAddress()
	port := config.GetDBPort()
	name := config.GetDBName()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=UTC", user, password, address, port, name)

	var err error
	for i := 1; i <= 10; i += 1 {
		log.Printf("the %d times connection trying", i)
		db := util.NewResult(sqlx.Connect("mysql", dsn))
		if db.Success {
			lastConnection = util.Some(db)
			return db
		}
		log.Printf("the %d times connection failed", i)
		time.Sleep(100 * time.Millisecond)
		err = db.Error
	}
	res := util.NewResultFailed[*sqlx.DB](err)
	lastConnection = util.Some(res)
	return res
}
