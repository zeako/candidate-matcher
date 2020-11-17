// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/zeako/candidate-matcher/ent/job"
	"github.com/zeako/candidate-matcher/ent/predicate"
	"github.com/zeako/candidate-matcher/ent/skill"
)

// JobUpdate is the builder for updating Job entities.
type JobUpdate struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where adds a new predicate for the builder.
func (ju *JobUpdate) Where(ps ...predicate.Job) *JobUpdate {
	ju.mutation.predicates = append(ju.mutation.predicates, ps...)
	return ju
}

// SetTitle sets the title field.
func (ju *JobUpdate) SetTitle(s string) *JobUpdate {
	ju.mutation.SetTitle(s)
	return ju
}

// SetSkillID sets the skill edge to Skill by id.
func (ju *JobUpdate) SetSkillID(id int) *JobUpdate {
	ju.mutation.SetSkillID(id)
	return ju
}

// SetSkill sets the skill edge to Skill.
func (ju *JobUpdate) SetSkill(s *Skill) *JobUpdate {
	return ju.SetSkillID(s.ID)
}

// Mutation returns the JobMutation object of the builder.
func (ju *JobUpdate) Mutation() *JobMutation {
	return ju.mutation
}

// ClearSkill clears the "skill" edge to type Skill.
func (ju *JobUpdate) ClearSkill() *JobUpdate {
	ju.mutation.ClearSkill()
	return ju
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ju *JobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ju.hooks) == 0 {
		if err = ju.check(); err != nil {
			return 0, err
		}
		affected, err = ju.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ju.check(); err != nil {
				return 0, err
			}
			ju.mutation = mutation
			affected, err = ju.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ju.hooks) - 1; i >= 0; i-- {
			mut = ju.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ju.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JobUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JobUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JobUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ju *JobUpdate) check() error {
	if _, ok := ju.mutation.SkillID(); ju.mutation.SkillCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"skill\"")
	}
	return nil
}

func (ju *JobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldTitle,
		})
	}
	if ju.mutation.SkillCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.SkillTable,
			Columns: []string{job.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.SkillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.SkillTable,
			Columns: []string{job.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// JobUpdateOne is the builder for updating a single Job entity.
type JobUpdateOne struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// SetTitle sets the title field.
func (juo *JobUpdateOne) SetTitle(s string) *JobUpdateOne {
	juo.mutation.SetTitle(s)
	return juo
}

// SetSkillID sets the skill edge to Skill by id.
func (juo *JobUpdateOne) SetSkillID(id int) *JobUpdateOne {
	juo.mutation.SetSkillID(id)
	return juo
}

// SetSkill sets the skill edge to Skill.
func (juo *JobUpdateOne) SetSkill(s *Skill) *JobUpdateOne {
	return juo.SetSkillID(s.ID)
}

// Mutation returns the JobMutation object of the builder.
func (juo *JobUpdateOne) Mutation() *JobMutation {
	return juo.mutation
}

// ClearSkill clears the "skill" edge to type Skill.
func (juo *JobUpdateOne) ClearSkill() *JobUpdateOne {
	juo.mutation.ClearSkill()
	return juo
}

// Save executes the query and returns the updated entity.
func (juo *JobUpdateOne) Save(ctx context.Context) (*Job, error) {
	var (
		err  error
		node *Job
	)
	if len(juo.hooks) == 0 {
		if err = juo.check(); err != nil {
			return nil, err
		}
		node, err = juo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = juo.check(); err != nil {
				return nil, err
			}
			juo.mutation = mutation
			node, err = juo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(juo.hooks) - 1; i >= 0; i-- {
			mut = juo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, juo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JobUpdateOne) SaveX(ctx context.Context) *Job {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JobUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JobUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (juo *JobUpdateOne) check() error {
	if _, ok := juo.mutation.SkillID(); juo.mutation.SkillCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"skill\"")
	}
	return nil
}

func (juo *JobUpdateOne) sqlSave(ctx context.Context) (_node *Job, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Job.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := juo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: job.FieldTitle,
		})
	}
	if juo.mutation.SkillCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.SkillTable,
			Columns: []string{job.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skill.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.SkillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.SkillTable,
			Columns: []string{job.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: skill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Job{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}