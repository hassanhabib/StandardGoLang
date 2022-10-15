package Storages

import "StandardGoLang/StandardGoLang.Core/Models/Students"

type IStorageBroker interface {
	InsertStudent(student Students.Student) Students.Student
}
