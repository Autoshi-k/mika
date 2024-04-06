package entity

import "time"

type Times []time.Time

func (ts Times) Contains(t time.Time) bool {
	for _, v := range ts {
		if v.Equal(t) {
			return true
		}
	}

	return false
}
