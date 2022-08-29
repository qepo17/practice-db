package seeder

import (
	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

func SeederUserNormalize(db *gorm.DB) {
	var usersNormalize []UserNormalize

	for i := 0; i < numOfInserteddata; i++ {
		user := UserNormalize{
			Name: faker.Name().Name(),
		}

		usersNormalize = append(usersNormalize, user)
	}

	db.CreateInBatches(usersNormalize, batchInsertSize)
}
func SeederInstructorNormalize(db *gorm.DB) {
	var instructorsNormalize []InstructorNormalize

	for i := 0; i < numOfInserteddata; i++ {
		instructor := InstructorNormalize{
			Name: faker.Name().Name(),
		}

		instructorsNormalize = append(instructorsNormalize, instructor)
	}

	db.CreateInBatches(instructorsNormalize, batchInsertSize)
}
func SeederCourseNormalize(db *gorm.DB) {
	var coursesNormalize []CourseNormalize

	for i := 0; i < numOfInserteddata; i++ {
		course := CourseNormalize{
			Slug:         faker.Internet().Slug(),
			Name:         faker.Commerce().ProductName(),
			Image:        faker.Avatar().Url("jpg", 640, 300),
			Price:        faker.Number().NumberInt(5),
			Stock:        faker.Number().NumberInt(1),
			InstructorID: uint(faker.Number().NumberInt(5)),
		}

		coursesNormalize = append(coursesNormalize, course)
	}

	db.CreateInBatches(coursesNormalize, batchInsertSize)
}
func SeederTransactionNormalize(db *gorm.DB) {
	var transactionsNormalize []TransactionNormalize

	for i := 0; i < numOfInserteddata; i++ {
		transaction := TransactionNormalize{
			Price:      faker.Number().NumberInt(5),
			Qty:        faker.Number().NumberInt(5),
			CourseName: faker.Commerce().ProductName(),
			UserID:     uint(faker.Number().NumberInt(5)),
			CourseID:   uint(faker.Number().NumberInt(5)),
		}

		transactionsNormalize = append(transactionsNormalize, transaction)
	}

	db.CreateInBatches(transactionsNormalize, batchInsertSize)
}
func SeederTransactionNormalizeDenormalize(db *gorm.DB) {
	var transactionsNormalizeDenormalize []TransactionNormalizeDenormalize

	for i := 0; i < numOfInserteddata; i++ {
		transaction := TransactionNormalizeDenormalize{
			Price:          faker.Number().NumberInt(5),
			Qty:            faker.Number().NumberInt(5),
			CourseName:     faker.Commerce().ProductName(),
			CourseImage:    faker.Avatar().Url("jpg", 640, 300),
			InstructorName: faker.Name().Name(),
			UserName:       faker.Name().Name(),
			UserID:         uint(faker.Number().NumberInt(5)),
			CourseID:       uint(faker.Number().NumberInt(5)),
		}

		transactionsNormalizeDenormalize = append(transactionsNormalizeDenormalize, transaction)
	}

	db.CreateInBatches(transactionsNormalizeDenormalize, batchInsertSize)
}
func SeederTransactionDenormalize(db *gorm.DB) {
	var transactionsDenormalize []TransactionDenormalize

	for i := 0; i < numOfInserteddata; i++ {
		transaction := TransactionDenormalize{
			Price:          faker.Number().NumberInt(5),
			Qty:            faker.Number().NumberInt(5),
			CourseSlug:     faker.Internet().Slug(),
			CourseName:     faker.Commerce().ProductName(),
			CourseImage:    faker.Avatar().Url("jpg", 640, 300),
			InstructorName: faker.Name().Name(),
			UserName:       faker.Name().Name(),
		}

		transactionsDenormalize = append(transactionsDenormalize, transaction)
	}

	db.CreateInBatches(transactionsDenormalize, batchInsertSize)
}
