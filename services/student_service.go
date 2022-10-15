package services

import "StandardGoLang/core"

func NewStudentService(broker core.StudentStorageBroker) *StudentService {
	return &StudentService{
		storageBroker: broker,
	}
}

type StudentService struct {
	storageBroker core.StudentStorageBroker
}

func (s *StudentService) AddStudent(student core.Student) core.Student {
	return s.storageBroker.InsertStudent(student)
}
