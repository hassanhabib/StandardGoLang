package core

type StudentStorageBroker interface {
	InsertStudent(student Student) Student
}

type StudentService interface {
	AddStudent(student Student) Student
}

type Student struct {
	ID   int
	Name string
}
