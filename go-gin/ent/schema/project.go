package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),

		field.String("name").
			NotEmpty(),

		field.String("description").
			Optional().
			Nillable(),

		field.Time("start_date"),

		field.Time("end_date").
			Optional().
			Nillable(),

		field.Enum("status").
			Values("created", "in_progress", "completed", "suspended").
			Default("created"),
	}
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assignments", Assignment.Type).
			Comment("A project can have many assignments."),
	}
}
