package Students

import (
	Students2 "StandardGoLang/StandardGoLang.Core/Models/Students"
	"StandardGoLang/StandardGoLang.Core/Services/Foundations/Students"
	"StandardGoLang/StandardGoLang.Tests/Mocks/Brokers/Storages"
	"testing"
)

func TestShouldAddStudent(testing *testing.T) {
	// given
	var storageBrokerMock = Storages.StorageBrokerMock{}
	var studentService = Students.StudentService{}

	var inputStudent = Students2.Student{
		Id:   10,
		Name: "Hassan Habib",
	}

	var expectedStudent Students2.Student = inputStudent
	studentService.Init(&storageBrokerMock)

	// when
	var actualStudent Students2.Student = studentService.AddStudent(inputStudent)

	// then
	if actualStudent != expectedStudent {
		testing.Errorf(
			"expected: %q actual %q",
			expectedStudent,
			actualStudent)
	}

	if storageBrokerMock.InsertStudentCallCount != 1 {
		testing.Errorf("Expected InsertStudent to be called once")
	}

	if storageBrokerMock.PassedInStudent != inputStudent {
		testing.Errorf(
			"Expected %q to be passed, but found %q",
			inputStudent,
			storageBrokerMock.PassedInStudent)
	}
}
