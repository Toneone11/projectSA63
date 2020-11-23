// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/Toneone11/app/ent/booking"
	"github.com/Toneone11/app/ent/operationroom"
	"github.com/Toneone11/app/ent/patient"
	"github.com/Toneone11/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Booking is the model entity for the Booking schema.
type Booking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingQuery when eager-loading is set.
	Edges                          BookingEdges `json:"edges"`
	operationroom_operationroom_id *int
	patient_patient_id             *int
	user_doctor_id                 *int
}

// BookingEdges holds the relations/edges for other nodes in the graph.
type BookingEdges struct {
	// DoctorID holds the value of the doctor_id edge.
	DoctorID *User
	// PatientID holds the value of the patient_id edge.
	PatientID *Patient
	// OperationroomID holds the value of the operationroom_id edge.
	OperationroomID *Operationroom
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DoctorIDOrErr returns the DoctorID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) DoctorIDOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.DoctorID == nil {
			// The edge doctor_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.DoctorID, nil
	}
	return nil, &NotLoadedError{edge: "doctor_id"}
}

// PatientIDOrErr returns the PatientID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) PatientIDOrErr() (*Patient, error) {
	if e.loadedTypes[1] {
		if e.PatientID == nil {
			// The edge patient_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patient.Label}
		}
		return e.PatientID, nil
	}
	return nil, &NotLoadedError{edge: "patient_id"}
}

// OperationroomIDOrErr returns the OperationroomID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookingEdges) OperationroomIDOrErr() (*Operationroom, error) {
	if e.loadedTypes[2] {
		if e.OperationroomID == nil {
			// The edge operationroom_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: operationroom.Label}
		}
		return e.OperationroomID, nil
	}
	return nil, &NotLoadedError{edge: "operationroom_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Booking) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // date
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Booking) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // operationroom_operationroom_id
		&sql.NullInt64{}, // patient_patient_id
		&sql.NullInt64{}, // user_doctor_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Booking fields.
func (b *Booking) assignValues(values ...interface{}) error {
	if m, n := len(values), len(booking.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field date", values[0])
	} else if value.Valid {
		b.Date = value.Time
	}
	values = values[1:]
	if len(values) == len(booking.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field operationroom_operationroom_id", value)
		} else if value.Valid {
			b.operationroom_operationroom_id = new(int)
			*b.operationroom_operationroom_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patient_patient_id", value)
		} else if value.Valid {
			b.patient_patient_id = new(int)
			*b.patient_patient_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_doctor_id", value)
		} else if value.Valid {
			b.user_doctor_id = new(int)
			*b.user_doctor_id = int(value.Int64)
		}
	}
	return nil
}

// QueryDoctorID queries the doctor_id edge of the Booking.
func (b *Booking) QueryDoctorID() *UserQuery {
	return (&BookingClient{config: b.config}).QueryDoctorID(b)
}

// QueryPatientID queries the patient_id edge of the Booking.
func (b *Booking) QueryPatientID() *PatientQuery {
	return (&BookingClient{config: b.config}).QueryPatientID(b)
}

// QueryOperationroomID queries the operationroom_id edge of the Booking.
func (b *Booking) QueryOperationroomID() *OperationroomQuery {
	return (&BookingClient{config: b.config}).QueryOperationroomID(b)
}

// Update returns a builder for updating this Booking.
// Note that, you need to call Booking.Unwrap() before calling this method, if this Booking
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Booking) Update() *BookingUpdateOne {
	return (&BookingClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Booking) Unwrap() *Booking {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Booking is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Booking) String() string {
	var builder strings.Builder
	builder.WriteString("Booking(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", date=")
	builder.WriteString(b.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Bookings is a parsable slice of Booking.
type Bookings []*Booking

func (b Bookings) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}