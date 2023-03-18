// Code generated by ent, DO NOT EDIT.

package ent

import (
	"lifthus-auth/ent/schema"
	"lifthus-auth/ent/session"
	"lifthus-auth/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescConnectedAt is the schema descriptor for connected_at field.
	sessionDescConnectedAt := sessionFields[2].Descriptor()
	// session.DefaultConnectedAt holds the default value on creation for the connected_at field.
	session.DefaultConnectedAt = sessionDescConnectedAt.Default.(func() time.Time)
	// sessionDescSignedAt is the schema descriptor for signed_at field.
	sessionDescSignedAt := sessionFields[3].Descriptor()
	// session.UpdateDefaultSignedAt holds the default value on update for the signed_at field.
	session.UpdateDefaultSignedAt = sessionDescSignedAt.UpdateDefault.(func() time.Time)
	// sessionDescUsed is the schema descriptor for used field.
	sessionDescUsed := sessionFields[4].Descriptor()
	// session.DefaultUsed holds the default value on creation for the used field.
	session.DefaultUsed = sessionDescUsed.Default.(bool)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionFields[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRegistered is the schema descriptor for registered field.
	userDescRegistered := userFields[1].Descriptor()
	// user.DefaultRegistered holds the default value on creation for the registered field.
	user.DefaultRegistered = userDescRegistered.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[11].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[12].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
