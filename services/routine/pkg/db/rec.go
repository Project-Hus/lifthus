package db

// import (
// 	"context"
// 	"routine/ent"
// 	"routine/ent/routineactrec"
// 	"routine/pkg/dto"
// )

// func UpdateRoutineActRec(
// 	dbClient *ent.Client,
// 	c context.Context,
// 	newRARDto *dto.UpdateRoutineActRecDto,
// ) (*ent.RoutineActRec, error) {
// 	updatingRAR := dbClient.RoutineActRec.UpdateOneID(newRARDto.ID).
// 		Where(routineactrec.AuthorEQ(*newRARDto.Author)).
// 		SetNillableCurrentReps(newRARDto.CurrentReps).
// 		SetNillableCurrentLap(newRARDto.CurrentLap).
// 		SetNillableStartedAt(newRARDto.StartedAt).
// 		SetNillableImage(newRARDto.Image).
// 		SetNillableComment(newRARDto.Comment)

// 	if newRARDto.Status != nil {
// 		updatingRAR = updatingRAR.SetStatus(routineactrec.Status(*newRARDto.Status))
// 	}
// 	updatedRAR, err := updatingRAR.Save(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return updatedRAR, nil
// }

// func CreateWeeklyProgramRec(
// 	dbClient *ent.Client,
// 	c context.Context,
// 	newPRec dto.CreateWeeklyProgramRecDto,
// ) (prid uint64, err error) {
// 	tx, err := dbClient.Tx(c)
// 	if err != nil {
// 		return 0, err
// 	}
// 	newRec, err := tx.ProgramRec.Create().
// 		SetAuthor(newPRec.Author).
// 		SetProgramID(newPRec.ProgramID).
// 		SetStartDate(newPRec.StartDate).
// 		SetEndDate(newPRec.EndDate).
// 		SetStatus("waiting").
// 		SetComment(newPRec.Comment).
// 		Save(c)
// 	if err != nil {
// 		return 0, rollback(tx, err)
// 	}

// 	prid = newRec.ID

// 	weeklyRoutineRecBulk := make([]*ent.WeeklyRoutineRecCreate, len(newPRec.WeeklyRoutineRecs))
// 	dailyRoutineRecDtos := make([][]dto.CreateWeeklyDailyRoutineRecDto, len(newPRec.WeeklyRoutineRecs))
// 	var cntDRR int
// 	for i, wrr := range newPRec.WeeklyRoutineRecs {
// 		weeklyRoutineRecBulk[i] = tx.WeeklyRoutineRec.Create().
// 			SetProgramRecID(prid).
// 			SetWeeklyRoutineID(wrr.WeeklyRoutineID).
// 			SetWeek(wrr.Week).
// 			SetStartDate(wrr.StartDate)
// 		for _, drr := range wrr.DayRoutineRecs {
// 			dailyRoutineRecDtos[i] = append(dailyRoutineRecDtos[i], drr)
// 			cntDRR++
// 		}
// 	}
// 	weeklyRoutineRecs, err := tx.WeeklyRoutineRec.CreateBulk(weeklyRoutineRecBulk...).Save(c)
// 	if err != nil {
// 		return 0, rollback(tx, err)
// 	}

// 	dailyRoutineRecBulk := make([]*ent.DailyRoutineRecCreate, cntDRR)
// 	routineActRecDtos := make([][]dto.CreateWeeklyRoutineActRecDto, cntDRR)
// 	var toDRR int
// 	var cntRAR int
// 	for i, wrr := range weeklyRoutineRecs {
// 		for _, drr := range dailyRoutineRecDtos[i] {
// 			dailyRoutineRecBulk[toDRR] =
// 				tx.DailyRoutineRec.Create().
// 					SetAuthor(newPRec.Author).
// 					SetWeeklyRoutineRec(wrr).
// 					SetStatus("waiting").
// 					SetDate(drr.Date)
// 			for _, rar := range drr.RoutineActRecs {
// 				routineActRecDtos[toDRR] = append(routineActRecDtos[toDRR], rar)
// 				cntRAR++
// 			}
// 			toDRR++
// 		}
// 	}
// 	dailyRoutineRecs, err := tx.DailyRoutineRec.CreateBulk(dailyRoutineRecBulk...).Save(c)
// 	if err != nil {
// 		return 0, rollback(tx, err)
// 	}

// 	routineActRecBulk := make([]*ent.RoutineActRecCreate, cntRAR)
// 	var toRAR int
// 	for i, drr := range dailyRoutineRecs {
// 		for _, rar := range routineActRecDtos[i] {
// 			routineActRecBulk[toRAR] =
// 				tx.RoutineActRec.Create().
// 					SetAuthor(newPRec.Author).
// 					SetDailyRoutineRec(drr).
// 					SetRoutineActID(rar.RoutineActID).
// 					SetActID(rar.ActID).
// 					SetOrder(rar.Order).
// 					SetNillableReps(rar.Reps).
// 					SetNillableLap(rar.Lap).
// 					SetStatus("waiting")
// 			toRAR++
// 		}
// 	}
// 	_, err = tx.RoutineActRec.CreateBulk(routineActRecBulk...).Save(c)
// 	if err != nil {
// 		return 0, rollback(tx, err)
// 	}

// 	tx.Commit()

// 	return prid, nil
// }
