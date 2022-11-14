package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Foo struct {
	gorm.Model
	Bars []Bar
}

type Bar struct {
	gorm.Model
	FooID uint
}

func main() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db.AutoMigrate(&Foo{})
	db.AutoMigrate(&Bar{})

	record := Foo{
		Bars: []Bar{{FooID: 1}, {}},
	}

	db.
		Where("1 = 1").
		Updates(&record)
}
