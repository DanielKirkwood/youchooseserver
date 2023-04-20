package schema

import (
	"errors"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Time mixin adds created_at and updated_at fields
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Validate(func(s string) error {
			if !isValidEmail(s) {
				return errors.New("email must be valid")
			}
			return nil
		}).Unique().NotEmpty(),
		field.String("otp").
			Sensitive().
			Optional().
			Nillable().
			MinLen(36).
			MaxLen(36),
		field.Time("otp_expires_at").
			Optional().
			Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", User.Type).Through("friendships", Friendship.Type),
	}
}
