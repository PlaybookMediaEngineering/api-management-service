package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Workspace represents a row from 'unkey.workspaces'.
type Workspace struct {
	ID                   string         `json:"id"`                     // id
	Name                 string         `json:"name"`                   // name
	Slug                 string         `json:"slug"`                   // slug
	TenantID             string         `json:"tenant_id"`              // tenant_id
	Internal             bool           `json:"internal"`               // internal
	StripeCustomerID     sql.NullString `json:"stripe_customer_id"`     // stripe_customer_id
	StripeSubscriptionID sql.NullString `json:"stripe_subscription_id"` // stripe_subscription_id
	Plan                 NullPlan       `json:"plan"`                   // plan
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Workspace] exists in the database.
func (w *Workspace) Exists() bool {
	return w._exists
}

// Deleted returns true when the [Workspace] has been marked for deletion
// from the database.
func (w *Workspace) Deleted() bool {
	return w._deleted
}

// Insert inserts the [Workspace] to the database.
func (w *Workspace) Insert(ctx context.Context, db DB) error {
	switch {
	case w._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case w._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO unkey.workspaces (` +
		`id, name, slug, tenant_id, internal, stripe_customer_id, stripe_subscription_id, plan` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, w.ID, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan)
	if _, err := db.ExecContext(ctx, sqlstr, w.ID, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan); err != nil {
		return logerror(err)
	}
	// set exists
	w._exists = true
	return nil
}

// Update updates a [Workspace] in the database.
func (w *Workspace) Update(ctx context.Context, db DB) error {
	switch {
	case !w._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case w._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE unkey.workspaces SET ` +
		`name = ?, slug = ?, tenant_id = ?, internal = ?, stripe_customer_id = ?, stripe_subscription_id = ?, plan = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan, w.ID)
	if _, err := db.ExecContext(ctx, sqlstr, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan, w.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Workspace] to the database.
func (w *Workspace) Save(ctx context.Context, db DB) error {
	if w.Exists() {
		return w.Update(ctx, db)
	}
	return w.Insert(ctx, db)
}

// Upsert performs an upsert for [Workspace].
func (w *Workspace) Upsert(ctx context.Context, db DB) error {
	switch {
	case w._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO unkey.workspaces (` +
		`id, name, slug, tenant_id, internal, stripe_customer_id, stripe_subscription_id, plan` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id = VALUES(id), name = VALUES(name), slug = VALUES(slug), tenant_id = VALUES(tenant_id), internal = VALUES(internal), stripe_customer_id = VALUES(stripe_customer_id), stripe_subscription_id = VALUES(stripe_subscription_id), plan = VALUES(plan)`
	// run
	logf(sqlstr, w.ID, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan)
	if _, err := db.ExecContext(ctx, sqlstr, w.ID, w.Name, w.Slug, w.TenantID, w.Internal, w.StripeCustomerID, w.StripeSubscriptionID, w.Plan); err != nil {
		return logerror(err)
	}
	// set exists
	w._exists = true
	return nil
}

// Delete deletes the [Workspace] from the database.
func (w *Workspace) Delete(ctx context.Context, db DB) error {
	switch {
	case !w._exists: // doesn't exist
		return nil
	case w._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM unkey.workspaces ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, w.ID)
	if _, err := db.ExecContext(ctx, sqlstr, w.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	w._deleted = true
	return nil
}

// WorkspaceBySlug retrieves a row from 'unkey.workspaces' as a [Workspace].
//
// Generated from index 'slug_idx'.
func WorkspaceBySlug(ctx context.Context, db DB, slug string) (*Workspace, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, slug, tenant_id, internal, stripe_customer_id, stripe_subscription_id, plan ` +
		`FROM unkey.workspaces ` +
		`WHERE slug = ?`
	// run
	logf(sqlstr, slug)
	w := Workspace{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, slug).Scan(&w.ID, &w.Name, &w.Slug, &w.TenantID, &w.Internal, &w.StripeCustomerID, &w.StripeSubscriptionID, &w.Plan); err != nil {
		return nil, logerror(err)
	}
	return &w, nil
}

// WorkspaceByTenantID retrieves a row from 'unkey.workspaces' as a [Workspace].
//
// Generated from index 'tenant_id_idx'.
func WorkspaceByTenantID(ctx context.Context, db DB, tenantID string) (*Workspace, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, slug, tenant_id, internal, stripe_customer_id, stripe_subscription_id, plan ` +
		`FROM unkey.workspaces ` +
		`WHERE tenant_id = ?`
	// run
	logf(sqlstr, tenantID)
	w := Workspace{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, tenantID).Scan(&w.ID, &w.Name, &w.Slug, &w.TenantID, &w.Internal, &w.StripeCustomerID, &w.StripeSubscriptionID, &w.Plan); err != nil {
		return nil, logerror(err)
	}
	return &w, nil
}

// WorkspaceByID retrieves a row from 'unkey.workspaces' as a [Workspace].
//
// Generated from index 'workspaces_id_pkey'.
func WorkspaceByID(ctx context.Context, db DB, id string) (*Workspace, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, slug, tenant_id, internal, stripe_customer_id, stripe_subscription_id, plan ` +
		`FROM unkey.workspaces ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	w := Workspace{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&w.ID, &w.Name, &w.Slug, &w.TenantID, &w.Internal, &w.StripeCustomerID, &w.StripeSubscriptionID, &w.Plan); err != nil {
		return nil, logerror(err)
	}
	return &w, nil
}