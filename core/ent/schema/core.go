package schema

import "entgo.io/ent"

// Core holds the schema definition for the Core entity.
type Core struct {
	ent.Schema
}

// Fields of the Core.
func (Core) Fields() []ent.Field {
	return nil
}

// Edges of the Core.
func (Core) Edges() []ent.Edge {
	return nil
}
