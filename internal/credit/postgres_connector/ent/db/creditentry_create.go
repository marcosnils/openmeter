// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/feature"
)

// CreditEntryCreate is the builder for creating a CreditEntry entity.
type CreditEntryCreate struct {
	config
	mutation *CreditEntryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cec *CreditEntryCreate) SetCreatedAt(t time.Time) *CreditEntryCreate {
	cec.mutation.SetCreatedAt(t)
	return cec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableCreatedAt(t *time.Time) *CreditEntryCreate {
	if t != nil {
		cec.SetCreatedAt(*t)
	}
	return cec
}

// SetUpdatedAt sets the "updated_at" field.
func (cec *CreditEntryCreate) SetUpdatedAt(t time.Time) *CreditEntryCreate {
	cec.mutation.SetUpdatedAt(t)
	return cec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableUpdatedAt(t *time.Time) *CreditEntryCreate {
	if t != nil {
		cec.SetUpdatedAt(*t)
	}
	return cec
}

// SetNamespace sets the "namespace" field.
func (cec *CreditEntryCreate) SetNamespace(s string) *CreditEntryCreate {
	cec.mutation.SetNamespace(s)
	return cec
}

// SetLedgerID sets the "ledger_id" field.
func (cec *CreditEntryCreate) SetLedgerID(s string) *CreditEntryCreate {
	cec.mutation.SetLedgerID(s)
	return cec
}

// SetEntryType sets the "entry_type" field.
func (cec *CreditEntryCreate) SetEntryType(ct credit.EntryType) *CreditEntryCreate {
	cec.mutation.SetEntryType(ct)
	return cec
}

// SetType sets the "type" field.
func (cec *CreditEntryCreate) SetType(ct credit.GrantType) *CreditEntryCreate {
	cec.mutation.SetType(ct)
	return cec
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableType(ct *credit.GrantType) *CreditEntryCreate {
	if ct != nil {
		cec.SetType(*ct)
	}
	return cec
}

// SetFeatureID sets the "feature_id" field.
func (cec *CreditEntryCreate) SetFeatureID(s string) *CreditEntryCreate {
	cec.mutation.SetFeatureID(s)
	return cec
}

// SetNillableFeatureID sets the "feature_id" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableFeatureID(s *string) *CreditEntryCreate {
	if s != nil {
		cec.SetFeatureID(*s)
	}
	return cec
}

// SetAmount sets the "amount" field.
func (cec *CreditEntryCreate) SetAmount(f float64) *CreditEntryCreate {
	cec.mutation.SetAmount(f)
	return cec
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableAmount(f *float64) *CreditEntryCreate {
	if f != nil {
		cec.SetAmount(*f)
	}
	return cec
}

// SetPriority sets the "priority" field.
func (cec *CreditEntryCreate) SetPriority(u uint8) *CreditEntryCreate {
	cec.mutation.SetPriority(u)
	return cec
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillablePriority(u *uint8) *CreditEntryCreate {
	if u != nil {
		cec.SetPriority(*u)
	}
	return cec
}

// SetEffectiveAt sets the "effective_at" field.
func (cec *CreditEntryCreate) SetEffectiveAt(t time.Time) *CreditEntryCreate {
	cec.mutation.SetEffectiveAt(t)
	return cec
}

// SetNillableEffectiveAt sets the "effective_at" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableEffectiveAt(t *time.Time) *CreditEntryCreate {
	if t != nil {
		cec.SetEffectiveAt(*t)
	}
	return cec
}

// SetExpirationPeriodDuration sets the "expiration_period_duration" field.
func (cec *CreditEntryCreate) SetExpirationPeriodDuration(cpd credit.ExpirationPeriodDuration) *CreditEntryCreate {
	cec.mutation.SetExpirationPeriodDuration(cpd)
	return cec
}

// SetNillableExpirationPeriodDuration sets the "expiration_period_duration" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableExpirationPeriodDuration(cpd *credit.ExpirationPeriodDuration) *CreditEntryCreate {
	if cpd != nil {
		cec.SetExpirationPeriodDuration(*cpd)
	}
	return cec
}

// SetExpirationPeriodCount sets the "expiration_period_count" field.
func (cec *CreditEntryCreate) SetExpirationPeriodCount(u uint8) *CreditEntryCreate {
	cec.mutation.SetExpirationPeriodCount(u)
	return cec
}

// SetNillableExpirationPeriodCount sets the "expiration_period_count" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableExpirationPeriodCount(u *uint8) *CreditEntryCreate {
	if u != nil {
		cec.SetExpirationPeriodCount(*u)
	}
	return cec
}

// SetExpirationAt sets the "expiration_at" field.
func (cec *CreditEntryCreate) SetExpirationAt(t time.Time) *CreditEntryCreate {
	cec.mutation.SetExpirationAt(t)
	return cec
}

// SetNillableExpirationAt sets the "expiration_at" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableExpirationAt(t *time.Time) *CreditEntryCreate {
	if t != nil {
		cec.SetExpirationAt(*t)
	}
	return cec
}

