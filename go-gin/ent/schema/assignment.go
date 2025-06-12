package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Assignment holds the schema definition for the Assignment entity.
type Assignment struct {
	ent.Schema
}

// Fields of the Assignment.
func (Assignment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),

		field.String("title").
			NotEmpty(),

		field.String("description").
			Optional().
			Nillable(),

		field.Time("start_date"),

		field.Time("end_date").
			Optional().
			Nillable(),

		field.Enum("status").
			Values("active", "in_active", "completed"),
	}
}

func (Assignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}

// Edges of the Assignment.
func (Assignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("assignments").
			Unique().
			Required().
			Immutable().
			Comment("An assignment only has one user and can't be changed after creation"),

		edge.From("project", Project.Type).
			Ref("assignments").
			Unique().
			Required().
			Immutable().
			Comment("An assignment only has one project and can't be changed after creation"),

		edge.To("statements", Statement.Type).
			Comment("An assignment can have many statements."),
	}
}
