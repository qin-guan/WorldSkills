package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DepartmentLocation holds the schema definition for the DepartmentLocation entity.
type DepartmentLocation struct {
	ent.Schema
}

// Fields of the DepartmentLocation.
func (DepartmentLocation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.Int("DepartmentID"),
		field.Int("LocationID"),
		field.Time("EndDate"),
		field.Time("StartDate"),
	}
}

// Edges of the DepartmentLocation.
func (DepartmentLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Asset", Asset.Type),
		edge.To("FromDepartmentLocation", AssetTransferLog.Type),
		edge.To("ToDepartmentLocation", AssetTransferLog.Type),
		edge.From("Department", Department.Type).
			Ref("DepartmentLocation").
			Field("DepartmentID").
			Required().
			Unique(),
		edge.From("Location", Location.Type).
			Ref("DepartmentLocation").
			Field("LocationID").
			Required().
			Unique(),
	}
}
