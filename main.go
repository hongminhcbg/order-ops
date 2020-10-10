package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/hongminhcbg/test/gin-mysql-redis/daos"
	my_redis "github.com/hongminhcbg/test/gin-mysql-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dbURL = `lhm:1@tcp(localhost:3306)/testdb?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local`

func main() {
	db, err := gorm.Open("mysql", dbURL)
	defer db.Close()
	if err != nil {
		panic("open db error: " + err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		panic("ping db error: " + err.Error())
	}

	personDao := daos.NewPersonDao(db)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RedisCache := my_redis.NewRedisPerson(client)
	engine := InitGin(personDao, RedisCache)
	engine.Run(":8080")
}
