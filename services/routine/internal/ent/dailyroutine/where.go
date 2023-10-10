// Code generated by ent, DO NOT EDIT.

package dailyroutine

import (
	"routine/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldCode, v))
}

// ProgramVersionCode applies equality check predicate on the "program_version_code" field. It's identical to ProgramVersionCodeEQ.
func ProgramVersionCode(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldProgramVersionCode, v))
}

// Day applies equality check predicate on the "day" field. It's identical to DayEQ.
func Day(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldDay, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldContainsFold(FieldCode, v))
}

// ProgramVersionCodeEQ applies the EQ predicate on the "program_version_code" field.
func ProgramVersionCodeEQ(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldProgramVersionCode, v))
}

// ProgramVersionCodeNEQ applies the NEQ predicate on the "program_version_code" field.
func ProgramVersionCodeNEQ(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNEQ(FieldProgramVersionCode, v))
}

// ProgramVersionCodeIn applies the In predicate on the "program_version_code" field.
func ProgramVersionCodeIn(vs ...string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldIn(FieldProgramVersionCode, vs...))
}

// ProgramVersionCodeNotIn applies the NotIn predicate on the "program_version_code" field.
func ProgramVersionCodeNotIn(vs ...string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNotIn(FieldProgramVersionCode, vs...))
}

// ProgramVersionCodeGT applies the GT predicate on the "program_version_code" field.
func ProgramVersionCodeGT(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGT(FieldProgramVersionCode, v))
}

// ProgramVersionCodeGTE applies the GTE predicate on the "program_version_code" field.
func ProgramVersionCodeGTE(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGTE(FieldProgramVersionCode, v))
}

// ProgramVersionCodeLT applies the LT predicate on the "program_version_code" field.
func ProgramVersionCodeLT(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLT(FieldProgramVersionCode, v))
}

// ProgramVersionCodeLTE applies the LTE predicate on the "program_version_code" field.
func ProgramVersionCodeLTE(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLTE(FieldProgramVersionCode, v))
}

// ProgramVersionCodeContains applies the Contains predicate on the "program_version_code" field.
func ProgramVersionCodeContains(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldContains(FieldProgramVersionCode, v))
}

// ProgramVersionCodeHasPrefix applies the HasPrefix predicate on the "program_version_code" field.
func ProgramVersionCodeHasPrefix(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldHasPrefix(FieldProgramVersionCode, v))
}

// ProgramVersionCodeHasSuffix applies the HasSuffix predicate on the "program_version_code" field.
func ProgramVersionCodeHasSuffix(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldHasSuffix(FieldProgramVersionCode, v))
}

// ProgramVersionCodeEqualFold applies the EqualFold predicate on the "program_version_code" field.
func ProgramVersionCodeEqualFold(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEqualFold(FieldProgramVersionCode, v))
}

// ProgramVersionCodeContainsFold applies the ContainsFold predicate on the "program_version_code" field.
func ProgramVersionCodeContainsFold(v string) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldContainsFold(FieldProgramVersionCode, v))
}

// DayEQ applies the EQ predicate on the "day" field.
func DayEQ(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldEQ(FieldDay, v))
}

// DayNEQ applies the NEQ predicate on the "day" field.
func DayNEQ(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNEQ(FieldDay, v))
}

// DayIn applies the In predicate on the "day" field.
func DayIn(vs ...uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldIn(FieldDay, vs...))
}

// DayNotIn applies the NotIn predicate on the "day" field.
func DayNotIn(vs ...uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldNotIn(FieldDay, vs...))
}

// DayGT applies the GT predicate on the "day" field.
func DayGT(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGT(FieldDay, v))
}

// DayGTE applies the GTE predicate on the "day" field.
func DayGTE(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldGTE(FieldDay, v))
}

// DayLT applies the LT predicate on the "day" field.
func DayLT(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLT(FieldDay, v))
}

// DayLTE applies the LTE predicate on the "day" field.
func DayLTE(v uint) predicate.DailyRoutine {
	return predicate.DailyRoutine(sql.FieldLTE(FieldDay, v))
}

// HasProgramVersion applies the HasEdge predicate on the "program_version" edge.
func HasProgramVersion() predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProgramVersionTable, ProgramVersionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramVersionWith applies the HasEdge predicate on the "program_version" edge with a given conditions (other predicates).
func HasProgramVersionWith(preds ...predicate.ProgramVersion) predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		step := newProgramVersionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoutineActs applies the HasEdge predicate on the "routine_acts" edge.
func HasRoutineActs() predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoutineActsTable, RoutineActsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoutineActsWith applies the HasEdge predicate on the "routine_acts" edge with a given conditions (other predicates).
func HasRoutineActsWith(preds ...predicate.RoutineAct) predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		step := newRoutineActsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DailyRoutine) predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DailyRoutine) predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
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
func Not(p predicate.DailyRoutine) predicate.DailyRoutine {
	return predicate.DailyRoutine(func(s *sql.Selector) {
		p(s.Not())
	})
}