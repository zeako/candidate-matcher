package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Candidate holds the schema definition for the Candidate entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate.
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().NotEmpty(),
	}
}

// Edges of the Candidate.
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("skills", Skill.Type),
	}
}
