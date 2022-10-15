package brokers

import "StandardGoLang/core"

type StudentStorageBroker struct {
}

func (s *StudentStorageBroker) InsertStudent(student core.Student) core.Student {
	return student
}
