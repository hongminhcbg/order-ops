package controler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/test/gin-mysql-redis/daos"
	"github.com/hongminhcbg/test/gin-mysql-redis/models"
	"github.com/hongminhcbg/test/gin-mysql-redis/redis"
)

type Controler struct {
	Dao   daos.PersonDao
	Redis redis.RedisPerson
}

func (c Controler) HealthCheck(contex *gin.Context) {
	contex.JSON(200, gin.H{
		"status": "running",
	})
}

func (c Controler) CreatePerson(context *gin.Context) {
	var person models.Person
	context.ShouldBindJSON(&person)
	err := c.Dao.Create(&person)
	if err != nil {
		context.JSON(400, gin.H{
			"message": "create person error",
			"errors":  err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"message": "success",
			"person":  person,
		})
	}
}

func (c Controler) GetPerson(context *gin.Context) {
	id := context.Query("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(400, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}
	if data, err := c.Redis.GetPerson(idint); err == nil {
		context.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
		return
	}

	data, err := c.Dao.Read(idint)
	if err != nil {
		context.JSON(400, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
	} else {
		context.JSON(200, gin.H{
			"status": "success",
			"data":   data,
		})
		err := c.Redis.Save(data)
		if err != nil {
			fmt.Println("save data to cache error")
		}
	}

}
