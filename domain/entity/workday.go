package entity

import (
	"time"
)

type WorkDay struct {
	timeSlots   []slot
	startTime   time.Time
	endTime     time.Time
	slotSpacing time.Duration
}

type slot struct {
	t     time.Time
	inUse bool
}

func (wd WorkDay) Reduce(durations []TimeBlockTimings, blockers []TimeBlockTimings) WorkDay {
	for _, duration := range append(durations, blockers...) {
		// todo what if changing timeslots spacing? won't find times
		startIndex, endIndex := getStartAndEndIndexes(wd.slotSpacing, wd.startTime, duration)

		i := startIndex
		for i <= endIndex {
			wd.timeSlots[i].inUse = true
			i++
		}
	}
	return wd
}

func (wd WorkDay) ReduceByDuration(duration time.Duration) WorkDay {
	for i, ts := range wd.timeSlots {
		if ts.inUse {
			curr := slot
			for curr.Before(duration.EndTime) || curr.Equal(duration.EndTime) {

			}
		}
	}
	return wd
}

func (wd WorkDay) AvailableTimeSlots() []time.Time {
	var availableSlots []time.Time
	for slot, ok := range wd.timeSlots {
		if ok {
			availableSlots = append(availableSlots, slot)
		}
	}

	return availableSlots
}

func getStartAndEndIndexes(spacing time.Duration, startTime time.Time, duration TimeBlockTimings) (start int, end int) {
	diff := duration.StartTime.Sub(startTime).Minutes()
	shit := diff / spacing.Minutes()
	start = int(shit) - 1

	diff = duration.EndTime.Sub(duration.StartTime).Minutes()
	shit = diff / spacing.Minutes()
	end = int(shit) - 1

	return
}
