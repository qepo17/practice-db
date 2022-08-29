package main

import (
	"normalization/seeder"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	println("start seeder")
	seeder.SeederUserNormalize(db)
	seeder.SeederInstructorNormalize(db)
	seeder.SeederCourseNormalize(db)
	seeder.SeederTransactionNormalize(db)
	seeder.SeederTransactionNormalizeDenormalize(db)
	seeder.SeederTransactionDenormalize(db)
	println("end seeder")
}

func initDB() *gorm.DB {
	dsn := "train:train@tcp(103.226.138.88:3306)/normalization?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
