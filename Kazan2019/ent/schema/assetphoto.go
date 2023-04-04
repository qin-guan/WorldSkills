package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AssetPhoto holds the schema definition for the AssetPhoto entity.
type AssetPhoto struct {
	ent.Schema
}

// Fields of the AssetPhoto.
func (AssetPhoto) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.String("AssetPhoto"),
	}
}

// Edges of the AssetPhoto.
func (AssetPhoto) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Asset", Asset.Type).
			Ref("AssetPhoto").
			Required().
			Unique(),
	}
}
