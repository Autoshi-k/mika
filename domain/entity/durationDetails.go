package entity

import "time"

type DurationDetails struct {
	Duration           time.Duration
	PreBufferDuration  time.Duration
	PostBufferDuration time.Duration
}

func (dd DurationDetails) CombinedDuration() time.Duration {
	return time.Duration(dd.PreBufferDuration.Seconds() + dd.PreBufferDuration.Seconds() + dd.PostBufferDuration.Seconds())
}
