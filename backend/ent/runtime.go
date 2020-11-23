// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/Toneone11/app/ent/operationroom"
	"github.com/Toneone11/app/ent/patient"
	"github.com/Toneone11/app/ent/schema"
	"github.com/Toneone11/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	operationroomFields := schema.Operationroom{}.Fields()
	_ = operationroomFields
	// operationroomDescOperationroomName is the schema descriptor for operationroom_name field.
	operationroomDescOperationroomName := operationroomFields[0].Descriptor()
	// operationroom.OperationroomNameValidator is a validator for the "operationroom_name" field. It is called by the builders before save.
	operationroom.OperationroomNameValidator = operationroomDescOperationroomName.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for patient_name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDoctorName is the schema descriptor for doctor_name field.
	userDescDoctorName := userFields[0].Descriptor()
	// user.DoctorNameValidator is a validator for the "doctor_name" field. It is called by the builders before save.
	user.DoctorNameValidator = userDescDoctorName.Validators[0].(func(string) error)
	// userDescDoctorEmail is the schema descriptor for doctor_email field.
	userDescDoctorEmail := userFields[1].Descriptor()
	// user.DoctorEmailValidator is a validator for the "doctor_email" field. It is called by the builders before save.
	user.DoctorEmailValidator = userDescDoctorEmail.Validators[0].(func(string) error)
}
