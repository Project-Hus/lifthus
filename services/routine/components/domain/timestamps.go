package domain

import "time"

func TimestampsFrom(createdAt time.Time, updatedAt *time.Time) Timestamps {
	return Timestamps{createdAt: createdAt, updatedAt: updatedAt}
}

type Timestamps struct {
	createdAt time.Time
	updatedAt *time.Time
}

func (t Timestamps) CreatedAt() time.Time {
	return t.createdAt
}

func (t Timestamps) UpdatedAt() *time.Time {
	return t.updatedAt
}
