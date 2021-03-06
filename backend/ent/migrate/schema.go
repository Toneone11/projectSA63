// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BookingsColumns holds the columns for the "bookings" table.
	BookingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "operationroom_operationroom_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_doctor_id", Type: field.TypeInt, Nullable: true},
	}
	// BookingsTable holds the schema information for the "bookings" table.
	BookingsTable = &schema.Table{
		Name:       "bookings",
		Columns:    BookingsColumns,
		PrimaryKey: []*schema.Column{BookingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bookings_operationrooms_operationroom_id",
				Columns: []*schema.Column{BookingsColumns[2]},

				RefColumns: []*schema.Column{OperationroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_patients_patient_id",
				Columns: []*schema.Column{BookingsColumns[3]},

				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bookings_users_doctor_id",
				Columns: []*schema.Column{BookingsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OperationroomsColumns holds the columns for the "operationrooms" table.
	OperationroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "operationroom_name", Type: field.TypeString},
	}
	// OperationroomsTable holds the schema information for the "operationrooms" table.
	OperationroomsTable = &schema.Table{
		Name:        "operationrooms",
		Columns:     OperationroomsColumns,
		PrimaryKey:  []*schema.Column{OperationroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "patient_name", Type: field.TypeString},
		{Name: "patient_age", Type: field.TypeInt},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:        "patients",
		Columns:     PatientsColumns,
		PrimaryKey:  []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doctor_name", Type: field.TypeString},
		{Name: "doctor_email", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BookingsTable,
		OperationroomsTable,
		PatientsTable,
		UsersTable,
	}
)

func init() {
	BookingsTable.ForeignKeys[0].RefTable = OperationroomsTable
	BookingsTable.ForeignKeys[1].RefTable = PatientsTable
	BookingsTable.ForeignKeys[2].RefTable = UsersTable
}
