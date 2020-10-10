package main

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/test/gin-mysql-redis/controler"
	"github.com/hongminhcbg/test/gin-mysql-redis/daos"
	"github.com/hongminhcbg/test/gin-mysql-redis/redis"
)

func InitGin(dao daos.PersonDao, redis redis.RedisPerson) *gin.Engine {
	controler := controler.Controler{Dao: dao, Redis: redis}
	engine := gin.Default()
	engine.GET("/health-check", controler.HealthCheck)
	engine.POST("/person", controler.CreatePerson)
	engine.GET("/person", controler.GetPerson)
	return engine
}
