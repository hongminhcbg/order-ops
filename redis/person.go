package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hongminhcbg/test/gin-mysql-redis/models"
	"time"
)

type RedisPerson interface {
	Save(person *models.Person) error
	GetPerson(id int) (*models.Person, error)
}

type redisPersonImpl struct {
	client *redis.Client
}

func NewRedisPerson(client *redis.Client) RedisPerson {
	return &redisPersonImpl{client:client}
}

func (r *redisPersonImpl) Save(person *models.Person) error {
	if person.ID == 0 {
		return errors.New("not valid id")
	}
	bytes, err := json.Marshal(person)
	if err != nil {
		return err
	}
	_, err = r.client.Set(fmt.Sprintf("%v", person.ID), bytes, 1*time.Hour).Result()
	return err
}

func (r *redisPersonImpl) GetPerson(id int) (*models.Person, error) {
	var person models.Person
	val, err := r.client.Get(fmt.Sprintf("%v", id)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
