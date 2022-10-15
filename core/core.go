package core

type StudentStorageBroker interface {
	InsertStudent(student Student) Student
}

type StudentService interface {
	AddStudent(student Student) Student
}

type LibraryCardService interface {
	CreateLibraryCardForStudent(student Student) LibraryCard
}

type LibraryCardBroker interface {
	InsertLibraryCard(card LibraryCard) LibraryCard
}

type Student struct {
	ID   int
	Name string
}

type LibraryCard struct {
	Owner  Student
	Number int
}
