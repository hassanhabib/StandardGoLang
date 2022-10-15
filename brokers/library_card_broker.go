package brokers

import "StandardGoLang/core"

type LibraryCardBroker struct {
}

func (l LibraryCardBroker) InsertLibraryCard(card core.LibraryCard) core.LibraryCard {
	return card
}

