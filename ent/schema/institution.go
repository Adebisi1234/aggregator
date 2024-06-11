package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Institution holds the schema definition for the Institution entity.
type Institution struct {
	ent.Schema
}

// Fields of the Institution.
func (Institution) Fields() []ent.Field {

	return []ent.Field{
		field.String("code").Unique(),
		field.String("name"),
		field.Enum("type").
			Values("bank", "mobile_money").
			Default("bank"), // Sets updated_at to the current time on update
	}
}

// Mixin of the FiatCurrency.
func (Institution) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fiat_currency", FiatCurrency.Type).
			Ref("institutions").
			Unique(),
	}
}
