// Code generated by ent, DO NOT EDIT.

package program

import (
	"routine/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldCode, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldTitle, v))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v uint64) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldAuthor, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldCreatedAt, v))
}

// VersionDerivedFrom applies equality check predicate on the "version_derived_from" field. It's identical to VersionDerivedFromEQ.
func VersionDerivedFrom(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldVersionDerivedFrom, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Program {
	return predicate.Program(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Program {
	return predicate.Program(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Program {
	return predicate.Program(sql.FieldContainsFold(FieldCode, v))
}

// ProgramTypeEQ applies the EQ predicate on the "program_type" field.
func ProgramTypeEQ(v ProgramType) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldProgramType, v))
}

// ProgramTypeNEQ applies the NEQ predicate on the "program_type" field.
func ProgramTypeNEQ(v ProgramType) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldProgramType, v))
}

// ProgramTypeIn applies the In predicate on the "program_type" field.
func ProgramTypeIn(vs ...ProgramType) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldProgramType, vs...))
}

// ProgramTypeNotIn applies the NotIn predicate on the "program_type" field.
func ProgramTypeNotIn(vs ...ProgramType) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldProgramType, vs...))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Program {
	return predicate.Program(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Program {
	return predicate.Program(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Program {
	return predicate.Program(sql.FieldContainsFold(FieldTitle, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v uint64) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v uint64) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...uint64) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...uint64) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v uint64) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v uint64) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v uint64) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v uint64) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldAuthor, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldCreatedAt, v))
}

// VersionDerivedFromEQ applies the EQ predicate on the "version_derived_from" field.
func VersionDerivedFromEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldEQ(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromNEQ applies the NEQ predicate on the "version_derived_from" field.
func VersionDerivedFromNEQ(v string) predicate.Program {
	return predicate.Program(sql.FieldNEQ(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromIn applies the In predicate on the "version_derived_from" field.
func VersionDerivedFromIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldIn(FieldVersionDerivedFrom, vs...))
}

// VersionDerivedFromNotIn applies the NotIn predicate on the "version_derived_from" field.
func VersionDerivedFromNotIn(vs ...string) predicate.Program {
	return predicate.Program(sql.FieldNotIn(FieldVersionDerivedFrom, vs...))
}

// VersionDerivedFromGT applies the GT predicate on the "version_derived_from" field.
func VersionDerivedFromGT(v string) predicate.Program {
	return predicate.Program(sql.FieldGT(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromGTE applies the GTE predicate on the "version_derived_from" field.
func VersionDerivedFromGTE(v string) predicate.Program {
	return predicate.Program(sql.FieldGTE(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromLT applies the LT predicate on the "version_derived_from" field.
func VersionDerivedFromLT(v string) predicate.Program {
	return predicate.Program(sql.FieldLT(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromLTE applies the LTE predicate on the "version_derived_from" field.
func VersionDerivedFromLTE(v string) predicate.Program {
	return predicate.Program(sql.FieldLTE(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromContains applies the Contains predicate on the "version_derived_from" field.
func VersionDerivedFromContains(v string) predicate.Program {
	return predicate.Program(sql.FieldContains(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromHasPrefix applies the HasPrefix predicate on the "version_derived_from" field.
func VersionDerivedFromHasPrefix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasPrefix(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromHasSuffix applies the HasSuffix predicate on the "version_derived_from" field.
func VersionDerivedFromHasSuffix(v string) predicate.Program {
	return predicate.Program(sql.FieldHasSuffix(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromEqualFold applies the EqualFold predicate on the "version_derived_from" field.
func VersionDerivedFromEqualFold(v string) predicate.Program {
	return predicate.Program(sql.FieldEqualFold(FieldVersionDerivedFrom, v))
}

// VersionDerivedFromContainsFold applies the ContainsFold predicate on the "version_derived_from" field.
func VersionDerivedFromContainsFold(v string) predicate.Program {
	return predicate.Program(sql.FieldContainsFold(FieldVersionDerivedFrom, v))
}

// HasProgramVersions applies the HasEdge predicate on the "program_versions" edge.
func HasProgramVersions() predicate.Program {
	return predicate.Program(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProgramVersionsTable, ProgramVersionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramVersionsWith applies the HasEdge predicate on the "program_versions" edge with a given conditions (other predicates).
func HasProgramVersionsWith(preds ...predicate.ProgramVersion) predicate.Program {
	return predicate.Program(func(s *sql.Selector) {
		step := newProgramVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Program) predicate.Program {
	return predicate.Program(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Program) predicate.Program {
	return predicate.Program(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Program) predicate.Program {
	return predicate.Program(func(s *sql.Selector) {
		p(s.Not())
	})
}
