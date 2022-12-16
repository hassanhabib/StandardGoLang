package services

import "StandardGoLang/core"

type mockStudentStorageBroker struct {
	passedInStudent core.Student
	callCount       int
}

func (m *mockStudentStorageBroker) InsertStudent(student core.Student) core.Student {
	m.passedInStudent = student
	m.callCount++

	return student
}
