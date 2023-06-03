package db

import (
	"context"
	"fmt"
	"routine/common/dto"
	"routine/ent"
	"sync"
)

func CreateDailyProgram(dbClient *ent.Client, c context.Context, p interface{}) (newProgram *ent.Program, err error) {
	return nil, nil
}

// CreateWeeklyProgram creates weekly program and returns created program's ID.
func CreateWeeklyProgram(dbClient *ent.Client, c context.Context, p *dto.CreateWeeklyProgramDto) (pid uint64, err error) {
	// first query tags and create tags if not exists
	tags, err := QueryAndCreateTags(dbClient, c, p.Tags)
	if err != nil {
		return 0, err
	}

	// to conduct atomic operation, use transaction
	tx, err := dbClient.Tx(c)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	// not to query db again, get each index first
	weekIdxMap := make(map[int]int)   // e.g. [week 1] : week 1's index
	dayIdxMap := make(map[[2]int]int) // e.g. [week 2, day 3] : week 2's day 3's index
	// do this concurrently
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, wr := range p.WeeklyRoutines {
			weekIdxMap[wr.Week] = i
		}
	}()
	go func() {
		defer wg.Done()
		for i, dr := range p.DailyRoutines {
			dayIdxMap[[2]int{dr.Week, dr.Day}] = i
		}
	}()

	// first create program while getting indices.
	newProgram, err := tx.Program.Create().
		SetTitle(p.Title).
		SetType("weekly").
		SetAuthor(p.Author).
		SetNillableDescription(p.Description).
		SetNillableImage(p.Image).
		AddTagIDs(tags...).
		Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create program: %w", err))
	}

	wg.Wait()

	// then create weekly routines
	bulkW := make([]*ent.WeeklyRoutineCreate, len(p.WeeklyRoutines))
	for i, wr := range p.WeeklyRoutines {
		bulkW[i] = tx.WeeklyRoutine.Create().SetProgramID(newProgram.ID).SetWeek(wr.Week)
	}
	newWeeklyRoutines, err := tx.WeeklyRoutine.CreateBulk(bulkW...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create weekly routines: %w", err))
	}

	// then create daily routines
	bulkD := make([]*ent.DailyRoutineCreate, len(p.DailyRoutines))
	for i, dr := range p.DailyRoutines {
		weekId := newWeeklyRoutines[weekIdxMap[dr.Week]].ID
		bulkD[i] = tx.DailyRoutine.Create().SetProgramID(newProgram.ID).SetWeeklyRoutineID(weekId).SetDay(dr.Day)
	}
	newDailyRoutines, err := tx.DailyRoutine.CreateBulk(bulkD...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create daily routines: %w", err))
	}

	// then create routine acts
	bulkA := make([]*ent.RoutineActCreate, len(p.RoutineActs))
	for i, ra := range p.RoutineActs {
		dayId := newDailyRoutines[dayIdxMap[[2]int{ra.Week, ra.Day}]].ID
		bulkA[i] = tx.RoutineAct.Create().
			SetDailyRoutineID(dayId).SetActID(ra.ActID).
			SetNillableReps(ra.Reps).SetNillableLap(ra.Lap).
			SetNillableWRatio(ra.WRatio).SetWarmup(ra.Warmup)
	}
	_, err = tx.RoutineAct.CreateBulk(bulkA...).Save(c)
	if err != nil {
		return 0, rollback(tx, fmt.Errorf("failed to create routine acts: %w", err))
	}

	return newProgram.ID, nil
}
