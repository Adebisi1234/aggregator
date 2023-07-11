// Code generated by ent, DO NOT EDIT.

package provideravailability

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/paycrest/paycrest-protocol/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLTE(FieldID, id))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldEndTime, v))
}

// CadenceEQ applies the EQ predicate on the "cadence" field.
func CadenceEQ(v Cadence) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldCadence, v))
}

// CadenceNEQ applies the NEQ predicate on the "cadence" field.
func CadenceNEQ(v Cadence) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNEQ(FieldCadence, v))
}

// CadenceIn applies the In predicate on the "cadence" field.
func CadenceIn(vs ...Cadence) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldIn(FieldCadence, vs...))
}

// CadenceNotIn applies the NotIn predicate on the "cadence" field.
func CadenceNotIn(vs ...Cadence) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNotIn(FieldCadence, vs...))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(sql.FieldLTE(FieldEndTime, v))
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.ProviderAvailability {
	return predicate.ProviderAvailability(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.ProviderProfile) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(func(s *sql.Selector) {
		step := newProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProviderAvailability) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProviderAvailability) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(func(s *sql.Selector) {
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
func Not(p predicate.ProviderAvailability) predicate.ProviderAvailability {
	return predicate.ProviderAvailability(func(s *sql.Selector) {
		p(s.Not())
	})
}
