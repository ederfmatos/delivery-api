package factory

import (
	"delivery-api/application/events"
	"delivery-api/application/repository"
	events2 "delivery-api/infra/events"
	repository2 "delivery-api/infra/repository"
	"delivery-api/infra/repository/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type InfraFactory struct {
	ProductRepository repository.ProductRepository
	EventEmitter      events.EventEmitter
}

func NewInfrastructureFactory() InfraFactory {
	database := GetDatabase()
	return InfraFactory{
		ProductRepository: repository2.NewDefaultProductRepository(database),
		EventEmitter:      events2.NewKafkaEventEmitter("localhost"),
	}
}

func GetDatabase() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(orm.ProductORM{})
	if err != nil {
		panic(err)
	}
	return database
}
