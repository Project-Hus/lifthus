package program

func CreateFirstWeeklyProgramVersion(desc string, wrs WeeklyRoutines) (*Version, error) {
	if !IsVersionDescValid(desc) {
		return nil, ErrInvalidDescription
	}
	return &Version{
		version:        1,
		description:    desc,
		weeklyRoutines: &wrs,
		dailyRoutines:  nil,
	}, nil
}

func UpgradeWeeklyProgramVersion(p *Program, desc string, wrs WeeklyRoutines) (*Version, error) {
	if !IsVersionDescValid(desc) {
		return nil, ErrInvalidDescription
	}
	return &Version{
		version:        p.versions.LatestVersion().version + 1,
		description:    desc,
		weeklyRoutines: &wrs,
		dailyRoutines:  nil,
	}, nil
}
