// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"lifthus-auth/ent/migrate"

	"lifthus-auth/ent/lifthusgroup"
	"lifthus-auth/ent/refreshtoken"
	"lifthus-auth/ent/session"
	"lifthus-auth/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LifthusGroup is the client for interacting with the LifthusGroup builders.
	LifthusGroup *LifthusGroupClient
	// RefreshToken is the client for interacting with the RefreshToken builders.
	RefreshToken *RefreshTokenClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.LifthusGroup = NewLifthusGroupClient(c.config)
	c.RefreshToken = NewRefreshTokenClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		LifthusGroup: NewLifthusGroupClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		Session:      NewSessionClient(cfg),
		User:         NewUserClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		LifthusGroup: NewLifthusGroupClient(cfg),
		RefreshToken: NewRefreshTokenClient(cfg),
		Session:      NewSessionClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LifthusGroup.
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
	c.LifthusGroup.Use(hooks...)
	c.RefreshToken.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.LifthusGroup.Intercept(interceptors...)
	c.RefreshToken.Intercept(interceptors...)
	c.Session.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LifthusGroupMutation:
		return c.LifthusGroup.mutate(ctx, m)
	case *RefreshTokenMutation:
		return c.RefreshToken.mutate(ctx, m)
	case *SessionMutation:
		return c.Session.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LifthusGroupClient is a client for the LifthusGroup schema.
type LifthusGroupClient struct {
	config
}

