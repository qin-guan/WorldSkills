// Code generated by ent, DO NOT EDIT.

package assetgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.AssetGroup {
	return predicate.AssetGroup(sql.FieldContainsFold(FieldName, v))
}

// HasAsset applies the HasEdge predicate on the "Asset" edge.
func HasAsset() predicate.AssetGroup {
	return predicate.AssetGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetWith applies the HasEdge predicate on the "Asset" edge with a given conditions (other predicates).
func HasAssetWith(preds ...predicate.Asset) predicate.AssetGroup {
	return predicate.AssetGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AssetGroup) predicate.AssetGroup {
	return predicate.AssetGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AssetGroup) predicate.AssetGroup {
	return predicate.AssetGroup(func(s *sql.Selector) {
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
func Not(p predicate.AssetGroup) predicate.AssetGroup {
	return predicate.AssetGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
