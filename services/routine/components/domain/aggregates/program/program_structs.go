package program

import (
	"routine/components/domain/aggregates/user"
	"time"
)

type NewProgramInfo struct {
	Parent      *uint64
	Title       string
	Author      user.User
	ProgramType ProgramType
	Iteration   int
	ImageSrcs   []string
	Description string
}

type ProgramInfo struct {
	Id *uint64

	Slug        string
	Code        string
	Parent      *uint64
	Title       string
	Author      user.User
	ProgramType ProgramType
	Iteration   int
	ImageSrcs   []string
	Description string

	createdAt *time.Time
	updatedAt *time.Time
}
