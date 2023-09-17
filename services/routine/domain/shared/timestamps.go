package domain

import "time"

type Timestamps struct {
	createdAt time.Time
	updatedAt *time.Time
}

func TimestampsFrom(
	createdAt time.Time,
	updatedAt *time.Time,
) Timestamps {
	return Timestamps{
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (ts Timestamps) CreatedAt() time.Time {
	return ts.createdAt
}

func (ts Timestamps) UpdatedAt() *time.Time {
	return ts.updatedAt
}
