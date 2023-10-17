package entrepo

import (
	"context"
	"routine/internal/domain/aggregates/program"
	"routine/internal/ent"
	"routine/internal/ent/act"
	"routine/internal/ent/routineact"
	"routine/internal/ent/s3image"
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
	var parcode *string
	var parver *int
	if parent := np.ParentProgramVersion(); parent != nil {
		parcode = (*string)(&parent.ProgramCode)
		parver = (*int)(&parent.ProgramVersionNumber)
	}
	ep, err := tx.Program.Create().
		SetCode(string(np.Code())).
		SetProgramType(ptype).
		SetTitle(string(np.Title())).
		SetAuthor(int64(np.Author())).
		SetCreatedAt(time.Time(np.CreatedAt())).
		SetNillableParentProgram(parcode).
		SetNillableParentVersion(parver).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = repo.insertNewProgramReleases(ctx, ep, np.Releases())
	if err != nil {
		return nil, err
	}
	return ProgramFromEntProgram(ep)
}

func (repo *EntProgramRepository) insertNewProgramReleases(ctx context.Context, ep *ent.Program, pvs program.ProgramReleases) ([]*ent.ProgramRelease, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	epvs := make([]*ent.ProgramRelease, len(pvs))
	for i, pv := range pvs {
		epv, err := tx.ProgramRelease.Create().
			SetProgram(ep).
			SetVersion(int(pv.Version())).
			SetCreatedAt(time.Time(pv.CreatedAt())).
			SetText(string(pv.Text())).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		_, err = repo.findImagesAndSetToProgramVersion(ctx, epv, pv.ImageSrcs())
		if err != nil {
			return nil, err
		}
		_, err = repo.insertNewRoutines(ctx, epv, pv.Routines())
		if err != nil {
			return nil, err
		}
		epvs[i] = epv
	}
	ep.Edges.ProgramReleases = epvs
	return epvs, nil
}

func (repo *EntProgramRepository) findImagesAndSetToProgramVersion(ctx context.Context, epv *ent.ProgramRelease, srcs program.ProgramImageSrcs) ([]*ent.S3ProgramImage, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	epis := make([]*ent.S3ProgramImage, len(srcs))
	for i, src := range srcs {
		ord := i + 1
		eimg, err := tx.S3Image.Query().Where(s3image.Src(src)).First(ctx)
		if err != nil {
			return nil, err
		}
		epimg, err := tx.S3ProgramImage.Create().SetOrder(int(ord)).SetS3Image(eimg).Save(ctx)
		if err != nil {
			return nil, err
		}
		epis[i] = epimg
	}
	epv.Edges.S3ProgramImages = epis
	return epis, nil
}

func (repo *EntProgramRepository) insertNewRoutines(ctx context.Context, epv *ent.ProgramRelease, drs program.Routines) ([]*ent.Routine, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	edrs := make([]*ent.Routine, len(drs))
	for i, dr := range drs {
		edr, err := tx.Routine.Create().
			SetProgramRelease(epv).
			SetDay(int(dr.Day())).
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
	epv.Edges.Routines = edrs
	return edrs, nil
}

func (repo *EntProgramRepository) insertNewRoutineActs(ctx context.Context, edr *ent.Routine, ras program.RoutineActs) ([]*ent.RoutineAct, error) {
	tx, finally, err := repo.Tx(ctx)
	defer finally(&err)
	if err != nil {
		return nil, err
	}
	eras := make([]*ent.RoutineAct, len(ras))
	for i, ra := range ras {
		eav, err := tx.Act.Query().Where(act.Code(string(ra.Act()))).First(ctx)
		if err != nil {
			return nil, err
		}
		era, err := tx.RoutineAct.Create().
			SetAct(eav).
			SetActCode(string(ra.Act())).
			SetRoutine(edr).
			SetOrder(int(ra.Order())).
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
