package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID"),
		field.String("Name"),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("DepartmentLocation", DepartmentLocation.Type),
	}
}
