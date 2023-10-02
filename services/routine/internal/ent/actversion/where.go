// Code generated by ent, DO NOT EDIT.

package actversion

import (
	"routine/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldCode, v))
}

// ActCode applies equality check predicate on the "act_code" field. It's identical to ActCodeEQ.
func ActCode(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldActCode, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldVersion, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldContainsFold(FieldCode, v))
}

// ActCodeEQ applies the EQ predicate on the "act_code" field.
func ActCodeEQ(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldActCode, v))
}

// ActCodeNEQ applies the NEQ predicate on the "act_code" field.
func ActCodeNEQ(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNEQ(FieldActCode, v))
}

// ActCodeIn applies the In predicate on the "act_code" field.
func ActCodeIn(vs ...string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldIn(FieldActCode, vs...))
}

// ActCodeNotIn applies the NotIn predicate on the "act_code" field.
func ActCodeNotIn(vs ...string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNotIn(FieldActCode, vs...))
}

// ActCodeGT applies the GT predicate on the "act_code" field.
func ActCodeGT(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGT(FieldActCode, v))
}

// ActCodeGTE applies the GTE predicate on the "act_code" field.
func ActCodeGTE(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGTE(FieldActCode, v))
}

// ActCodeLT applies the LT predicate on the "act_code" field.
func ActCodeLT(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLT(FieldActCode, v))
}

// ActCodeLTE applies the LTE predicate on the "act_code" field.
func ActCodeLTE(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLTE(FieldActCode, v))
}

// ActCodeContains applies the Contains predicate on the "act_code" field.
func ActCodeContains(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldContains(FieldActCode, v))
}

// ActCodeHasPrefix applies the HasPrefix predicate on the "act_code" field.
func ActCodeHasPrefix(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldHasPrefix(FieldActCode, v))
}

// ActCodeHasSuffix applies the HasSuffix predicate on the "act_code" field.
func ActCodeHasSuffix(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldHasSuffix(FieldActCode, v))
}

// ActCodeEqualFold applies the EqualFold predicate on the "act_code" field.
func ActCodeEqualFold(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEqualFold(FieldActCode, v))
}

// ActCodeContainsFold applies the ContainsFold predicate on the "act_code" field.
func ActCodeContainsFold(v string) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldContainsFold(FieldActCode, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v uint) predicate.ActVersion {
	return predicate.ActVersion(sql.FieldLTE(FieldVersion, v))
}

// HasActImages applies the HasEdge predicate on the "act_images" edge.
func HasActImages() predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ActImagesTable, ActImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActImagesWith applies the HasEdge predicate on the "act_images" edge with a given conditions (other predicates).
func HasActImagesWith(preds ...predicate.ActImage) predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		step := newActImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAct applies the HasEdge predicate on the "act" edge.
func HasAct() predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActTable, ActColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActWith applies the HasEdge predicate on the "act" edge with a given conditions (other predicates).
func HasActWith(preds ...predicate.Act) predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		step := newActStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ActVersion) predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ActVersion) predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
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
func Not(p predicate.ActVersion) predicate.ActVersion {
	return predicate.ActVersion(func(s *sql.Selector) {
		p(s.Not())
	})
}