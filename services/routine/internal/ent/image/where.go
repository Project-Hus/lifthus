// Code generated by ent, DO NOT EDIT.

package image

import (
	"routine/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldID, id))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldKey, v))
}

// Src applies equality check predicate on the "src" field. It's identical to SrcEQ.
func Src(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSrc, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldKey, v))
}

// SrcEQ applies the EQ predicate on the "src" field.
func SrcEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSrc, v))
}

// SrcNEQ applies the NEQ predicate on the "src" field.
func SrcNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldSrc, v))
}

// SrcIn applies the In predicate on the "src" field.
func SrcIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldSrc, vs...))
}

// SrcNotIn applies the NotIn predicate on the "src" field.
func SrcNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldSrc, vs...))
}

// SrcGT applies the GT predicate on the "src" field.
func SrcGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldSrc, v))
}

// SrcGTE applies the GTE predicate on the "src" field.
func SrcGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldSrc, v))
}

// SrcLT applies the LT predicate on the "src" field.
func SrcLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldSrc, v))
}

// SrcLTE applies the LTE predicate on the "src" field.
func SrcLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldSrc, v))
}

// SrcContains applies the Contains predicate on the "src" field.
func SrcContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldSrc, v))
}

// SrcHasPrefix applies the HasPrefix predicate on the "src" field.
func SrcHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldSrc, v))
}

// SrcHasSuffix applies the HasSuffix predicate on the "src" field.
func SrcHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldSrc, v))
}

// SrcEqualFold applies the EqualFold predicate on the "src" field.
func SrcEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldSrc, v))
}

// SrcContainsFold applies the ContainsFold predicate on the "src" field.
func SrcContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldSrc, v))
}

// HasActVersions applies the HasEdge predicate on the "act_versions" edge.
func HasActVersions() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ActVersionsTable, ActVersionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActVersionsWith applies the HasEdge predicate on the "act_versions" edge with a given conditions (other predicates).
func HasActVersionsWith(preds ...predicate.ActVersion) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newActVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProgramVersions applies the HasEdge predicate on the "program_versions" edge.
func HasProgramVersions() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProgramVersionsTable, ProgramVersionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramVersionsWith applies the HasEdge predicate on the "program_versions" edge with a given conditions (other predicates).
func HasProgramVersionsWith(preds ...predicate.ProgramVersion) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newProgramVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasActImages applies the HasEdge predicate on the "act_images" edge.
func HasActImages() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ActImagesTable, ActImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActImagesWith applies the HasEdge predicate on the "act_images" edge with a given conditions (other predicates).
func HasActImagesWith(preds ...predicate.ActImage) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newActImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProgramImages applies the HasEdge predicate on the "program_images" edge.
func HasProgramImages() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ProgramImagesTable, ProgramImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramImagesWith applies the HasEdge predicate on the "program_images" edge with a given conditions (other predicates).
func HasProgramImagesWith(preds ...predicate.ProgramImage) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newProgramImagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
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
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		p(s.Not())
	})
}
