package main

import (
	"fmt"
	"order-ops/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", config.MySQLConnectURL)
	defer db.Close()
	if err != nil {
		panic("open db error: " + err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		panic("ping db error: " + err.Error())
	}

	fmt.Println("success")
	// personDao := daos.NewPersonDao(db)

	// RedisCache := my_redis.NewRedisPerson(client)
	engine := InitGin(db)
	_ = engine.Run(":8080")
}
