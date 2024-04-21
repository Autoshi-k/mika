package memory

import (
	"something/domain/repositories/timetable"
	"sync"
)

type memoryRepository struct {
	sync.Mutex
}

func NewTimetableRepository() timetable.Repository {
	return memoryRepository{}
}
