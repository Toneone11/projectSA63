// Code generated by entc, DO NOT EDIT.

package operationroom

const (
	// Label holds the string label denoting the operationroom type in the database.
	Label = "operationroom"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOperationroomName holds the string denoting the operationroom_name field in the database.
	FieldOperationroomName = "operationroom_name"

	// EdgeOperationroomID holds the string denoting the operationroom_id edge name in mutations.
	EdgeOperationroomID = "operationroom_id"

	// Table holds the table name of the operationroom in the database.
	Table = "operationrooms"
	// OperationroomIDTable is the table the holds the operationroom_id relation/edge.
	OperationroomIDTable = "bookings"
	// OperationroomIDInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	OperationroomIDInverseTable = "bookings"
	// OperationroomIDColumn is the table column denoting the operationroom_id relation/edge.
	OperationroomIDColumn = "operationroom_operationroom_id"
)

// Columns holds all SQL columns for operationroom fields.
var Columns = []string{
	FieldID,
	FieldOperationroomName,
}

var (
	// OperationroomNameValidator is a validator for the "operationroom_name" field. It is called by the builders before save.
	OperationroomNameValidator func(string) error
)
