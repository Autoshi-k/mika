package entity

import "time"

type TimeBlockTimings struct {
	DurationDetails
	StartTime time.Time
	EndTime   time.Time
}

func NewTimeBlockTimings(startTime time.Time, durations DurationDetails) TimeBlockTimings {
	return TimeBlockTimings{
		DurationDetails: durations,
		StartTime:       startTime,
		EndTime:         startTime.Add(durations.Duration),
	}
}
