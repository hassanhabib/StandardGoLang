package services

import (
	"StandardGoLang/core"
	"testing"
)

func TestLibraryCardService(t *testing.T) {
	broker := &mockLibraryCardBroker{}

	inputStudent := core.Student{
		ID:   10,
		Name: "Hassan Habib",
	}

	expectedCard := core.LibraryCard{
		Owner:  inputStudent,
		Number: 1,
	}

	service := NewLibraryCardService(broker)

	actualCard := service.CreateLibraryCardForStudent(inputStudent)

	if actualCard != expectedCard {
		t.Errorf(
			"expected: %q actual %q",
			expectedCard,
			actualCard)
	}

	if broker.callCount != 1 {
		t.Errorf("Expected CreateLibraryCardForStudent to be called once")
	}

	//Test increment of card number
	actualCard2 := service.CreateLibraryCardForStudent(inputStudent)
	if actualCard2.Number != 2 {
		t.Errorf("Expected number of 2nd LibraryCard to be 2, got %d", actualCard2.Number)
	}
}
