package models

type Person struct {
	ID        int64  `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string `gorm:"unique;column:name" json:"name"`
	Age       int64  `gorm:"column:age;type:int(10)" json:"age"`
	Address   string `gorm:"column:address;type:varchar(128)" json:"address"`
	IsMarried bool   `gorm:"column:is_married;type:tinyint(1)" json:"is_married"`
}

func (Person) TableName() string {
	return "persons"
}
