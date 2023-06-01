// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActsColumns holds the columns for the "acts" table.
	ActsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"rep", "lap", "simple"}},
		{Name: "author", Type: field.TypeUint64},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 5000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "weight", Type: field.TypeBool, Default: false},
		{Name: "bodyweight", Type: field.TypeBool, Default: false},
		{Name: "cardio", Type: field.TypeBool, Default: false},
		{Name: "upper", Type: field.TypeBool, Default: false},
		{Name: "lower", Type: field.TypeBool, Default: false},
		{Name: "full", Type: field.TypeBool, Default: false},
		{Name: "arms", Type: field.TypeBool, Default: false},
		{Name: "shoulders", Type: field.TypeBool, Default: false},
		{Name: "chest", Type: field.TypeBool, Default: false},
		{Name: "core", Type: field.TypeBool, Default: false},
		{Name: "upper_back", Type: field.TypeBool, Default: false},
		{Name: "lower_back", Type: field.TypeBool, Default: false},
		{Name: "legs", Type: field.TypeBool, Default: false},
		{Name: "legs_front", Type: field.TypeBool, Default: false},
		{Name: "legs_back", Type: field.TypeBool, Default: false},
		{Name: "etc", Type: field.TypeBool, Default: false},
	}
	// ActsTable holds the schema information for the "acts" table.
	ActsTable = &schema.Table{
		Name:       "acts",
		Columns:    ActsColumns,
		PrimaryKey: []*schema.Column{ActsColumns[0]},
	}
	// BodyInfosColumns holds the columns for the "body_infos" table.
	BodyInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "author", Type: field.TypeUint64, Unique: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "height", Type: field.TypeFloat64, Nullable: true},
		{Name: "body_weight", Type: field.TypeFloat64, Nullable: true},
		{Name: "body_fat_mass", Type: field.TypeFloat64, Nullable: true},
		{Name: "skeletal_muscle_mass", Type: field.TypeFloat64, Nullable: true},
		{Name: "program_rec_id", Type: field.TypeUint64, Nullable: true},
	}
	// BodyInfosTable holds the schema information for the "body_infos" table.
	BodyInfosTable = &schema.Table{
		Name:       "body_infos",
		Columns:    BodyInfosColumns,
		PrimaryKey: []*schema.Column{BodyInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "body_infos_program_recs_body_info",
				Columns:    []*schema.Column{BodyInfosColumns[7]},
				RefColumns: []*schema.Column{ProgramRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DailyRoutinesColumns holds the columns for the "daily_routines" table.
	DailyRoutinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "day", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "program_id", Type: field.TypeUint64, Nullable: true},
		{Name: "weekly_routine_id", Type: field.TypeUint64, Nullable: true},
	}
	// DailyRoutinesTable holds the schema information for the "daily_routines" table.
	DailyRoutinesTable = &schema.Table{
		Name:       "daily_routines",
		Columns:    DailyRoutinesColumns,
		PrimaryKey: []*schema.Column{DailyRoutinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_routines_programs_daily_routines",
				Columns:    []*schema.Column{DailyRoutinesColumns[4]},
				RefColumns: []*schema.Column{ProgramsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "daily_routines_weekly_routines_daily_routines",
				Columns:    []*schema.Column{DailyRoutinesColumns[5]},
				RefColumns: []*schema.Column{WeeklyRoutinesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// DailyRoutineRecsColumns holds the columns for the "daily_routine_recs" table.
	DailyRoutineRecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "author", Type: field.TypeUint64},
		{Name: "date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "daily_routine_id", Type: field.TypeUint64, Nullable: true},
		{Name: "program_rec_id", Type: field.TypeUint64, Nullable: true},
		{Name: "weekly_routine_rec_id", Type: field.TypeUint64, Nullable: true},
	}
	// DailyRoutineRecsTable holds the schema information for the "daily_routine_recs" table.
	DailyRoutineRecsTable = &schema.Table{
		Name:       "daily_routine_recs",
		Columns:    DailyRoutineRecsColumns,
		PrimaryKey: []*schema.Column{DailyRoutineRecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_routine_recs_daily_routines_daily_routine_recs",
				Columns:    []*schema.Column{DailyRoutineRecsColumns[7]},
				RefColumns: []*schema.Column{DailyRoutinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "daily_routine_recs_program_recs_daily_routine_recs",
				Columns:    []*schema.Column{DailyRoutineRecsColumns[8]},
				RefColumns: []*schema.Column{ProgramRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "daily_routine_recs_weekly_routine_recs_daily_routine_recs",
				Columns:    []*schema.Column{DailyRoutineRecsColumns[9]},
				RefColumns: []*schema.Column{WeeklyRoutineRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// OneRepMaxesColumns holds the columns for the "one_rep_maxes" table.
	OneRepMaxesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "author", Type: field.TypeUint64, Unique: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "one_rep_max", Type: field.TypeFloat64, Nullable: true},
		{Name: "certified", Type: field.TypeBool, Default: false},
		{Name: "calculated", Type: field.TypeBool, Default: false},
		{Name: "act_id", Type: field.TypeUint64},
		{Name: "program_rec_id", Type: field.TypeUint64, Nullable: true},
	}
	// OneRepMaxesTable holds the schema information for the "one_rep_maxes" table.
	OneRepMaxesTable = &schema.Table{
		Name:       "one_rep_maxes",
		Columns:    OneRepMaxesColumns,
		PrimaryKey: []*schema.Column{OneRepMaxesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "one_rep_maxes_acts_one_rep_maxes",
				Columns:    []*schema.Column{OneRepMaxesColumns[6]},
				RefColumns: []*schema.Column{ActsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "one_rep_maxes_program_recs_one_rep_max",
				Columns:    []*schema.Column{OneRepMaxesColumns[7]},
				RefColumns: []*schema.Column{ProgramRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProgramsColumns holds the columns for the "programs" table.
	ProgramsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 50},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"weekly", "daily"}},
		{Name: "author", Type: field.TypeUint64},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 5000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ProgramsTable holds the schema information for the "programs" table.
	ProgramsTable = &schema.Table{
		Name:       "programs",
		Columns:    ProgramsColumns,
		PrimaryKey: []*schema.Column{ProgramsColumns[0]},
	}
	// ProgramRecsColumns holds the columns for the "program_recs" table.
	ProgramRecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "author", Type: field.TypeUint64},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 1000},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "program_id", Type: field.TypeUint64},
	}
	// ProgramRecsTable holds the schema information for the "program_recs" table.
	ProgramRecsTable = &schema.Table{
		Name:       "program_recs",
		Columns:    ProgramRecsColumns,
		PrimaryKey: []*schema.Column{ProgramRecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "program_recs_programs_program_recs",
				Columns:    []*schema.Column{ProgramRecsColumns[8]},
				RefColumns: []*schema.Column{ProgramsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RoutineActsColumns holds the columns for the "routine_acts" table.
	RoutineActsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "order", Type: field.TypeInt},
		{Name: "reps", Type: field.TypeInt, Nullable: true},
		{Name: "lap", Type: field.TypeInt, Nullable: true},
		{Name: "warmup", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "act_id", Type: field.TypeUint64},
		{Name: "daily_routine_id", Type: field.TypeUint64},
	}
	// RoutineActsTable holds the schema information for the "routine_acts" table.
	RoutineActsTable = &schema.Table{
		Name:       "routine_acts",
		Columns:    RoutineActsColumns,
		PrimaryKey: []*schema.Column{RoutineActsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routine_acts_acts_routine_acts",
				Columns:    []*schema.Column{RoutineActsColumns[7]},
				RefColumns: []*schema.Column{ActsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "routine_acts_daily_routines_routine_acts",
				Columns:    []*schema.Column{RoutineActsColumns[8]},
				RefColumns: []*schema.Column{DailyRoutinesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoutineActRecsColumns holds the columns for the "routine_act_recs" table.
	RoutineActRecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "routine_act_id", Type: field.TypeUint64, Nullable: true},
		{Name: "order", Type: field.TypeInt},
		{Name: "reps", Type: field.TypeInt, Nullable: true},
		{Name: "lap", Type: field.TypeInt, Nullable: true},
		{Name: "current_reps", Type: field.TypeInt, Default: 0},
		{Name: "current_lap", Type: field.TypeInt, Default: 0},
		{Name: "image", Type: field.TypeString, Nullable: true},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"history", "waiting", "proceeding", "completed", "failed", "canceled"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "act_id", Type: field.TypeUint64},
		{Name: "daily_routine_rec_id", Type: field.TypeUint64},
		{Name: "routine_act_routine_act_recs", Type: field.TypeUint64, Nullable: true},
	}
	// RoutineActRecsTable holds the schema information for the "routine_act_recs" table.
	RoutineActRecsTable = &schema.Table{
		Name:       "routine_act_recs",
		Columns:    RoutineActRecsColumns,
		PrimaryKey: []*schema.Column{RoutineActRecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routine_act_recs_acts_routine_act_recs",
				Columns:    []*schema.Column{RoutineActRecsColumns[12]},
				RefColumns: []*schema.Column{ActsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "routine_act_recs_daily_routine_recs_routine_act_recs",
				Columns:    []*schema.Column{RoutineActRecsColumns[13]},
				RefColumns: []*schema.Column{DailyRoutineRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "routine_act_recs_routine_acts_routine_act_recs",
				Columns:    []*schema.Column{RoutineActRecsColumns[14]},
				RefColumns: []*schema.Column{RoutineActsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "tag", Type: field.TypeString, Size: 20},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// WeeklyRoutinesColumns holds the columns for the "weekly_routines" table.
	WeeklyRoutinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "week", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "program_id", Type: field.TypeUint64},
	}
	// WeeklyRoutinesTable holds the schema information for the "weekly_routines" table.
	WeeklyRoutinesTable = &schema.Table{
		Name:       "weekly_routines",
		Columns:    WeeklyRoutinesColumns,
		PrimaryKey: []*schema.Column{WeeklyRoutinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weekly_routines_programs_weekly_routines",
				Columns:    []*schema.Column{WeeklyRoutinesColumns[4]},
				RefColumns: []*schema.Column{ProgramsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WeeklyRoutineRecsColumns holds the columns for the "weekly_routine_recs" table.
	WeeklyRoutineRecsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "program_rec_id", Type: field.TypeUint64},
		{Name: "weekly_routine_id", Type: field.TypeUint64},
	}
	// WeeklyRoutineRecsTable holds the schema information for the "weekly_routine_recs" table.
	WeeklyRoutineRecsTable = &schema.Table{
		Name:       "weekly_routine_recs",
		Columns:    WeeklyRoutineRecsColumns,
		PrimaryKey: []*schema.Column{WeeklyRoutineRecsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weekly_routine_recs_program_recs_weekly_routine_recs",
				Columns:    []*schema.Column{WeeklyRoutineRecsColumns[4]},
				RefColumns: []*schema.Column{ProgramRecsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "weekly_routine_recs_weekly_routines_weekly_routine_recs",
				Columns:    []*schema.Column{WeeklyRoutineRecsColumns[5]},
				RefColumns: []*schema.Column{WeeklyRoutinesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TagActsColumns holds the columns for the "tag_acts" table.
	TagActsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeUint64},
		{Name: "act_id", Type: field.TypeUint64},
	}
	// TagActsTable holds the schema information for the "tag_acts" table.
	TagActsTable = &schema.Table{
		Name:       "tag_acts",
		Columns:    TagActsColumns,
		PrimaryKey: []*schema.Column{TagActsColumns[0], TagActsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_acts_tag_id",
				Columns:    []*schema.Column{TagActsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_acts_act_id",
				Columns:    []*schema.Column{TagActsColumns[1]},
				RefColumns: []*schema.Column{ActsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TagProgramsColumns holds the columns for the "tag_programs" table.
	TagProgramsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeUint64},
		{Name: "program_id", Type: field.TypeUint64},
	}
	// TagProgramsTable holds the schema information for the "tag_programs" table.
	TagProgramsTable = &schema.Table{
		Name:       "tag_programs",
		Columns:    TagProgramsColumns,
		PrimaryKey: []*schema.Column{TagProgramsColumns[0], TagProgramsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_programs_tag_id",
				Columns:    []*schema.Column{TagProgramsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_programs_program_id",
				Columns:    []*schema.Column{TagProgramsColumns[1]},
				RefColumns: []*schema.Column{ProgramsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActsTable,
		BodyInfosTable,
		DailyRoutinesTable,
		DailyRoutineRecsTable,
		OneRepMaxesTable,
		ProgramsTable,
		ProgramRecsTable,
		RoutineActsTable,
		RoutineActRecsTable,
		TagsTable,
		WeeklyRoutinesTable,
		WeeklyRoutineRecsTable,
		TagActsTable,
		TagProgramsTable,
	}
)

func init() {
	BodyInfosTable.ForeignKeys[0].RefTable = ProgramRecsTable
	DailyRoutinesTable.ForeignKeys[0].RefTable = ProgramsTable
	DailyRoutinesTable.ForeignKeys[1].RefTable = WeeklyRoutinesTable
	DailyRoutineRecsTable.ForeignKeys[0].RefTable = DailyRoutinesTable
	DailyRoutineRecsTable.ForeignKeys[1].RefTable = ProgramRecsTable
	DailyRoutineRecsTable.ForeignKeys[2].RefTable = WeeklyRoutineRecsTable
	OneRepMaxesTable.ForeignKeys[0].RefTable = ActsTable
	OneRepMaxesTable.ForeignKeys[1].RefTable = ProgramRecsTable
	ProgramRecsTable.ForeignKeys[0].RefTable = ProgramsTable
	RoutineActsTable.ForeignKeys[0].RefTable = ActsTable
	RoutineActsTable.ForeignKeys[1].RefTable = DailyRoutinesTable
	RoutineActRecsTable.ForeignKeys[0].RefTable = ActsTable
	RoutineActRecsTable.ForeignKeys[1].RefTable = DailyRoutineRecsTable
	RoutineActRecsTable.ForeignKeys[2].RefTable = RoutineActsTable
	WeeklyRoutinesTable.ForeignKeys[0].RefTable = ProgramsTable
	WeeklyRoutineRecsTable.ForeignKeys[0].RefTable = ProgramRecsTable
	WeeklyRoutineRecsTable.ForeignKeys[1].RefTable = WeeklyRoutinesTable
	TagActsTable.ForeignKeys[0].RefTable = TagsTable
	TagActsTable.ForeignKeys[1].RefTable = ActsTable
	TagProgramsTable.ForeignKeys[0].RefTable = TagsTable
	TagProgramsTable.ForeignKeys[1].RefTable = ProgramsTable
}