// SetRolloverType sets the "rollover_type" field.
func (cec *CreditEntryCreate) SetRolloverType(crt credit.GrantRolloverType) *CreditEntryCreate {
	cec.mutation.SetRolloverType(crt)
	return cec
}

// SetNillableRolloverType sets the "rollover_type" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableRolloverType(crt *credit.GrantRolloverType) *CreditEntryCreate {
	if crt != nil {
		cec.SetRolloverType(*crt)
	}
	return cec
}

// SetRolloverMaxAmount sets the "rollover_max_amount" field.
func (cec *CreditEntryCreate) SetRolloverMaxAmount(f float64) *CreditEntryCreate {
	cec.mutation.SetRolloverMaxAmount(f)
	return cec
}

// SetNillableRolloverMaxAmount sets the "rollover_max_amount" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableRolloverMaxAmount(f *float64) *CreditEntryCreate {
	if f != nil {
		cec.SetRolloverMaxAmount(*f)
	}
	return cec
}

// SetMetadata sets the "metadata" field.
func (cec *CreditEntryCreate) SetMetadata(m map[string]string) *CreditEntryCreate {
	cec.mutation.SetMetadata(m)
	return cec
}

// SetParentID sets the "parent_id" field.
func (cec *CreditEntryCreate) SetParentID(s string) *CreditEntryCreate {
	cec.mutation.SetParentID(s)
	return cec
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableParentID(s *string) *CreditEntryCreate {
	if s != nil {
		cec.SetParentID(*s)
	}
	return cec
}

// SetID sets the "id" field.
func (cec *CreditEntryCreate) SetID(s string) *CreditEntryCreate {
	cec.mutation.SetID(s)
	return cec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableID(s *string) *CreditEntryCreate {
	if s != nil {
		cec.SetID(*s)
	}
	return cec
}

// SetParent sets the "parent" edge to the CreditEntry entity.
func (cec *CreditEntryCreate) SetParent(c *CreditEntry) *CreditEntryCreate {
	return cec.SetParentID(c.ID)
}

// SetChildrenID sets the "children" edge to the CreditEntry entity by ID.
func (cec *CreditEntryCreate) SetChildrenID(id string) *CreditEntryCreate {
	cec.mutation.SetChildrenID(id)
	return cec
}

// SetNillableChildrenID sets the "children" edge to the CreditEntry entity by ID if the given value is not nil.
func (cec *CreditEntryCreate) SetNillableChildrenID(id *string) *CreditEntryCreate {
	if id != nil {
		cec = cec.SetChildrenID(*id)
	}
	return cec
}

// SetChildren sets the "children" edge to the CreditEntry entity.
func (cec *CreditEntryCreate) SetChildren(c *CreditEntry) *CreditEntryCreate {
	return cec.SetChildrenID(c.ID)
}

// SetFeature sets the "feature" edge to the Feature entity.
func (cec *CreditEntryCreate) SetFeature(f *Feature) *CreditEntryCreate {
	return cec.SetFeatureID(f.ID)
}

// Mutation returns the CreditEntryMutation object of the builder.
func (cec *CreditEntryCreate) Mutation() *CreditEntryMutation {
	return cec.mutation
}

// Save creates the CreditEntry in the database.
func (cec *CreditEntryCreate) Save(ctx context.Context) (*CreditEntry, error) {
	cec.defaults()
	return withHooks(ctx, cec.sqlSave, cec.mutation, cec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cec *CreditEntryCreate) SaveX(ctx context.Context) *CreditEntry {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cec *CreditEntryCreate) Exec(ctx context.Context) error {
	_, err := cec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cec *CreditEntryCreate) ExecX(ctx context.Context) {
	if err := cec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cec *CreditEntryCreate) defaults() {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		v := creditentry.DefaultCreatedAt()
		cec.mutation.SetCreatedAt(v)
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		v := creditentry.DefaultUpdatedAt()
		cec.mutation.SetUpdatedAt(v)
	}
	if _, ok := cec.mutation.Priority(); !ok {
		v := creditentry.DefaultPriority
		cec.mutation.SetPriority(v)
	}
	if _, ok := cec.mutation.EffectiveAt(); !ok {
		v := creditentry.DefaultEffectiveAt()
		cec.mutation.SetEffectiveAt(v)
	}
	if _, ok := cec.mutation.ID(); !ok {
		v := creditentry.DefaultID()
		cec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cec *CreditEntryCreate) check() error {
	if _, ok := cec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "CreditEntry.created_at"`)}
	}
	if _, ok := cec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "CreditEntry.updated_at"`)}
	}
	if _, ok := cec.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "CreditEntry.namespace"`)}
	}
	if v, ok := cec.mutation.Namespace(); ok {
		if err := creditentry.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "CreditEntry.namespace": %w`, err)}
		}
	}
	if _, ok := cec.mutation.LedgerID(); !ok {
		return &ValidationError{Name: "ledger_id", err: errors.New(`db: missing required field "CreditEntry.ledger_id"`)}
	}
	if _, ok := cec.mutation.EntryType(); !ok {
		return &ValidationError{Name: "entry_type", err: errors.New(`db: missing required field "CreditEntry.entry_type"`)}
	}
	if v, ok := cec.mutation.EntryType(); ok {
		if err := creditentry.EntryTypeValidator(v); err != nil {
			return &ValidationError{Name: "entry_type", err: fmt.Errorf(`db: validator failed for field "CreditEntry.entry_type": %w`, err)}
		}
	}
	if v, ok := cec.mutation.GetType(); ok {
		if err := creditentry.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "CreditEntry.type": %w`, err)}
		}
	}
	if _, ok := cec.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`db: missing required field "CreditEntry.priority"`)}
	}
	if _, ok := cec.mutation.EffectiveAt(); !ok {
		return &ValidationError{Name: "effective_at", err: errors.New(`db: missing required field "CreditEntry.effective_at"`)}
	}
	if v, ok := cec.mutation.ExpirationPeriodDuration(); ok {
		if err := creditentry.ExpirationPeriodDurationValidator(v); err != nil {
			return &ValidationError{Name: "expiration_period_duration", err: fmt.Errorf(`db: validator failed for field "CreditEntry.expiration_period_duration": %w`, err)}
		}
	}
	if v, ok := cec.mutation.RolloverType(); ok {
		if err := creditentry.RolloverTypeValidator(v); err != nil {
			return &ValidationError{Name: "rollover_type", err: fmt.Errorf(`db: validator failed for field "CreditEntry.rollover_type": %w`, err)}
		}
	}
	return nil
}

func (cec *CreditEntryCreate) sqlSave(ctx context.Context) (*CreditEntry, error) {
	if err := cec.check(); err != nil {
		return nil, err
	}
	_node, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected CreditEntry.ID type: %T", _spec.ID.Value)
		}
	}
	cec.mutation.id = &_node.ID
	cec.mutation.done = true
	return _node, nil
}

func (cec *CreditEntryCreate) createSpec() (*CreditEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &CreditEntry{config: cec.config}
		_spec = sqlgraph.NewCreateSpec(creditentry.Table, sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString))
	)
	_spec.OnConflict = cec.conflict
	if id, ok := cec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cec.mutation.CreatedAt(); ok {
		_spec.SetField(creditentry.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cec.mutation.UpdatedAt(); ok {
		_spec.SetField(creditentry.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cec.mutation.Namespace(); ok {
		_spec.SetField(creditentry.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := cec.mutation.LedgerID(); ok {
		_spec.SetField(creditentry.FieldLedgerID, field.TypeString, value)
		_node.LedgerID = value
	}
	if value, ok := cec.mutation.EntryType(); ok {
		_spec.SetField(creditentry.FieldEntryType, field.TypeEnum, value)
		_node.EntryType = value
	}
	if value, ok := cec.mutation.GetType(); ok {
		_spec.SetField(creditentry.FieldType, field.TypeEnum, value)
		_node.Type = &value
	}
	if value, ok := cec.mutation.Amount(); ok {
		_spec.SetField(creditentry.FieldAmount, field.TypeFloat64, value)
		_node.Amount = &value
	}
	if value, ok := cec.mutation.Priority(); ok {
		_spec.SetField(creditentry.FieldPriority, field.TypeUint8, value)
		_node.Priority = value
	}
	if value, ok := cec.mutation.EffectiveAt(); ok {
		_spec.SetField(creditentry.FieldEffectiveAt, field.TypeTime, value)
		_node.EffectiveAt = value
	}
	if value, ok := cec.mutation.ExpirationPeriodDuration(); ok {
		_spec.SetField(creditentry.FieldExpirationPeriodDuration, field.TypeEnum, value)
		_node.ExpirationPeriodDuration = &value
	}
	if value, ok := cec.mutation.ExpirationPeriodCount(); ok {
		_spec.SetField(creditentry.FieldExpirationPeriodCount, field.TypeUint8, value)
		_node.ExpirationPeriodCount = &value
	}
	if value, ok := cec.mutation.ExpirationAt(); ok {
		_spec.SetField(creditentry.FieldExpirationAt, field.TypeTime, value)
		_node.ExpirationAt = &value
	}
	if value, ok := cec.mutation.RolloverType(); ok {
		_spec.SetField(creditentry.FieldRolloverType, field.TypeEnum, value)
		_node.RolloverType = &value
	}
	if value, ok := cec.mutation.RolloverMaxAmount(); ok {
		_spec.SetField(creditentry.FieldRolloverMaxAmount, field.TypeFloat64, value)
		_node.RolloverMaxAmount = &value
	}
	if value, ok := cec.mutation.Metadata(); ok {
		_spec.SetField(creditentry.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := cec.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   creditentry.ParentTable,
			Columns: []string{creditentry.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   creditentry.ChildrenTable,
			Columns: []string{creditentry.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creditentry.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cec.mutation.FeatureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   creditentry.FeatureTable,
			Columns: []string{creditentry.FeatureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feature.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FeatureID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CreditEntry.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CreditEntryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cec *CreditEntryCreate) OnConflict(opts ...sql.ConflictOption) *CreditEntryUpsertOne {
	cec.conflict = opts
	return &CreditEntryUpsertOne{
		create: cec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cec *CreditEntryCreate) OnConflictColumns(columns ...string) *CreditEntryUpsertOne {
	cec.conflict = append(cec.conflict, sql.ConflictColumns(columns...))
	return &CreditEntryUpsertOne{
		create: cec,
	}
}

type (
	// CreditEntryUpsertOne is the builder for "upsert"-ing
	//  one CreditEntry node.
	CreditEntryUpsertOne struct {
		create *CreditEntryCreate
	}

	// CreditEntryUpsert is the "OnConflict" setter.
	CreditEntryUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *CreditEntryUpsert) SetUpdatedAt(v time.Time) *CreditEntryUpsert {
	u.Set(creditentry.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CreditEntryUpsert) UpdateUpdatedAt() *CreditEntryUpsert {
	u.SetExcluded(creditentry.FieldUpdatedAt)
	return u
}

// SetMetadata sets the "metadata" field.
func (u *CreditEntryUpsert) SetMetadata(v map[string]string) *CreditEntryUpsert {
	u.Set(creditentry.FieldMetadata, v)
	return u
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *CreditEntryUpsert) UpdateMetadata() *CreditEntryUpsert {
	u.SetExcluded(creditentry.FieldMetadata)
	return u
}

// ClearMetadata clears the value of the "metadata" field.
func (u *CreditEntryUpsert) ClearMetadata() *CreditEntryUpsert {
	u.SetNull(creditentry.FieldMetadata)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(creditentry.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CreditEntryUpsertOne) UpdateNewValues() *CreditEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(creditentry.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(creditentry.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(creditentry.FieldNamespace)
		}
		if _, exists := u.create.mutation.LedgerID(); exists {
			s.SetIgnore(creditentry.FieldLedgerID)
		}
		if _, exists := u.create.mutation.EntryType(); exists {
			s.SetIgnore(creditentry.FieldEntryType)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(creditentry.FieldType)
		}
		if _, exists := u.create.mutation.FeatureID(); exists {
			s.SetIgnore(creditentry.FieldFeatureID)
		}
		if _, exists := u.create.mutation.Amount(); exists {
			s.SetIgnore(creditentry.FieldAmount)
		}
		if _, exists := u.create.mutation.Priority(); exists {
			s.SetIgnore(creditentry.FieldPriority)
		}
		if _, exists := u.create.mutation.EffectiveAt(); exists {
			s.SetIgnore(creditentry.FieldEffectiveAt)
		}
		if _, exists := u.create.mutation.ExpirationPeriodDuration(); exists {
			s.SetIgnore(creditentry.FieldExpirationPeriodDuration)
		}
		if _, exists := u.create.mutation.ExpirationPeriodCount(); exists {
			s.SetIgnore(creditentry.FieldExpirationPeriodCount)
		}
		if _, exists := u.create.mutation.ExpirationAt(); exists {
			s.SetIgnore(creditentry.FieldExpirationAt)
		}
		if _, exists := u.create.mutation.RolloverType(); exists {
			s.SetIgnore(creditentry.FieldRolloverType)
		}
		if _, exists := u.create.mutation.RolloverMaxAmount(); exists {
			s.SetIgnore(creditentry.FieldRolloverMaxAmount)
		}
		if _, exists := u.create.mutation.ParentID(); exists {
			s.SetIgnore(creditentry.FieldParentID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CreditEntryUpsertOne) Ignore() *CreditEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CreditEntryUpsertOne) DoNothing() *CreditEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CreditEntryCreate.OnConflict
// documentation for more info.
func (u *CreditEntryUpsertOne) Update(set func(*CreditEntryUpsert)) *CreditEntryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CreditEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CreditEntryUpsertOne) SetUpdatedAt(v time.Time) *CreditEntryUpsertOne {
	return u.Update(func(s *CreditEntryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CreditEntryUpsertOne) UpdateUpdatedAt() *CreditEntryUpsertOne {
	return u.Update(func(s *CreditEntryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetMetadata sets the "metadata" field.
func (u *CreditEntryUpsertOne) SetMetadata(v map[string]string) *CreditEntryUpsertOne {
	return u.Update(func(s *CreditEntryUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *CreditEntryUpsertOne) UpdateMetadata() *CreditEntryUpsertOne {
	return u.Update(func(s *CreditEntryUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *CreditEntryUpsertOne) ClearMetadata() *CreditEntryUpsertOne {
	return u.Update(func(s *CreditEntryUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *CreditEntryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for CreditEntryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CreditEntryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CreditEntryUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: CreditEntryUpsertOne.ID is not supported by MySQL driver. Use CreditEntryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CreditEntryUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CreditEntryCreateBulk is the builder for creating many CreditEntry entities in bulk.
type CreditEntryCreateBulk struct {
	config
	err      error
	builders []*CreditEntryCreate
	conflict []sql.ConflictOption
}

// Save creates the CreditEntry entities in the database.
func (cecb *CreditEntryCreateBulk) Save(ctx context.Context) ([]*CreditEntry, error) {
	if cecb.err != nil {
		return nil, cecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cecb.builders))
	nodes := make([]*CreditEntry, len(cecb.builders))
	mutators := make([]Mutator, len(cecb.builders))
	for i := range cecb.builders {
		func(i int, root context.Context) {
			builder := cecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CreditEntryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cecb *CreditEntryCreateBulk) SaveX(ctx context.Context) []*CreditEntry {
	v, err := cecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cecb *CreditEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := cecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cecb *CreditEntryCreateBulk) ExecX(ctx context.Context) {
	if err := cecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CreditEntry.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CreditEntryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cecb *CreditEntryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CreditEntryUpsertBulk {
	cecb.conflict = opts
	return &CreditEntryUpsertBulk{
		create: cecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cecb *CreditEntryCreateBulk) OnConflictColumns(columns ...string) *CreditEntryUpsertBulk {
	cecb.conflict = append(cecb.conflict, sql.ConflictColumns(columns...))
	return &CreditEntryUpsertBulk{
		create: cecb,
	}
}

// CreditEntryUpsertBulk is the builder for "upsert"-ing
// a bulk of CreditEntry nodes.
type CreditEntryUpsertBulk struct {
	create *CreditEntryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(creditentry.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CreditEntryUpsertBulk) UpdateNewValues() *CreditEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(creditentry.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(creditentry.FieldCreatedAt)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(creditentry.FieldNamespace)
			}
			if _, exists := b.mutation.LedgerID(); exists {
				s.SetIgnore(creditentry.FieldLedgerID)
			}
			if _, exists := b.mutation.EntryType(); exists {
				s.SetIgnore(creditentry.FieldEntryType)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(creditentry.FieldType)
			}
			if _, exists := b.mutation.FeatureID(); exists {
				s.SetIgnore(creditentry.FieldFeatureID)
			}
			if _, exists := b.mutation.Amount(); exists {
				s.SetIgnore(creditentry.FieldAmount)
			}
			if _, exists := b.mutation.Priority(); exists {
				s.SetIgnore(creditentry.FieldPriority)
			}
			if _, exists := b.mutation.EffectiveAt(); exists {
				s.SetIgnore(creditentry.FieldEffectiveAt)
			}
			if _, exists := b.mutation.ExpirationPeriodDuration(); exists {
				s.SetIgnore(creditentry.FieldExpirationPeriodDuration)
			}
			if _, exists := b.mutation.ExpirationPeriodCount(); exists {
				s.SetIgnore(creditentry.FieldExpirationPeriodCount)
			}
			if _, exists := b.mutation.ExpirationAt(); exists {
				s.SetIgnore(creditentry.FieldExpirationAt)
			}
			if _, exists := b.mutation.RolloverType(); exists {
				s.SetIgnore(creditentry.FieldRolloverType)
			}
			if _, exists := b.mutation.RolloverMaxAmount(); exists {
				s.SetIgnore(creditentry.FieldRolloverMaxAmount)
			}
			if _, exists := b.mutation.ParentID(); exists {
				s.SetIgnore(creditentry.FieldParentID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CreditEntry.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CreditEntryUpsertBulk) Ignore() *CreditEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CreditEntryUpsertBulk) DoNothing() *CreditEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CreditEntryCreateBulk.OnConflict
// documentation for more info.
func (u *CreditEntryUpsertBulk) Update(set func(*CreditEntryUpsert)) *CreditEntryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CreditEntryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CreditEntryUpsertBulk) SetUpdatedAt(v time.Time) *CreditEntryUpsertBulk {
	return u.Update(func(s *CreditEntryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CreditEntryUpsertBulk) UpdateUpdatedAt() *CreditEntryUpsertBulk {
	return u.Update(func(s *CreditEntryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetMetadata sets the "metadata" field.
func (u *CreditEntryUpsertBulk) SetMetadata(v map[string]string) *CreditEntryUpsertBulk {
	return u.Update(func(s *CreditEntryUpsert) {
		s.SetMetadata(v)
	})
}

// UpdateMetadata sets the "metadata" field to the value that was provided on create.
func (u *CreditEntryUpsertBulk) UpdateMetadata() *CreditEntryUpsertBulk {
	return u.Update(func(s *CreditEntryUpsert) {
		s.UpdateMetadata()
	})
}

// ClearMetadata clears the value of the "metadata" field.
func (u *CreditEntryUpsertBulk) ClearMetadata() *CreditEntryUpsertBulk {
	return u.Update(func(s *CreditEntryUpsert) {
		s.ClearMetadata()
	})
}

// Exec executes the query.
func (u *CreditEntryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the CreditEntryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for CreditEntryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CreditEntryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
