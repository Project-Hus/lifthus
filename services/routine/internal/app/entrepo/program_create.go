package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/program"
	"routine/internal/ent"
	eactversion "routine/internal/ent/actversion"
	"routine/internal/ent/image"
	"routine/internal/ent/routineact"
	"time"
)

func (repo *EntProgramRepository) insertNewProgram(ctx context.Context, np *program.Program) (*program.Program, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	ptype, err := entProgramTypeFrom(np.ProgramType().Type())
	if err != nil {
		return nil, err
	}
	ep, err := tx.Program.Create().
		SetCode(string(np.Code())).
		SetProgramType(ptype).
		SetTitle(string(np.Title())).
		SetAuthor(uint64(np.Author())).
		SetCreatedAt(time.Time(np.CreatedAt())).
		SetNillableVersionDerivedFrom((*string)(np.DerivedFrom())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = repo.insertNewProgramVersions(ctx, ep, np.Versions())
	if err != nil {
		return nil, err
	}
	return ProgramFromEntProgram(ep)
}

func (repo *EntProgramRepository) insertNewProgramVersions(ctx context.Context, ep *ent.Program, pvs program.ProgramVersions) ([]*ent.ProgramVersion, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	epvs := make([]*ent.ProgramVersion, len(pvs))
	for i, pv := range pvs {
		epv, err := tx.ProgramVersion.Create().
			SetCode(string(pv.Code())).
			SetProgram(ep).
			SetProgramCode(string(pv.Program())).
			SetCreatedAt(time.Time(pv.CreatedAt())).
			SetVersion(uint(pv.Version())).
			SetText(string(pv.Text())).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		_, err = repo.findImagesAndSetToProgramVersion(ctx, epv, pv.ImageSrcs())
		if err != nil {
			return nil, err
		}
		_, err = repo.insertNewDailyRoutines(ctx, epv, pv.DailyRoutines())
		if err != nil {
			return nil, err
		}
		epvs[i] = epv
	}
	ep.Edges.ProgramVersions = epvs
	return epvs, nil
}

func (repo *EntProgramRepository) findImagesAndSetToProgramVersion(ctx context.Context, epv *ent.ProgramVersion, srcs program.ProgramImageSrcs) ([]*ent.ProgramImage, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	epis := make([]*ent.ProgramImage, len(srcs))
	for i, src := range srcs {
		ord := i + 1
		eimg, err := tx.Image.Query().Where(image.Src(src)).First(ctx)
		if err != nil {
			return nil, err
		}
		epimg, err := tx.ProgramImage.Create().SetOrder(uint(ord)).SetImage(eimg).Save(ctx)
		if err != nil {
			return nil, err
		}
		epis[i] = epimg
	}
	epv.Edges.ProgramImages = epis
	return epis, nil
}

func (repo *EntProgramRepository) insertNewDailyRoutines(ctx context.Context, epv *ent.ProgramVersion, drs program.DailyRoutines) ([]*ent.DailyRoutine, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	edrs := make([]*ent.DailyRoutine, len(drs))
	for i, dr := range drs {
		edr, err := tx.DailyRoutine.Create().
			SetProgramVersion(epv).
			SetProgramVersionCode(epv.Code).
			SetCode(string(dr.Code())).
			SetDay(uint(dr.Day())).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		_, err = repo.insertNewRoutineActs(ctx, edr, dr.RoutineActs())
		if err != nil {
			return nil, err
		}
		edrs[i] = edr
	}
	epv.Edges.DailyRoutines = edrs
	return edrs, nil
}

func (repo *EntProgramRepository) insertNewRoutineActs(ctx context.Context, edr *ent.DailyRoutine, ras program.RoutineActs) ([]*ent.RoutineAct, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	eras := make([]*ent.RoutineAct, len(ras))
	for i, ra := range ras {
		eav, err := tx.ActVersion.Query().Where(eactversion.Code(string(ra.ActVersion()))).First(ctx)
		if err != nil {
			return nil, err
		}
		era, err := tx.RoutineAct.Create().
			SetActVersion(eav).
			SetActVersionCode(eav.Code).
			SetDailyRoutine(edr).
			SetDailyRoutineCode(edr.Code).
			SetOrder(uint(ra.Order())).
			SetStage(routineact.Stage(ra.Stage().Type())).
			SetRepsOrMeters(uint(ra.RepsOrMeters())).
			SetRatioOrSecs(float64(ra.RatioOrSecs())).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		eras[i] = era
	}
	edr.Edges.RoutineActs = eras
	return eras, nil
}
