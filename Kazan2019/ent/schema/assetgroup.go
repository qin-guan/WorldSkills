package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AssetGroup holds the schema definition for the AssetGroup entity.
type AssetGroup struct {
	ent.Schema
}

// Fields of the AssetGroup.
func (AssetGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.String("Name"),
	}
}

// Edges of the AssetGroup.
func (AssetGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Asset", Asset.Type),
	}
}
