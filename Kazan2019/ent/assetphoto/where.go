// Code generated by ent, DO NOT EDIT.

package assetphoto

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldLTE(FieldID, id))
}

// AssetPhoto applies equality check predicate on the "AssetPhoto" field. It's identical to AssetPhotoEQ.
func AssetPhoto(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldEQ(FieldAssetPhoto, v))
}

// AssetPhotoEQ applies the EQ predicate on the "AssetPhoto" field.
func AssetPhotoEQ(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldEQ(FieldAssetPhoto, v))
}

// AssetPhotoNEQ applies the NEQ predicate on the "AssetPhoto" field.
func AssetPhotoNEQ(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldNEQ(FieldAssetPhoto, v))
}

// AssetPhotoIn applies the In predicate on the "AssetPhoto" field.
func AssetPhotoIn(vs ...string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldIn(FieldAssetPhoto, vs...))
}

// AssetPhotoNotIn applies the NotIn predicate on the "AssetPhoto" field.
func AssetPhotoNotIn(vs ...string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldNotIn(FieldAssetPhoto, vs...))
}

// AssetPhotoGT applies the GT predicate on the "AssetPhoto" field.
func AssetPhotoGT(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldGT(FieldAssetPhoto, v))
}

// AssetPhotoGTE applies the GTE predicate on the "AssetPhoto" field.
func AssetPhotoGTE(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldGTE(FieldAssetPhoto, v))
}

// AssetPhotoLT applies the LT predicate on the "AssetPhoto" field.
func AssetPhotoLT(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldLT(FieldAssetPhoto, v))
}

// AssetPhotoLTE applies the LTE predicate on the "AssetPhoto" field.
func AssetPhotoLTE(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldLTE(FieldAssetPhoto, v))
}

// AssetPhotoContains applies the Contains predicate on the "AssetPhoto" field.
func AssetPhotoContains(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldContains(FieldAssetPhoto, v))
}

// AssetPhotoHasPrefix applies the HasPrefix predicate on the "AssetPhoto" field.
func AssetPhotoHasPrefix(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldHasPrefix(FieldAssetPhoto, v))
}

// AssetPhotoHasSuffix applies the HasSuffix predicate on the "AssetPhoto" field.
func AssetPhotoHasSuffix(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldHasSuffix(FieldAssetPhoto, v))
}

// AssetPhotoEqualFold applies the EqualFold predicate on the "AssetPhoto" field.
func AssetPhotoEqualFold(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldEqualFold(FieldAssetPhoto, v))
}

// AssetPhotoContainsFold applies the ContainsFold predicate on the "AssetPhoto" field.
func AssetPhotoContainsFold(v string) predicate.AssetPhoto {
	return predicate.AssetPhoto(sql.FieldContainsFold(FieldAssetPhoto, v))
}

// HasAsset applies the HasEdge predicate on the "Asset" edge.
func HasAsset() predicate.AssetPhoto {
	return predicate.AssetPhoto(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetWith applies the HasEdge predicate on the "Asset" edge with a given conditions (other predicates).
func HasAssetWith(preds ...predicate.Asset) predicate.AssetPhoto {
	return predicate.AssetPhoto(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssetTable, AssetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AssetPhoto) predicate.AssetPhoto {
	return predicate.AssetPhoto(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AssetPhoto) predicate.AssetPhoto {
	return predicate.AssetPhoto(func(s *sql.Selector) {
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
func Not(p predicate.AssetPhoto) predicate.AssetPhoto {
	return predicate.AssetPhoto(func(s *sql.Selector) {
		p(s.Not())
	})
}
