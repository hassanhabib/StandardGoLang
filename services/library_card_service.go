package services

import "StandardGoLang/core"

var libraryCardNumber int = 1

func NewLibraryCardService(broker core.LibraryCardBroker) *LibraryCardService {
	return &LibraryCardService{
		broker: broker,
	}
}

type LibraryCardService struct {
	broker core.LibraryCardBroker
}

func (s LibraryCardService) CreateLibraryCardForStudent(student core.Student) core.LibraryCard {
	card := core.LibraryCard{
		Owner:  student,
		Number: libraryCardNumber,
	}

	libraryCardNumber++

	return s.broker.InsertLibraryCard(card)
}
