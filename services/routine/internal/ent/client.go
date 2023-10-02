// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"routine/internal/ent/migrate"

	"routine/internal/ent/act"
	"routine/internal/ent/actimage"
	"routine/internal/ent/actversion"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Act is the client for interacting with the Act builders.
	Act *ActClient
	// ActImage is the client for interacting with the ActImage builders.
	ActImage *ActImageClient
	// ActVersion is the client for interacting with the ActVersion builders.
	ActVersion *ActVersionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Act = NewActClient(c.config)
	c.ActImage = NewActImageClient(c.config)
	c.ActVersion = NewActVersionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Act:        NewActClient(cfg),
		ActImage:   NewActImageClient(cfg),
		ActVersion: NewActVersionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Act:        NewActClient(cfg),
		ActImage:   NewActImageClient(cfg),
		ActVersion: NewActVersionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Act.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
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
	c.Act.Use(hooks...)
	c.ActImage.Use(hooks...)
	c.ActVersion.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Act.Intercept(interceptors...)
	c.ActImage.Intercept(interceptors...)
	c.ActVersion.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ActMutation:
		return c.Act.mutate(ctx, m)
	case *ActImageMutation:
		return c.ActImage.mutate(ctx, m)
	case *ActVersionMutation:
		return c.ActVersion.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ActClient is a client for the Act schema.
type ActClient struct {
	config
}

