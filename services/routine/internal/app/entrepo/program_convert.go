package entrepo

import (
	"fmt"
	"routine/internal/domain"
	"routine/internal/domain/aggregates/program"
	"routine/internal/domain/aggregates/user"
	"routine/internal/ent"
	eprogram "routine/internal/ent/program"
)

func ProgramFromEntProgram(ep *ent.Program) (*program.Program, error) {
	switch ep.ProgramType {
	case eprogram.ProgramTypeWeekly:
		author := user.UserFrom(user.UserId(ep.Author))
		versions, err := VersionsFromEntVersions(ep.Edges.ProgramVersions)
		if err != nil {
			return nil, err
		}
		p, err := program.WeeklyProgramFrom(
			program.ProgramCode(ep.Code),
			program.ProgramTitle(ep.Title),
			*author,
			domain.CreatedAt(ep.CreatedAt),
			(*program.ProgramVersionCode)(ep.VersionDerivedFrom),
			versions,
		)
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unknown program type: %s", ep.ProgramType)
	}
}

func VersionsFromEntVersions(epvs []*ent.ProgramVersion) (pvs []*program.ProgramVersion, err error) {
	return nil, nil
}
