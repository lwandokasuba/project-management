package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Statement holds the schema definition for the Statement entity.
type Statement struct {
	ent.Schema
}

// Fields of the Statement.
func (Statement) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),

		field.Enum("type").
			Values("expense", "profit").
			Immutable(),

		field.Float("amount"),

		field.String("description").
			Optional().
			Nillable(),
	}
}

func (Statement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Edges of the Statement.
func (Statement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assignment", Assignment.Type).
			Ref("statements").
			Unique().
			Required().
			Immutable().
			Comment("A statement only has one assignment and can't be changed after creation"),
	}
}
