package schema

import (
	"hash/maphash"
	"math/rand"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProviderProfile holds the schema definition for the ProviderProfile entity.
type ProviderProfile struct {
	ent.Schema
}

// Mixin of the ProviderProfile.
func (ProviderProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ProviderProfile.
func (ProviderProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(generateProviderID).
			Unique(),
		field.String("trading_name").MaxLen(80),
		field.String("country").MaxLen(80),
	}
}

// Edges of the ProviderProfile.
func (ProviderProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("api_key", APIKey.Type).
			Ref("provider_profile").
			Unique().
			Required(),
		edge.To("order_tokens", ProviderOrderToken.Type),
		edge.To("availability", ProviderAvailability.Type).
			Unique(),
	}
}

// generateProviderID generates a random string of the specified length
func generateProviderID() string {
	// Define the character set for the random string
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Create a random string of 8 characters
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))

	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}