// NewActClient returns a client for the Act from the given config.
func NewActClient(c config) *ActClient {
	return &ActClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `act.Hooks(f(g(h())))`.
func (c *ActClient) Use(hooks ...Hook) {
	c.hooks.Act = append(c.hooks.Act, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `act.Intercept(f(g(h())))`.
func (c *ActClient) Intercept(interceptors ...Interceptor) {
	c.inters.Act = append(c.inters.Act, interceptors...)
}

// Create returns a builder for creating a Act entity.
func (c *ActClient) Create() *ActCreate {
	mutation := newActMutation(c.config, OpCreate)
	return &ActCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Act entities.
func (c *ActClient) CreateBulk(builders ...*ActCreate) *ActCreateBulk {
	return &ActCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Act.
func (c *ActClient) Update() *ActUpdate {
	mutation := newActMutation(c.config, OpUpdate)
	return &ActUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActClient) UpdateOne(a *Act) *ActUpdateOne {
	mutation := newActMutation(c.config, OpUpdateOne, withAct(a))
	return &ActUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActClient) UpdateOneID(id uint64) *ActUpdateOne {
	mutation := newActMutation(c.config, OpUpdateOne, withActID(id))
	return &ActUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Act.
func (c *ActClient) Delete() *ActDelete {
	mutation := newActMutation(c.config, OpDelete)
	return &ActDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ActClient) DeleteOne(a *Act) *ActDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ActClient) DeleteOneID(id uint64) *ActDeleteOne {
	builder := c.Delete().Where(act.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActDeleteOne{builder}
}

// Query returns a query builder for Act.
func (c *ActClient) Query() *ActQuery {
	return &ActQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAct},
		inters: c.Interceptors(),
	}
}

// Get returns a Act entity by its id.
func (c *ActClient) Get(ctx context.Context, id uint64) (*Act, error) {
	return c.Query().Where(act.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActClient) GetX(ctx context.Context, id uint64) *Act {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActVersions queries the act_versions edge of a Act.
func (c *ActClient) QueryActVersions(a *Act) *ActVersionQuery {
	query := (&ActVersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(act.Table, act.FieldID, id),
			sqlgraph.To(actversion.Table, actversion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, act.ActVersionsTable, act.ActVersionsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActClient) Hooks() []Hook {
	return c.hooks.Act
}

// Interceptors returns the client interceptors.
func (c *ActClient) Interceptors() []Interceptor {
	return c.inters.Act
}

func (c *ActClient) mutate(ctx context.Context, m *ActMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ActCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ActUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ActUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ActDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Act mutation op: %q", m.Op())
	}
}

// ActImageClient is a client for the ActImage schema.
type ActImageClient struct {
	config
}

// NewActImageClient returns a client for the ActImage from the given config.
func NewActImageClient(c config) *ActImageClient {
	return &ActImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `actimage.Hooks(f(g(h())))`.
func (c *ActImageClient) Use(hooks ...Hook) {
	c.hooks.ActImage = append(c.hooks.ActImage, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `actimage.Intercept(f(g(h())))`.
func (c *ActImageClient) Intercept(interceptors ...Interceptor) {
	c.inters.ActImage = append(c.inters.ActImage, interceptors...)
}

// Create returns a builder for creating a ActImage entity.
func (c *ActImageClient) Create() *ActImageCreate {
	mutation := newActImageMutation(c.config, OpCreate)
	return &ActImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ActImage entities.
func (c *ActImageClient) CreateBulk(builders ...*ActImageCreate) *ActImageCreateBulk {
	return &ActImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ActImage.
func (c *ActImageClient) Update() *ActImageUpdate {
	mutation := newActImageMutation(c.config, OpUpdate)
	return &ActImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActImageClient) UpdateOne(ai *ActImage) *ActImageUpdateOne {
	mutation := newActImageMutation(c.config, OpUpdateOne, withActImage(ai))
	return &ActImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActImageClient) UpdateOneID(id uint64) *ActImageUpdateOne {
	mutation := newActImageMutation(c.config, OpUpdateOne, withActImageID(id))
	return &ActImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ActImage.
func (c *ActImageClient) Delete() *ActImageDelete {
	mutation := newActImageMutation(c.config, OpDelete)
	return &ActImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ActImageClient) DeleteOne(ai *ActImage) *ActImageDeleteOne {
	return c.DeleteOneID(ai.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ActImageClient) DeleteOneID(id uint64) *ActImageDeleteOne {
	builder := c.Delete().Where(actimage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActImageDeleteOne{builder}
}

// Query returns a query builder for ActImage.
func (c *ActImageClient) Query() *ActImageQuery {
	return &ActImageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeActImage},
		inters: c.Interceptors(),
	}
}

// Get returns a ActImage entity by its id.
func (c *ActImageClient) Get(ctx context.Context, id uint64) (*ActImage, error) {
	return c.Query().Where(actimage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActImageClient) GetX(ctx context.Context, id uint64) *ActImage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActVersion queries the act_version edge of a ActImage.
func (c *ActImageClient) QueryActVersion(ai *ActImage) *ActVersionQuery {
	query := (&ActVersionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ai.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actimage.Table, actimage.FieldID, id),
			sqlgraph.To(actversion.Table, actversion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, actimage.ActVersionTable, actimage.ActVersionColumn),
		)
		fromV = sqlgraph.Neighbors(ai.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActImageClient) Hooks() []Hook {
	return c.hooks.ActImage
}

// Interceptors returns the client interceptors.
func (c *ActImageClient) Interceptors() []Interceptor {
	return c.inters.ActImage
}

func (c *ActImageClient) mutate(ctx context.Context, m *ActImageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ActImageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ActImageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ActImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ActImageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ActImage mutation op: %q", m.Op())
	}
}

// ActVersionClient is a client for the ActVersion schema.
type ActVersionClient struct {
	config
}

// NewActVersionClient returns a client for the ActVersion from the given config.
func NewActVersionClient(c config) *ActVersionClient {
	return &ActVersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `actversion.Hooks(f(g(h())))`.
func (c *ActVersionClient) Use(hooks ...Hook) {
	c.hooks.ActVersion = append(c.hooks.ActVersion, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `actversion.Intercept(f(g(h())))`.
func (c *ActVersionClient) Intercept(interceptors ...Interceptor) {
	c.inters.ActVersion = append(c.inters.ActVersion, interceptors...)
}

// Create returns a builder for creating a ActVersion entity.
func (c *ActVersionClient) Create() *ActVersionCreate {
	mutation := newActVersionMutation(c.config, OpCreate)
	return &ActVersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ActVersion entities.
func (c *ActVersionClient) CreateBulk(builders ...*ActVersionCreate) *ActVersionCreateBulk {
	return &ActVersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ActVersion.
func (c *ActVersionClient) Update() *ActVersionUpdate {
	mutation := newActVersionMutation(c.config, OpUpdate)
	return &ActVersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActVersionClient) UpdateOne(av *ActVersion) *ActVersionUpdateOne {
	mutation := newActVersionMutation(c.config, OpUpdateOne, withActVersion(av))
	return &ActVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActVersionClient) UpdateOneID(id uint64) *ActVersionUpdateOne {
	mutation := newActVersionMutation(c.config, OpUpdateOne, withActVersionID(id))
	return &ActVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ActVersion.
func (c *ActVersionClient) Delete() *ActVersionDelete {
	mutation := newActVersionMutation(c.config, OpDelete)
	return &ActVersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ActVersionClient) DeleteOne(av *ActVersion) *ActVersionDeleteOne {
	return c.DeleteOneID(av.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ActVersionClient) DeleteOneID(id uint64) *ActVersionDeleteOne {
	builder := c.Delete().Where(actversion.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActVersionDeleteOne{builder}
}

// Query returns a query builder for ActVersion.
func (c *ActVersionClient) Query() *ActVersionQuery {
	return &ActVersionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeActVersion},
		inters: c.Interceptors(),
	}
}

// Get returns a ActVersion entity by its id.
func (c *ActVersionClient) Get(ctx context.Context, id uint64) (*ActVersion, error) {
	return c.Query().Where(actversion.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActVersionClient) GetX(ctx context.Context, id uint64) *ActVersion {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryActImages queries the act_images edge of a ActVersion.
func (c *ActVersionClient) QueryActImages(av *ActVersion) *ActImageQuery {
	query := (&ActImageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := av.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actversion.Table, actversion.FieldID, id),
			sqlgraph.To(actimage.Table, actimage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, actversion.ActImagesTable, actversion.ActImagesColumn),
		)
		fromV = sqlgraph.Neighbors(av.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAct queries the act edge of a ActVersion.
func (c *ActVersionClient) QueryAct(av *ActVersion) *ActQuery {
	query := (&ActClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := av.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(actversion.Table, actversion.FieldID, id),
			sqlgraph.To(act.Table, act.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, actversion.ActTable, actversion.ActColumn),
		)
		fromV = sqlgraph.Neighbors(av.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActVersionClient) Hooks() []Hook {
	return c.hooks.ActVersion
}

// Interceptors returns the client interceptors.
func (c *ActVersionClient) Interceptors() []Interceptor {
	return c.inters.ActVersion
}

func (c *ActVersionClient) mutate(ctx context.Context, m *ActVersionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ActVersionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ActVersionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ActVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ActVersionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ActVersion mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Act, ActImage, ActVersion []ent.Hook
	}
	inters struct {
		Act, ActImage, ActVersion []ent.Interceptor
	}
)
