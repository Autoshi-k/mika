package entity

import (
	"fmt"
	"time"
)

type WorkDay interface {
	Reduce(durations []TimeBlockTimings, blockers []TimeBlockTimings) WorkDay
	ReduceByDuration(duration time.Duration) WorkDay
	AvailableTimeSlots() []time.Time
}

type WorkDayListedSlots struct {
	timeSlots   []slot
	startTime   time.Time
	endTime     time.Time
	slotSpacing time.Duration
}

type slot struct {
	t                  time.Time
	inUse              bool
	invalidForDuration bool
}

func (wd WorkDayListedSlots) Reduce(durations []TimeBlockTimings, blockers []TimeBlockTimings) WorkDay {
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

func (wd WorkDayListedSlots) ReduceByDuration(duration time.Duration) WorkDay {

	for i, ts := range wd.timeSlots {
		if !ts.inUse {

		}
	}
	return wd
}

func (wd WorkDayListedSlots) AvailableTimeSlots() []time.Time {
	var availableSlots []time.Time
	for slot, ok := range wd.timeSlots {
		if ok {
			availableSlots = append(availableSlots, slot)
		}
	}

	return availableSlots
}

func NewWorkday(startTime, endTime time.Time, spacingMinutes int) (WorkDay, error) {
	if startTime.IsZero() {
		return nil, fmt.Errorf("startTime is empty %v", startTime)
	}

	if endTime.IsZero() {
		return nil, fmt.Errorf("endTime is empty %v", endTime)
	}

	if spacingMinutes <= 0 {
		return nil, fmt.Errorf("spacing should be above zero, recommanded [10, 15, 30, ...] %d", spacingMinutes)
	}

	spacingDuration := time.Duration(spacingMinutes) * time.Minute

	var slots []slot
	timestamp := startTime
	for timestamp.Before(endTime) {
		slots = append(slots, slot{
			t:     timestamp,
			inUse: false,
		})
		timestamp = timestamp.Add(spacingDuration)
	}

	return WorkDayListedSlots{
		timeSlots:   slots,
		startTime:   startTime,
		endTime:     endTime,
		slotSpacing: spacingDuration,
	}, nil
}

func getStartAndEndIndexes(spacing time.Duration, startTime time.Time, duration TimeBlockTimings) (start int, end int) {
	diff := duration.StartTime.Sub(startTime).Minutes()

	if spacing == 0 {
		spacing = 1
	}

	shit := diff / spacing.Minutes()
	start = int(shit) - 1

	diff = duration.EndTime.Sub(duration.StartTime).Minutes()
	shit = diff / spacing.Minutes()
	end = int(shit) - 1

	return
}
