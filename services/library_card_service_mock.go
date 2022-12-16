package services

import "StandardGoLang/core"

type mockLibraryCardBroker struct {
	callCount int
}

func (m *mockLibraryCardBroker) InsertLibraryCard(card core.LibraryCard) core.LibraryCard {
	m.callCount++

	return card
}
