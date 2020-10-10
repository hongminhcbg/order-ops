package daos

import (
	"github.com/hongminhcbg/test/gin-mysql-redis/models"
	"github.com/jinzhu/gorm"
)

type PersonDao interface {
	Create(person *models.Person) error
	Save(person *models.Person) error
	Read(id int) (*models.Person, error)
}

type personDaoImpl struct {
	db *gorm.DB
}
func NewPersonDao(db *gorm.DB) PersonDao {
	return &personDaoImpl{
		db:db,
	}
}

func (c *personDaoImpl)Create(person *models.Person) error{
	return c.db.Create(person).Error
}

func (c *personDaoImpl)Save(person *models.Person) error{
	return c.db.Save(person).Error
}
func (c *personDaoImpl)Read(id int) (*models.Person, error)  {
	var person models.Person
	if err := c.db.Where("id = ?", id).First(&person).Error; err != nil {
		return nil, err
	}
	return &person, nil
}