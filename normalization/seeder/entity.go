package seeder

const (
	numOfInserteddata = 1000000
	batchInsertSize   = 1000
)

type (
	UserNormalize struct {
		ID   uint
		Name string
	}

	CourseNormalize struct {
		ID           uint
		Slug         string
		Image        string
		Name         string
		Price        int
		Stock        int
		InstructorID uint
	}

	InstructorNormalize struct {
		ID   uint
		Name string
	}

	TransactionNormalize struct {
		ID         uint
		Price      int
		Qty        int
		CourseName string
		UserID     uint
		CourseID   uint
	}

	TransactionNormalizeDenormalize struct {
		ID             uint
		Price          int
		Qty            int
		CourseName     string
		UserName       string
		InstructorName string
		CourseImage    string
		UserID         uint
		CourseID       uint
	}

	TransactionDenormalize struct {
		ID             uint
		Price          int
		Qty            int
		CourseSlug     string
		CourseImage    string
		CourseName     string
		UserName       string
		InstructorName string
	}
)

func (UserNormalize) TableName() string {
	return "user_normalize"
}

func (CourseNormalize) TableName() string {
	return "course_normalize"
}

func (InstructorNormalize) TableName() string {
	return "instructor_normalize"
}

func (TransactionNormalize) TableName() string {
	return "transaction_normalize"
}

func (TransactionNormalizeDenormalize) TableName() string {
	return "transaction_normalize_denormalize"
}

func (TransactionDenormalize) TableName() string {
	return "transaction_denormalize"
}
