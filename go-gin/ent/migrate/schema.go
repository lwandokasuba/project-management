// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AssignmentsColumns holds the columns for the "assignments" table.
	AssignmentsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"active", "in_active", "completed"}},
		{Name: "project_assignments", Type: field.TypeUUID},
		{Name: "user_assignments", Type: field.TypeUUID},
	}
	// AssignmentsTable holds the schema information for the "assignments" table.
	AssignmentsTable = &schema.Table{
		Name:       "assignments",
		Columns:    AssignmentsColumns,
		PrimaryKey: []*schema.Column{AssignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "assignments_projects_assignments",
				Columns:    []*schema.Column{AssignmentsColumns[8]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "assignments_users_assignments",
				Columns:    []*schema.Column{AssignmentsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"created", "in_progress", "completed", "suspended"}, Default: "created"},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// StatementsColumns holds the columns for the "statements" table.
	StatementsColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"expense", "profit"}},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "assignment_statements", Type: field.TypeUUID},
	}
	// StatementsTable holds the schema information for the "statements" table.
	StatementsTable = &schema.Table{
		Name:       "statements",
		Columns:    StatementsColumns,
		PrimaryKey: []*schema.Column{StatementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "statements_assignments_statements",
				Columns:    []*schema.Column{StatementsColumns[6]},
				RefColumns: []*schema.Column{AssignmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "oid", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AssignmentsTable,
		ProjectsTable,
		StatementsTable,
		UsersTable,
	}
)

func init() {
	AssignmentsTable.ForeignKeys[0].RefTable = ProjectsTable
	AssignmentsTable.ForeignKeys[1].RefTable = UsersTable
	StatementsTable.ForeignKeys[0].RefTable = AssignmentsTable
}