// NewLifthusGroupClient returns a client for the LifthusGroup from the given config.
func NewLifthusGroupClient(c config) *LifthusGroupClient {
	return &LifthusGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `lifthusgroup.Hooks(f(g(h())))`.
func (c *LifthusGroupClient) Use(hooks ...Hook) {
	c.hooks.LifthusGroup = append(c.hooks.LifthusGroup, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `lifthusgroup.Intercept(f(g(h())))`.
func (c *LifthusGroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.LifthusGroup = append(c.inters.LifthusGroup, interceptors...)
}

// Create returns a builder for creating a LifthusGroup entity.
func (c *LifthusGroupClient) Create() *LifthusGroupCreate {
	mutation := newLifthusGroupMutation(c.config, OpCreate)
	return &LifthusGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LifthusGroup entities.
func (c *LifthusGroupClient) CreateBulk(builders ...*LifthusGroupCreate) *LifthusGroupCreateBulk {
	return &LifthusGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LifthusGroup.
func (c *LifthusGroupClient) Update() *LifthusGroupUpdate {
	mutation := newLifthusGroupMutation(c.config, OpUpdate)
	return &LifthusGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LifthusGroupClient) UpdateOne(lg *LifthusGroup) *LifthusGroupUpdateOne {
	mutation := newLifthusGroupMutation(c.config, OpUpdateOne, withLifthusGroup(lg))
	return &LifthusGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LifthusGroupClient) UpdateOneID(id int) *LifthusGroupUpdateOne {
	mutation := newLifthusGroupMutation(c.config, OpUpdateOne, withLifthusGroupID(id))
	return &LifthusGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LifthusGroup.
func (c *LifthusGroupClient) Delete() *LifthusGroupDelete {
	mutation := newLifthusGroupMutation(c.config, OpDelete)
	return &LifthusGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LifthusGroupClient) DeleteOne(lg *LifthusGroup) *LifthusGroupDeleteOne {
	return c.DeleteOneID(lg.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LifthusGroupClient) DeleteOneID(id int) *LifthusGroupDeleteOne {
	builder := c.Delete().Where(lifthusgroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LifthusGroupDeleteOne{builder}
}

// Query returns a query builder for LifthusGroup.
func (c *LifthusGroupClient) Query() *LifthusGroupQuery {
	return &LifthusGroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLifthusGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a LifthusGroup entity by its id.
func (c *LifthusGroupClient) Get(ctx context.Context, id int) (*LifthusGroup, error) {
	return c.Query().Where(lifthusgroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LifthusGroupClient) GetX(ctx context.Context, id int) *LifthusGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LifthusGroupClient) Hooks() []Hook {
	return c.hooks.LifthusGroup
}

// Interceptors returns the client interceptors.
func (c *LifthusGroupClient) Interceptors() []Interceptor {
	return c.inters.LifthusGroup
}

func (c *LifthusGroupClient) mutate(ctx context.Context, m *LifthusGroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LifthusGroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LifthusGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LifthusGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LifthusGroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LifthusGroup mutation op: %q", m.Op())
	}
}

// RefreshTokenClient is a client for the RefreshToken schema.
type RefreshTokenClient struct {
	config
}

// NewRefreshTokenClient returns a client for the RefreshToken from the given config.
func NewRefreshTokenClient(c config) *RefreshTokenClient {
	return &RefreshTokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `refreshtoken.Hooks(f(g(h())))`.
func (c *RefreshTokenClient) Use(hooks ...Hook) {
	c.hooks.RefreshToken = append(c.hooks.RefreshToken, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `refreshtoken.Intercept(f(g(h())))`.
func (c *RefreshTokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.RefreshToken = append(c.inters.RefreshToken, interceptors...)
}

// Create returns a builder for creating a RefreshToken entity.
func (c *RefreshTokenClient) Create() *RefreshTokenCreate {
	mutation := newRefreshTokenMutation(c.config, OpCreate)
	return &RefreshTokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RefreshToken entities.
func (c *RefreshTokenClient) CreateBulk(builders ...*RefreshTokenCreate) *RefreshTokenCreateBulk {
	return &RefreshTokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RefreshToken.
func (c *RefreshTokenClient) Update() *RefreshTokenUpdate {
	mutation := newRefreshTokenMutation(c.config, OpUpdate)
	return &RefreshTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RefreshTokenClient) UpdateOne(rt *RefreshToken) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshToken(rt))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RefreshTokenClient) UpdateOneID(id uuid.UUID) *RefreshTokenUpdateOne {
	mutation := newRefreshTokenMutation(c.config, OpUpdateOne, withRefreshTokenID(id))
	return &RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RefreshToken.
func (c *RefreshTokenClient) Delete() *RefreshTokenDelete {
	mutation := newRefreshTokenMutation(c.config, OpDelete)
	return &RefreshTokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RefreshTokenClient) DeleteOne(rt *RefreshToken) *RefreshTokenDeleteOne {
	return c.DeleteOneID(rt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RefreshTokenClient) DeleteOneID(id uuid.UUID) *RefreshTokenDeleteOne {
	builder := c.Delete().Where(refreshtoken.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RefreshTokenDeleteOne{builder}
}

// Query returns a query builder for RefreshToken.
func (c *RefreshTokenClient) Query() *RefreshTokenQuery {
	return &RefreshTokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRefreshToken},
		inters: c.Interceptors(),
	}
}

// Get returns a RefreshToken entity by its id.
func (c *RefreshTokenClient) Get(ctx context.Context, id uuid.UUID) (*RefreshToken, error) {
	return c.Query().Where(refreshtoken.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RefreshTokenClient) GetX(ctx context.Context, id uuid.UUID) *RefreshToken {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a RefreshToken.
func (c *RefreshTokenClient) QueryUser(rt *RefreshToken) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(refreshtoken.Table, refreshtoken.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, refreshtoken.UserTable, refreshtoken.UserColumn),
		)
		fromV = sqlgraph.Neighbors(rt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RefreshTokenClient) Hooks() []Hook {
	return c.hooks.RefreshToken
}

// Interceptors returns the client interceptors.
func (c *RefreshTokenClient) Interceptors() []Interceptor {
	return c.inters.RefreshToken
}

func (c *RefreshTokenClient) mutate(ctx context.Context, m *RefreshTokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RefreshTokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RefreshTokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RefreshTokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RefreshTokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RefreshToken mutation op: %q", m.Op())
	}
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `session.Intercept(f(g(h())))`.
func (c *SessionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Session = append(c.inters.Session, interceptors...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id uuid.UUID) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id uuid.UUID) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSession},
		inters: c.Interceptors(),
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id uuid.UUID) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id uuid.UUID) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Session.
func (c *SessionClient) QueryUser(s *Session) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, session.UserTable, session.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// Interceptors returns the client interceptors.
func (c *SessionClient) Interceptors() []Interceptor {
	return c.inters.Session
}

func (c *SessionClient) mutate(ctx context.Context, m *SessionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SessionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SessionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Session mutation op: %q", m.Op())
	}
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

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySessions queries the sessions edge of a User.
func (c *UserClient) QuerySessions(u *User) *SessionQuery {
	query := (&SessionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.SessionsTable, user.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLifthusTokens queries the lifthus_tokens edge of a User.
func (c *UserClient) QueryLifthusTokens(u *User) *RefreshTokenQuery {
	query := (&RefreshTokenClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(refreshtoken.Table, refreshtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.LifthusTokensTable, user.LifthusTokensColumn),
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

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		LifthusGroup, RefreshToken, Session, User []ent.Hook
	}
	inters struct {
		LifthusGroup, RefreshToken, Session, User []ent.Interceptor
	}
)
