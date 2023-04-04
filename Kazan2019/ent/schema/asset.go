package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.Int("AssetGroupID"),
		field.Int("EmployeeID"),
		field.Int("DepartmentLocationID"),
		field.Int("AssetSN").
			Positive(),
		field.String("AssetName"),
		field.String("Description"),
		field.Time("WarrantyDate"),
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("AssetPhoto", AssetPhoto.Type),
		edge.To("AssetTransferLog", AssetTransferLog.Type),
		edge.From("DepartmentLocation", DepartmentLocation.Type).
			Ref("Asset").
			Field("DepartmentLocationID").
			Required().
			Unique(),
		edge.From("Employee", Employee.Type).
			Ref("Asset").
			Field("EmployeeID").
			Required().
			Unique(),
		edge.From("AssetGroup", AssetGroup.Type).
			Ref("Asset").
			Field("AssetGroupID").
			Required().
			Unique(),
	}
}
