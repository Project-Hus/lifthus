// Code generated by ent, DO NOT EDIT.

package ent

import (
	"routine/internal/ent/act"
	"routine/internal/ent/actimage"
	"routine/internal/ent/actversion"
	"routine/internal/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	actFields := schema.Act{}.Fields()
	_ = actFields
	// actDescCode is the schema descriptor for code field.
	actDescCode := actFields[1].Descriptor()
	// act.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	act.CodeValidator = actDescCode.Validators[0].(func(string) error)
	// actDescName is the schema descriptor for name field.
	actDescName := actFields[3].Descriptor()
	// act.NameValidator is a validator for the "name" field. It is called by the builders before save.
	act.NameValidator = actDescName.Validators[0].(func(string) error)
	actimageFields := schema.ActImage{}.Fields()
	_ = actimageFields
	// actimageDescSrc is the schema descriptor for src field.
	actimageDescSrc := actimageFields[2].Descriptor()
	// actimage.SrcValidator is a validator for the "src" field. It is called by the builders before save.
	actimage.SrcValidator = actimageDescSrc.Validators[0].(func(string) error)
	actversionFields := schema.ActVersion{}.Fields()
	_ = actversionFields
	// actversionDescCode is the schema descriptor for code field.
	actversionDescCode := actversionFields[1].Descriptor()
	// actversion.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	actversion.CodeValidator = actversionDescCode.Validators[0].(func(string) error)
	// actversionDescActCode is the schema descriptor for act_code field.
	actversionDescActCode := actversionFields[2].Descriptor()
	// actversion.ActCodeValidator is a validator for the "act_code" field. It is called by the builders before save.
	actversion.ActCodeValidator = actversionDescActCode.Validators[0].(func(string) error)
}