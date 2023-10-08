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
		{Name: "code", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "act_type", Type: field.TypeEnum, Enums: []string{"weight", "time", "simple"}},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "author", Type: field.TypeUint64},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ActsTable holds the schema information for the "acts" table.
	ActsTable = &schema.Table{
		Name:       "acts",
		Columns:    ActsColumns,
		PrimaryKey: []*schema.Column{ActsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "act_code",
				Unique:  true,
				Columns: []*schema.Column{ActsColumns[1]},
			},
			{
				Name:    "act_name",
				Unique:  false,
				Columns: []*schema.Column{ActsColumns[3]},
			},
		},
	}
	// ActImagesColumns holds the columns for the "act_images" table.
	ActImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "order", Type: field.TypeUint},
		{Name: "act_version_id", Type: field.TypeUint64},
		{Name: "image_id", Type: field.TypeUint64},
	}
	// ActImagesTable holds the schema information for the "act_images" table.
	ActImagesTable = &schema.Table{
		Name:       "act_images",
		Columns:    ActImagesColumns,
		PrimaryKey: []*schema.Column{ActImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_images_act_versions_act_version",
				Columns:    []*schema.Column{ActImagesColumns[2]},
				RefColumns: []*schema.Column{ActVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "act_images_images_image",
				Columns:    []*schema.Column{ActImagesColumns[3]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "actimage_act_version_id_image_id",
				Unique:  true,
				Columns: []*schema.Column{ActImagesColumns[2], ActImagesColumns[3]},
			},
		},
	}
	// ActVersionsColumns holds the columns for the "act_versions" table.
	ActVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "act_code", Type: field.TypeString, Size: 20},
		{Name: "version", Type: field.TypeUint},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "act_act_versions", Type: field.TypeUint64},
	}
	// ActVersionsTable holds the schema information for the "act_versions" table.
	ActVersionsTable = &schema.Table{
		Name:       "act_versions",
		Columns:    ActVersionsColumns,
		PrimaryKey: []*schema.Column{ActVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "act_versions_acts_act_versions",
				Columns:    []*schema.Column{ActVersionsColumns[6]},
				RefColumns: []*schema.Column{ActsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DailyRoutinesColumns holds the columns for the "daily_routines" table.
	DailyRoutinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "program_version_code", Type: field.TypeString, Size: 20},
		{Name: "day", Type: field.TypeUint},
		{Name: "program_version_daily_routines", Type: field.TypeUint64},
	}
	// DailyRoutinesTable holds the schema information for the "daily_routines" table.
	DailyRoutinesTable = &schema.Table{
		Name:       "daily_routines",
		Columns:    DailyRoutinesColumns,
		PrimaryKey: []*schema.Column{DailyRoutinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_routines_program_versions_daily_routines",
				Columns:    []*schema.Column{DailyRoutinesColumns[4]},
				RefColumns: []*schema.Column{ProgramVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "src", Type: field.TypeString, Unique: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
	}
	// ProgramsColumns holds the columns for the "programs" table.
	ProgramsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "program_type", Type: field.TypeEnum, Enums: []string{"weekly", "daily"}},
		{Name: "title", Type: field.TypeString, Size: 50},
		{Name: "author", Type: field.TypeUint64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "version_derived_from", Type: field.TypeString, Size: 20},
	}
	// ProgramsTable holds the schema information for the "programs" table.
	ProgramsTable = &schema.Table{
		Name:       "programs",
		Columns:    ProgramsColumns,
		PrimaryKey: []*schema.Column{ProgramsColumns[0]},
	}
	// ProgramImagesColumns holds the columns for the "program_images" table.
	ProgramImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "order", Type: field.TypeUint},
		{Name: "program_version_id", Type: field.TypeUint64},
		{Name: "image_id", Type: field.TypeUint64},
	}
	// ProgramImagesTable holds the schema information for the "program_images" table.
	ProgramImagesTable = &schema.Table{
		Name:       "program_images",
		Columns:    ProgramImagesColumns,
		PrimaryKey: []*schema.Column{ProgramImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "program_images_program_versions_program_version",
				Columns:    []*schema.Column{ProgramImagesColumns[2]},
				RefColumns: []*schema.Column{ProgramVersionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "program_images_images_image",
				Columns:    []*schema.Column{ProgramImagesColumns[3]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "programimage_program_version_id_image_id",
				Unique:  true,
				Columns: []*schema.Column{ProgramImagesColumns[2], ProgramImagesColumns[3]},
			},
		},
	}
	// ProgramVersionsColumns holds the columns for the "program_versions" table.
	ProgramVersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "program_code", Type: field.TypeString, Size: 20},
		{Name: "version", Type: field.TypeUint},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "program_program_versions", Type: field.TypeUint64},
	}
	// ProgramVersionsTable holds the schema information for the "program_versions" table.
	ProgramVersionsTable = &schema.Table{
		Name:       "program_versions",
		Columns:    ProgramVersionsColumns,
		PrimaryKey: []*schema.Column{ProgramVersionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "program_versions_programs_program_versions",
				Columns:    []*schema.Column{ProgramVersionsColumns[6]},
				RefColumns: []*schema.Column{ProgramsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RoutineActsColumns holds the columns for the "routine_acts" table.
	RoutineActsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "daily_routine_code", Type: field.TypeString, Size: 20},
		{Name: "order", Type: field.TypeUint},
		{Name: "act_version", Type: field.TypeString, Size: 20},
		{Name: "stage", Type: field.TypeEnum, Enums: []string{"warmup", "main", "cooldown"}},
		{Name: "reps_or_meters", Type: field.TypeUint},
		{Name: "ratio_or_secs", Type: field.TypeFloat64},
		{Name: "daily_routine_routine_acts", Type: field.TypeUint64},
	}
	// RoutineActsTable holds the schema information for the "routine_acts" table.
	RoutineActsTable = &schema.Table{
		Name:       "routine_acts",
		Columns:    RoutineActsColumns,
		PrimaryKey: []*schema.Column{RoutineActsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "routine_acts_daily_routines_routine_acts",
				Columns:    []*schema.Column{RoutineActsColumns[7]},
				RefColumns: []*schema.Column{DailyRoutinesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActsTable,
		ActImagesTable,
		ActVersionsTable,
		DailyRoutinesTable,
		ImagesTable,
		ProgramsTable,
		ProgramImagesTable,
		ProgramVersionsTable,
		RoutineActsTable,
	}
)

func init() {
	ActImagesTable.ForeignKeys[0].RefTable = ActVersionsTable
	ActImagesTable.ForeignKeys[1].RefTable = ImagesTable
	ActVersionsTable.ForeignKeys[0].RefTable = ActsTable
	DailyRoutinesTable.ForeignKeys[0].RefTable = ProgramVersionsTable
	ProgramImagesTable.ForeignKeys[0].RefTable = ProgramVersionsTable
	ProgramImagesTable.ForeignKeys[1].RefTable = ImagesTable
	ProgramVersionsTable.ForeignKeys[0].RefTable = ProgramsTable
	RoutineActsTable.ForeignKeys[0].RefTable = DailyRoutinesTable
}
