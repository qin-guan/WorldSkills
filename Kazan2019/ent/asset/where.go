// Code generated by ent, DO NOT EDIT.

package asset

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/qin-guan/WorldSkills/Kazan2019/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldID, id))
}

// AssetGroupID applies equality check predicate on the "AssetGroupID" field. It's identical to AssetGroupIDEQ.
func AssetGroupID(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetGroupID, v))
}

// EmployeeID applies equality check predicate on the "EmployeeID" field. It's identical to EmployeeIDEQ.
func EmployeeID(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldEmployeeID, v))
}

// DepartmentLocationID applies equality check predicate on the "DepartmentLocationID" field. It's identical to DepartmentLocationIDEQ.
func DepartmentLocationID(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldDepartmentLocationID, v))
}

// AssetSN applies equality check predicate on the "AssetSN" field. It's identical to AssetSNEQ.
func AssetSN(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetSN, v))
}

// AssetName applies equality check predicate on the "AssetName" field. It's identical to AssetNameEQ.
func AssetName(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetName, v))
}

// Description applies equality check predicate on the "Description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldDescription, v))
}

// WarrantyDate applies equality check predicate on the "WarrantyDate" field. It's identical to WarrantyDateEQ.
func WarrantyDate(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldWarrantyDate, v))
}

// AssetGroupIDEQ applies the EQ predicate on the "AssetGroupID" field.
func AssetGroupIDEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetGroupID, v))
}

// AssetGroupIDNEQ applies the NEQ predicate on the "AssetGroupID" field.
func AssetGroupIDNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldAssetGroupID, v))
}

// AssetGroupIDIn applies the In predicate on the "AssetGroupID" field.
func AssetGroupIDIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldAssetGroupID, vs...))
}

// AssetGroupIDNotIn applies the NotIn predicate on the "AssetGroupID" field.
func AssetGroupIDNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldAssetGroupID, vs...))
}

// EmployeeIDEQ applies the EQ predicate on the "EmployeeID" field.
func EmployeeIDEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "EmployeeID" field.
func EmployeeIDNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "EmployeeID" field.
func EmployeeIDIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "EmployeeID" field.
func EmployeeIDNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// DepartmentLocationIDEQ applies the EQ predicate on the "DepartmentLocationID" field.
func DepartmentLocationIDEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldDepartmentLocationID, v))
}

// DepartmentLocationIDNEQ applies the NEQ predicate on the "DepartmentLocationID" field.
func DepartmentLocationIDNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldDepartmentLocationID, v))
}

// DepartmentLocationIDIn applies the In predicate on the "DepartmentLocationID" field.
func DepartmentLocationIDIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldDepartmentLocationID, vs...))
}

// DepartmentLocationIDNotIn applies the NotIn predicate on the "DepartmentLocationID" field.
func DepartmentLocationIDNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldDepartmentLocationID, vs...))
}

// AssetSNEQ applies the EQ predicate on the "AssetSN" field.
func AssetSNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetSN, v))
}

// AssetSNNEQ applies the NEQ predicate on the "AssetSN" field.
func AssetSNNEQ(v int) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldAssetSN, v))
}

// AssetSNIn applies the In predicate on the "AssetSN" field.
func AssetSNIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldAssetSN, vs...))
}

// AssetSNNotIn applies the NotIn predicate on the "AssetSN" field.
func AssetSNNotIn(vs ...int) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldAssetSN, vs...))
}

// AssetSNGT applies the GT predicate on the "AssetSN" field.
func AssetSNGT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldAssetSN, v))
}

// AssetSNGTE applies the GTE predicate on the "AssetSN" field.
func AssetSNGTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldAssetSN, v))
}

// AssetSNLT applies the LT predicate on the "AssetSN" field.
func AssetSNLT(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldAssetSN, v))
}

// AssetSNLTE applies the LTE predicate on the "AssetSN" field.
func AssetSNLTE(v int) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldAssetSN, v))
}

// AssetNameEQ applies the EQ predicate on the "AssetName" field.
func AssetNameEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldAssetName, v))
}

// AssetNameNEQ applies the NEQ predicate on the "AssetName" field.
func AssetNameNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldAssetName, v))
}

// AssetNameIn applies the In predicate on the "AssetName" field.
func AssetNameIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldAssetName, vs...))
}

// AssetNameNotIn applies the NotIn predicate on the "AssetName" field.
func AssetNameNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldAssetName, vs...))
}

// AssetNameGT applies the GT predicate on the "AssetName" field.
func AssetNameGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldAssetName, v))
}

// AssetNameGTE applies the GTE predicate on the "AssetName" field.
func AssetNameGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldAssetName, v))
}

// AssetNameLT applies the LT predicate on the "AssetName" field.
func AssetNameLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldAssetName, v))
}

