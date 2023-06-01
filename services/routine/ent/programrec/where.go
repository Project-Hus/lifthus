// Code generated by ent, DO NOT EDIT.

package programrec

import (
	"routine/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldID, id))
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldAuthor, v))
}

// ProgramID applies equality check predicate on the "program_id" field. It's identical to ProgramIDEQ.
func ProgramID(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldProgramID, v))
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldStartDate, v))
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldEndDate, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldComment, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldUpdatedAt, v))
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldAuthor, v))
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldAuthor, v))
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldAuthor, vs...))
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldAuthor, vs...))
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldAuthor, v))
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldAuthor, v))
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldAuthor, v))
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldAuthor, v))
}

// ProgramIDEQ applies the EQ predicate on the "program_id" field.
func ProgramIDEQ(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldProgramID, v))
}

// ProgramIDNEQ applies the NEQ predicate on the "program_id" field.
func ProgramIDNEQ(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldProgramID, v))
}

// ProgramIDIn applies the In predicate on the "program_id" field.
func ProgramIDIn(vs ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldProgramID, vs...))
}

// ProgramIDNotIn applies the NotIn predicate on the "program_id" field.
func ProgramIDNotIn(vs ...uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldProgramID, vs...))
}

// ProgramIDGT applies the GT predicate on the "program_id" field.
func ProgramIDGT(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldProgramID, v))
}

// ProgramIDGTE applies the GTE predicate on the "program_id" field.
func ProgramIDGTE(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldProgramID, v))
}

// ProgramIDLT applies the LT predicate on the "program_id" field.
func ProgramIDLT(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldProgramID, v))
}

// ProgramIDLTE applies the LTE predicate on the "program_id" field.
func ProgramIDLTE(v uint64) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldProgramID, v))
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldStartDate, v))
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldStartDate, v))
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldStartDate, vs...))
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldStartDate, vs...))
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldStartDate, v))
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldStartDate, v))
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldStartDate, v))
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldStartDate, v))
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldEndDate, v))
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldEndDate, v))
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldEndDate, vs...))
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldEndDate, vs...))
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldEndDate, v))
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldEndDate, v))
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldEndDate, v))
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldEndDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldStatus, vs...))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldContainsFold(FieldComment, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProgramRec {
	return predicate.ProgramRec(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasProgram applies the HasEdge predicate on the "program" edge.
func HasProgram() predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProgramTable, ProgramColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramWith applies the HasEdge predicate on the "program" edge with a given conditions (other predicates).
func HasProgramWith(preds ...predicate.Program) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := newProgramStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWeeklyRoutineRecs applies the HasEdge predicate on the "weekly_routine_recs" edge.
func HasWeeklyRoutineRecs() predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WeeklyRoutineRecsTable, WeeklyRoutineRecsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeeklyRoutineRecsWith applies the HasEdge predicate on the "weekly_routine_recs" edge with a given conditions (other predicates).
func HasWeeklyRoutineRecsWith(preds ...predicate.WeeklyRoutineRec) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := newWeeklyRoutineRecsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDailyRoutineRecs applies the HasEdge predicate on the "daily_routine_recs" edge.
func HasDailyRoutineRecs() predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DailyRoutineRecsTable, DailyRoutineRecsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDailyRoutineRecsWith applies the HasEdge predicate on the "daily_routine_recs" edge with a given conditions (other predicates).
func HasDailyRoutineRecsWith(preds ...predicate.DailyRoutineRec) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		step := newDailyRoutineRecsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProgramRec) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProgramRec) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
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
func Not(p predicate.ProgramRec) predicate.ProgramRec {
	return predicate.ProgramRec(func(s *sql.Selector) {
		p(s.Not())
	})
}
