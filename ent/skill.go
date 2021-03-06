// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/zeako/candidate-matcher/ent/skill"
)

// Skill is the model entity for the Skill schema.
type Skill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkillQuery when eager-loading is set.
	Edges SkillEdges `json:"edges"`
}

// SkillEdges holds the relations/edges for other nodes in the graph.
type SkillEdges struct {
	// Candidate holds the value of the candidate edge.
	Candidate []*Candidate
	// Job holds the value of the job edge.
	Job []*Job
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CandidateOrErr returns the Candidate value or an error if the edge
// was not loaded in eager-loading.
func (e SkillEdges) CandidateOrErr() ([]*Candidate, error) {
	if e.loadedTypes[0] {
		return e.Candidate, nil
	}
	return nil, &NotLoadedError{edge: "candidate"}
}

// JobOrErr returns the Job value or an error if the edge
// was not loaded in eager-loading.
func (e SkillEdges) JobOrErr() ([]*Job, error) {
	if e.loadedTypes[1] {
		return e.Job, nil
	}
	return nil, &NotLoadedError{edge: "job"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Skill) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Skill fields.
func (s *Skill) assignValues(values ...interface{}) error {
	if m, n := len(values), len(skill.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		s.Name = value.String
	}
	return nil
}

// QueryCandidate queries the candidate edge of the Skill.
func (s *Skill) QueryCandidate() *CandidateQuery {
	return (&SkillClient{config: s.config}).QueryCandidate(s)
}

// QueryJob queries the job edge of the Skill.
func (s *Skill) QueryJob() *JobQuery {
	return (&SkillClient{config: s.config}).QueryJob(s)
}

// Update returns a builder for updating this Skill.
// Note that, you need to call Skill.Unwrap() before calling this method, if this Skill
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Skill) Update() *SkillUpdateOne {
	return (&SkillClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Skill) Unwrap() *Skill {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Skill is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Skill) String() string {
	var builder strings.Builder
	builder.WriteString("Skill(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Skills is a parsable slice of Skill.
type Skills []*Skill

func (s Skills) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
