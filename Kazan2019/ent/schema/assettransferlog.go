package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AssetTransferLog holds the schema definition for the AssetTransferLog entity.
type AssetTransferLog struct {
	ent.Schema
}

// Fields of the AssetTransferLog.
func (AssetTransferLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.Int("AssetID"),
		field.Int("FromAssetSN").Positive(),
		field.Int("ToAssetSN").Positive(),
		field.Int("FromDepartmentLocationID"),
		field.Int("ToDepartmentLocationID"),
		field.Time("TransferDate"),
	}
}

// Edges of the AssetTransferLog.
func (AssetTransferLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Asset", Asset.Type).
			Ref("AssetTransferLog").
			Field("AssetID").
			Required().
			Unique(),
		edge.From("FromDepartmentLocation", DepartmentLocation.Type).
			Ref("FromDepartmentLocation").
			Field("FromDepartmentLocationID").
			Required().
			Unique(),
		edge.From("ToDepartmentLocation", DepartmentLocation.Type).
			Ref("ToDepartmentLocation").
			Field("ToDepartmentLocationID").
			Required().
			Unique(),
	}
}
