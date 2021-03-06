// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Toneone11/app/ent/migrate"

	"github.com/Toneone11/app/ent/booking"
	"github.com/Toneone11/app/ent/operationroom"
	"github.com/Toneone11/app/ent/patient"
	"github.com/Toneone11/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Booking is the client for interacting with the Booking builders.
	Booking *BookingClient
	// Operationroom is the client for interacting with the Operationroom builders.
	Operationroom *OperationroomClient
	// Patient is the client for interacting with the Patient builders.
	Patient *PatientClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Booking = NewBookingClient(c.config)
	c.Operationroom = NewOperationroomClient(c.config)
	c.Patient = NewPatientClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Booking:       NewBookingClient(cfg),
		Operationroom: NewOperationroomClient(cfg),
		Patient:       NewPatientClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:        cfg,
		Booking:       NewBookingClient(cfg),
		Operationroom: NewOperationroomClient(cfg),
		Patient:       NewPatientClient(cfg),
		User:          NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Booking.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Booking.Use(hooks...)
	c.Operationroom.Use(hooks...)
	c.Patient.Use(hooks...)
	c.User.Use(hooks...)
}

// BookingClient is a client for the Booking schema.
type BookingClient struct {
	config
}

// NewBookingClient returns a client for the Booking from the given config.
func NewBookingClient(c config) *BookingClient {
	return &BookingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `booking.Hooks(f(g(h())))`.
func (c *BookingClient) Use(hooks ...Hook) {
	c.hooks.Booking = append(c.hooks.Booking, hooks...)
}

// Create returns a create builder for Booking.
func (c *BookingClient) Create() *BookingCreate {
	mutation := newBookingMutation(c.config, OpCreate)
	return &BookingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Booking.
func (c *BookingClient) Update() *BookingUpdate {
	mutation := newBookingMutation(c.config, OpUpdate)
	return &BookingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookingClient) UpdateOne(b *Booking) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBooking(b))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookingClient) UpdateOneID(id int) *BookingUpdateOne {
	mutation := newBookingMutation(c.config, OpUpdateOne, withBookingID(id))
	return &BookingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Booking.
func (c *BookingClient) Delete() *BookingDelete {
	mutation := newBookingMutation(c.config, OpDelete)
	return &BookingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BookingClient) DeleteOne(b *Booking) *BookingDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BookingClient) DeleteOneID(id int) *BookingDeleteOne {
	builder := c.Delete().Where(booking.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookingDeleteOne{builder}
}

// Create returns a query builder for Booking.
func (c *BookingClient) Query() *BookingQuery {
	return &BookingQuery{config: c.config}
}

// Get returns a Booking entity by its id.
func (c *BookingClient) Get(ctx context.Context, id int) (*Booking, error) {
	return c.Query().Where(booking.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookingClient) GetX(ctx context.Context, id int) *Booking {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QueryDoctorID queries the doctor_id edge of a Booking.
func (c *BookingClient) QueryDoctorID(b *Booking) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.DoctorIDTable, booking.DoctorIDColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPatientID queries the patient_id edge of a Booking.
func (c *BookingClient) QueryPatientID(b *Booking) *PatientQuery {
	query := &PatientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.PatientIDTable, booking.PatientIDColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOperationroomID queries the operationroom_id edge of a Booking.
func (c *BookingClient) QueryOperationroomID(b *Booking) *OperationroomQuery {
	query := &OperationroomQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(booking.Table, booking.FieldID, id),
			sqlgraph.To(operationroom.Table, operationroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, booking.OperationroomIDTable, booking.OperationroomIDColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookingClient) Hooks() []Hook {
	return c.hooks.Booking
}

// OperationroomClient is a client for the Operationroom schema.
type OperationroomClient struct {
	config
}

// NewOperationroomClient returns a client for the Operationroom from the given config.
func NewOperationroomClient(c config) *OperationroomClient {
	return &OperationroomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `operationroom.Hooks(f(g(h())))`.
func (c *OperationroomClient) Use(hooks ...Hook) {
	c.hooks.Operationroom = append(c.hooks.Operationroom, hooks...)
}

// Create returns a create builder for Operationroom.
func (c *OperationroomClient) Create() *OperationroomCreate {
	mutation := newOperationroomMutation(c.config, OpCreate)
	return &OperationroomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Operationroom.
func (c *OperationroomClient) Update() *OperationroomUpdate {
	mutation := newOperationroomMutation(c.config, OpUpdate)
	return &OperationroomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OperationroomClient) UpdateOne(o *Operationroom) *OperationroomUpdateOne {
	mutation := newOperationroomMutation(c.config, OpUpdateOne, withOperationroom(o))
	return &OperationroomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OperationroomClient) UpdateOneID(id int) *OperationroomUpdateOne {
	mutation := newOperationroomMutation(c.config, OpUpdateOne, withOperationroomID(id))
	return &OperationroomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Operationroom.
func (c *OperationroomClient) Delete() *OperationroomDelete {
	mutation := newOperationroomMutation(c.config, OpDelete)
	return &OperationroomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *OperationroomClient) DeleteOne(o *Operationroom) *OperationroomDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *OperationroomClient) DeleteOneID(id int) *OperationroomDeleteOne {
	builder := c.Delete().Where(operationroom.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OperationroomDeleteOne{builder}
}

// Create returns a query builder for Operationroom.
func (c *OperationroomClient) Query() *OperationroomQuery {
	return &OperationroomQuery{config: c.config}
}

// Get returns a Operationroom entity by its id.
func (c *OperationroomClient) Get(ctx context.Context, id int) (*Operationroom, error) {
	return c.Query().Where(operationroom.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OperationroomClient) GetX(ctx context.Context, id int) *Operationroom {
	o, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return o
}

// QueryOperationroomID queries the operationroom_id edge of a Operationroom.
func (c *OperationroomClient) QueryOperationroomID(o *Operationroom) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(operationroom.Table, operationroom.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, operationroom.OperationroomIDTable, operationroom.OperationroomIDColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OperationroomClient) Hooks() []Hook {
	return c.hooks.Operationroom
}

// PatientClient is a client for the Patient schema.
type PatientClient struct {
	config
}

// NewPatientClient returns a client for the Patient from the given config.
func NewPatientClient(c config) *PatientClient {
	return &PatientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `patient.Hooks(f(g(h())))`.
func (c *PatientClient) Use(hooks ...Hook) {
	c.hooks.Patient = append(c.hooks.Patient, hooks...)
}

// Create returns a create builder for Patient.
func (c *PatientClient) Create() *PatientCreate {
	mutation := newPatientMutation(c.config, OpCreate)
	return &PatientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Patient.
func (c *PatientClient) Update() *PatientUpdate {
	mutation := newPatientMutation(c.config, OpUpdate)
	return &PatientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PatientClient) UpdateOne(pa *Patient) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatient(pa))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PatientClient) UpdateOneID(id int) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatientID(id))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Patient.
func (c *PatientClient) Delete() *PatientDelete {
	mutation := newPatientMutation(c.config, OpDelete)
	return &PatientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PatientClient) DeleteOne(pa *Patient) *PatientDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PatientClient) DeleteOneID(id int) *PatientDeleteOne {
	builder := c.Delete().Where(patient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PatientDeleteOne{builder}
}

// Create returns a query builder for Patient.
func (c *PatientClient) Query() *PatientQuery {
	return &PatientQuery{config: c.config}
}

// Get returns a Patient entity by its id.
func (c *PatientClient) Get(ctx context.Context, id int) (*Patient, error) {
	return c.Query().Where(patient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PatientClient) GetX(ctx context.Context, id int) *Patient {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QueryPatientID queries the patient_id edge of a Patient.
func (c *PatientClient) QueryPatientID(pa *Patient) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patient.PatientIDTable, patient.PatientIDColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PatientClient) Hooks() []Hook {
	return c.hooks.Patient
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryDoctorID queries the doctor_id edge of a User.
func (c *UserClient) QueryDoctorID(u *User) *BookingQuery {
	query := &BookingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DoctorIDTable, user.DoctorIDColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