// AssetNameLTE applies the LTE predicate on the "AssetName" field.
func AssetNameLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldAssetName, v))
}

// AssetNameContains applies the Contains predicate on the "AssetName" field.
func AssetNameContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldAssetName, v))
}

// AssetNameHasPrefix applies the HasPrefix predicate on the "AssetName" field.
func AssetNameHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldAssetName, v))
}

// AssetNameHasSuffix applies the HasSuffix predicate on the "AssetName" field.
func AssetNameHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldAssetName, v))
}

// AssetNameEqualFold applies the EqualFold predicate on the "AssetName" field.
func AssetNameEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldAssetName, v))
}

// AssetNameContainsFold applies the ContainsFold predicate on the "AssetName" field.
func AssetNameContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldAssetName, v))
}

// DescriptionEQ applies the EQ predicate on the "Description" field.
func DescriptionEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "Description" field.
func DescriptionNEQ(v string) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "Description" field.
func DescriptionIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "Description" field.
func DescriptionNotIn(vs ...string) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "Description" field.
func DescriptionGT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "Description" field.
func DescriptionGTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "Description" field.
func DescriptionLT(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "Description" field.
func DescriptionLTE(v string) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "Description" field.
func DescriptionContains(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "Description" field.
func DescriptionHasPrefix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "Description" field.
func DescriptionHasSuffix(v string) predicate.Asset {
	return predicate.Asset(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "Description" field.
func DescriptionEqualFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "Description" field.
func DescriptionContainsFold(v string) predicate.Asset {
	return predicate.Asset(sql.FieldContainsFold(FieldDescription, v))
}

// WarrantyDateEQ applies the EQ predicate on the "WarrantyDate" field.
func WarrantyDateEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldEQ(FieldWarrantyDate, v))
}

// WarrantyDateNEQ applies the NEQ predicate on the "WarrantyDate" field.
func WarrantyDateNEQ(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNEQ(FieldWarrantyDate, v))
}

// WarrantyDateIn applies the In predicate on the "WarrantyDate" field.
func WarrantyDateIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldIn(FieldWarrantyDate, vs...))
}

// WarrantyDateNotIn applies the NotIn predicate on the "WarrantyDate" field.
func WarrantyDateNotIn(vs ...time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldNotIn(FieldWarrantyDate, vs...))
}

// WarrantyDateGT applies the GT predicate on the "WarrantyDate" field.
func WarrantyDateGT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGT(FieldWarrantyDate, v))
}

// WarrantyDateGTE applies the GTE predicate on the "WarrantyDate" field.
func WarrantyDateGTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldGTE(FieldWarrantyDate, v))
}

// WarrantyDateLT applies the LT predicate on the "WarrantyDate" field.
func WarrantyDateLT(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLT(FieldWarrantyDate, v))
}

// WarrantyDateLTE applies the LTE predicate on the "WarrantyDate" field.
func WarrantyDateLTE(v time.Time) predicate.Asset {
	return predicate.Asset(sql.FieldLTE(FieldWarrantyDate, v))
}

// HasAssetPhoto applies the HasEdge predicate on the "AssetPhoto" edge.
func HasAssetPhoto() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetPhotoTable, AssetPhotoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetPhotoWith applies the HasEdge predicate on the "AssetPhoto" edge with a given conditions (other predicates).
func HasAssetPhotoWith(preds ...predicate.AssetPhoto) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetPhotoInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetPhotoTable, AssetPhotoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssetTransferLog applies the HasEdge predicate on the "AssetTransferLog" edge.
func HasAssetTransferLog() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetTransferLogTable, AssetTransferLogColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetTransferLogWith applies the HasEdge predicate on the "AssetTransferLog" edge with a given conditions (other predicates).
func HasAssetTransferLogWith(preds ...predicate.AssetTransferLog) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetTransferLogInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AssetTransferLogTable, AssetTransferLogColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartmentLocation applies the HasEdge predicate on the "DepartmentLocation" edge.
func HasDepartmentLocation() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentLocationTable, DepartmentLocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentLocationWith applies the HasEdge predicate on the "DepartmentLocation" edge with a given conditions (other predicates).
func HasDepartmentLocationWith(preds ...predicate.DepartmentLocation) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentLocationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentLocationTable, DepartmentLocationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmployee applies the HasEdge predicate on the "Employee" edge.
func HasEmployee() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "Employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmployeeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAssetGroup applies the HasEdge predicate on the "AssetGroup" edge.
func HasAssetGroup() predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssetGroupTable, AssetGroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAssetGroupWith applies the HasEdge predicate on the "AssetGroup" edge with a given conditions (other predicates).
func HasAssetGroupWith(preds ...predicate.AssetGroup) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AssetGroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AssetGroupTable, AssetGroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
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
func Not(p predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		p(s.Not())
	})
}
