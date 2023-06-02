// Code generated by ent, DO NOT EDIT.

package onerepmax

import (
	"routine/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLTE(FieldID, id))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldAuthor, v))
}

// ActID applies equality check predicate on the "act_id" field. It's identical to ActIDEQ.
func ActID(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldActID, v))
}

// ProgramRecID applies equality check predicate on the "program_rec_id" field. It's identical to ProgramRecIDEQ.
func ProgramRecID(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldProgramRecID, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldDate, v))
}

// OneRepMax applies equality check predicate on the "one_rep_max" field. It's identical to OneRepMaxEQ.
func OneRepMax(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldOneRepMax, v))
}

// Certified applies equality check predicate on the "certified" field. It's identical to CertifiedEQ.
func Certified(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldCertified, v))
}

// Calculated applies equality check predicate on the "calculated" field. It's identical to CalculatedEQ.
func Calculated(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldCalculated, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLTE(FieldAuthor, v))
}

// ActIDEQ applies the EQ predicate on the "act_id" field.
func ActIDEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldActID, v))
}

// ActIDNEQ applies the NEQ predicate on the "act_id" field.
func ActIDNEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldActID, v))
}

// ActIDIn applies the In predicate on the "act_id" field.
func ActIDIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldActID, vs...))
}

// ActIDNotIn applies the NotIn predicate on the "act_id" field.
func ActIDNotIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldActID, vs...))
}

// ProgramRecIDEQ applies the EQ predicate on the "program_rec_id" field.
func ProgramRecIDEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldProgramRecID, v))
}

// ProgramRecIDNEQ applies the NEQ predicate on the "program_rec_id" field.
func ProgramRecIDNEQ(v uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldProgramRecID, v))
}

// ProgramRecIDIn applies the In predicate on the "program_rec_id" field.
func ProgramRecIDIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldProgramRecID, vs...))
}

// ProgramRecIDNotIn applies the NotIn predicate on the "program_rec_id" field.
func ProgramRecIDNotIn(vs ...uint64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldProgramRecID, vs...))
}

// ProgramRecIDIsNil applies the IsNil predicate on the "program_rec_id" field.
func ProgramRecIDIsNil() predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIsNull(FieldProgramRecID))
}

// ProgramRecIDNotNil applies the NotNil predicate on the "program_rec_id" field.
func ProgramRecIDNotNil() predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotNull(FieldProgramRecID))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLTE(FieldDate, v))
}

// OneRepMaxEQ applies the EQ predicate on the "one_rep_max" field.
func OneRepMaxEQ(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldOneRepMax, v))
}

// OneRepMaxNEQ applies the NEQ predicate on the "one_rep_max" field.
func OneRepMaxNEQ(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldOneRepMax, v))
}

// OneRepMaxIn applies the In predicate on the "one_rep_max" field.
func OneRepMaxIn(vs ...float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIn(FieldOneRepMax, vs...))
}

// OneRepMaxNotIn applies the NotIn predicate on the "one_rep_max" field.
func OneRepMaxNotIn(vs ...float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotIn(FieldOneRepMax, vs...))
}

// OneRepMaxGT applies the GT predicate on the "one_rep_max" field.
func OneRepMaxGT(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGT(FieldOneRepMax, v))
}

// OneRepMaxGTE applies the GTE predicate on the "one_rep_max" field.
func OneRepMaxGTE(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldGTE(FieldOneRepMax, v))
}

// OneRepMaxLT applies the LT predicate on the "one_rep_max" field.
func OneRepMaxLT(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLT(FieldOneRepMax, v))
}

// OneRepMaxLTE applies the LTE predicate on the "one_rep_max" field.
func OneRepMaxLTE(v float64) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldLTE(FieldOneRepMax, v))
}

// OneRepMaxIsNil applies the IsNil predicate on the "one_rep_max" field.
func OneRepMaxIsNil() predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldIsNull(FieldOneRepMax))
}

// OneRepMaxNotNil applies the NotNil predicate on the "one_rep_max" field.
func OneRepMaxNotNil() predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNotNull(FieldOneRepMax))
}

// CertifiedEQ applies the EQ predicate on the "certified" field.
func CertifiedEQ(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldCertified, v))
}

// CertifiedNEQ applies the NEQ predicate on the "certified" field.
func CertifiedNEQ(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldCertified, v))
}

// CalculatedEQ applies the EQ predicate on the "calculated" field.
func CalculatedEQ(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldEQ(FieldCalculated, v))
}

// CalculatedNEQ applies the NEQ predicate on the "calculated" field.
func CalculatedNEQ(v bool) predicate.OneRepMax {
	return predicate.OneRepMax(sql.FieldNEQ(FieldCalculated, v))
}

// HasAct applies the HasEdge predicate on the "act" edge.
func HasAct() predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActTable, ActColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActWith applies the HasEdge predicate on the "act" edge with a given conditions (other predicates).
func HasActWith(preds ...predicate.Act) predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		step := newActStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProgramRec applies the HasEdge predicate on the "program_rec" edge.
func HasProgramRec() predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProgramRecTable, ProgramRecColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramRecWith applies the HasEdge predicate on the "program_rec" edge with a given conditions (other predicates).
func HasProgramRecWith(preds ...predicate.ProgramRec) predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		step := newProgramRecStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OneRepMax) predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OneRepMax) predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
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
func Not(p predicate.OneRepMax) predicate.OneRepMax {
	return predicate.OneRepMax(func(s *sql.Selector) {
		p(s.Not())
	})
}