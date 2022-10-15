package Storages

import "StandardGoLang/StandardGoLang.Core/Models/Students"

type StorageBroker struct {
}

func (storageBroker StorageBroker) InsertStudent(student Students.Student) Students.Student {
	return student
}
