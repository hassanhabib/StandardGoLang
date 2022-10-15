package Storages

import (
	"StandardGoLang/StandardGoLang.Core/Models/Students"
)

type StorageBrokerMock struct {
	PassedInStudent        Students.Student
	InsertStudentCallCount int
}

func (storageBrokerMock *StorageBrokerMock) InsertStudent(student Students.Student) Students.Student {
	storageBrokerMock.InsertStudentCallCount++
	storageBrokerMock.PassedInStudent = student

	return student
}