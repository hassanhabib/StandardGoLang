package Students

import (
	"StandardGoLang/StandardGoLang.Core/Brokers/Storages"
	"StandardGoLang/StandardGoLang.Core/Models/Students"
)

type StudentService struct {
	storageBroker Storages.IStorageBroker
}

func (studentService *StudentService) Init(storageBroker Storages.IStorageBroker) {
	studentService.storageBroker = storageBroker
}

func (studentService *StudentService) AddStudent(student Students.Student) Students.Student {
	return studentService.storageBroker.InsertStudent(student)
}
