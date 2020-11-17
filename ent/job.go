// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/zeako/candidate-matcher/ent/job"
	"github.com/zeako/candidate-matcher/ent/skill"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobQuery when eager-loading is set.
	Edges     JobEdges `json:"edges"`
	job_skill *int
}

// JobEdges holds the relations/edges for other nodes in the graph.
type JobEdges struct {
	// Skill holds the value of the skill edge.
	Skill *Skill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SkillOrErr returns the Skill value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JobEdges) SkillOrErr() (*Skill, error) {
	if e.loadedTypes[0] {
		if e.Skill == nil {
			// The edge skill was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: skill.Label}
		}
		return e.Skill, nil
	}
	return nil, &NotLoadedError{edge: "skill"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Job) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // job_skill
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(values ...interface{}) error {
	if m, n := len(values), len(job.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	j.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		j.Title = value.String
	}
	values = values[1:]
	if len(values) == len(job.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field job_skill", value)
		} else if value.Valid {
			j.job_skill = new(int)
			*j.job_skill = int(value.Int64)
		}
	}
	return nil
}

// QuerySkill queries the skill edge of the Job.
func (j *Job) QuerySkill() *SkillQuery {
	return (&JobClient{config: j.config}).QuerySkill(j)
}

// Update returns a builder for updating this Job.
// Note that, you need to call Job.Unwrap() before calling this method, if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return (&JobClient{config: j.config}).UpdateOne(j)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v", j.ID))
	builder.WriteString(", title=")
	builder.WriteString(j.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job

func (j Jobs) config(cfg config) {
	for _i := range j {
		j[_i].config = cfg
	}
}
