package main

import (
	"normalization/seeder"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	db.AutoMigrate(seeder.UserNormalize{})
	db.AutoMigrate(seeder.CourseNormalize{})
	db.AutoMigrate(seeder.InstructorNormalize{})
	db.AutoMigrate(seeder.TransactionNormalize{})
	db.AutoMigrate(seeder.TransactionNormalizeDenormalize{})
	db.AutoMigrate(seeder.TransactionDenormalize{})

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
	dsn := "host=103.52.114.29 user=metabase password=thepassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug()

	return db
}
