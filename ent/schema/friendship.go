package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Friendship holds the edge schema definition of the Friendship relationship.
type Friendship struct {
	ent.Schema
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Immutable(),
		field.Int("friend_id").
			Immutable(),
		field.Enum("status").
			Values("requested", "accepted"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("friend", User.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
