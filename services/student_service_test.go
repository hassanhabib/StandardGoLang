package services

import (
	"StandardGoLang/core"
	"testing"
)

func TestStudentService(t *testing.T) {
	broker := &mockStudentStorageBroker{}

	inputStudent := core.Student{
		ID:   10,
		Name: "Hassan Habib",
	}

	expectedStudent := inputStudent

	studentService := NewStudentService(broker)

	actualStudent := studentService.AddStudent(inputStudent)

	if actualStudent != expectedStudent {
		t.Errorf(
			"expected: %q actual %q",
			expectedStudent,
			actualStudent)
	}

	if broker.callCount != 1 {
		t.Errorf("Expected InsertStudent to be called once")
	}

	if broker.passedInStudent != inputStudent {
		t.Errorf(
			"Expected %q to be passed, but found %q",
			inputStudent,
			broker.passedInStudent)
	}
}
