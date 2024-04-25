// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: providers.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createProvider = `-- name: CreateProvider :one
INSERT INTO providers (
    name,
    project_id,
    class,
    implements,
    definition,
    auth_flows
    ) VALUES ($1, $2, $3, $4, $5::jsonb, $6) RETURNING id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class
`

type CreateProviderParams struct {
	Name       string              `json:"name"`
	ProjectID  uuid.UUID           `json:"project_id"`
	Class      ProviderClass       `json:"class"`
	Implements []ProviderType      `json:"implements"`
	Definition json.RawMessage     `json:"definition"`
	AuthFlows  []AuthorizationFlow `json:"auth_flows"`
}

func (q *Queries) CreateProvider(ctx context.Context, arg CreateProviderParams) (Provider, error) {
	row := q.db.QueryRowContext(ctx, createProvider,
		arg.Name,
		arg.ProjectID,
		arg.Class,
		pq.Array(arg.Implements),
		arg.Definition,
		pq.Array(arg.AuthFlows),
	)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.ProjectID,
		pq.Array(&i.Implements),
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.AuthFlows),
		&i.Class,
	)
	return i, err
}

const deleteProvider = `-- name: DeleteProvider :exec
DELETE FROM providers
   WHERE id = $1 AND project_id = $2
`

type DeleteProviderParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) DeleteProvider(ctx context.Context, arg DeleteProviderParams) error {
	_, err := q.db.ExecContext(ctx, deleteProvider, arg.ID, arg.ProjectID)
	return err
}

const findProviders = `-- name: FindProviders :many

SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers
WHERE project_id = ANY($1::uuid[])
    AND ($2::provider_type = ANY(implements) OR $2::provider_type IS NULL)
    AND (lower(name) = lower($3::text) OR $3::text IS NULL)
`

type FindProvidersParams struct {
	Projects []uuid.UUID      `json:"projects"`
	Trait    NullProviderType `json:"trait"`
	Name     sql.NullString   `json:"name"`
}

// FindProviders allows us to take a trait and filter
// providers by it. It also optionally takes a name, in case we want to
// filter by name as well.
func (q *Queries) FindProviders(ctx context.Context, arg FindProvidersParams) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, findProviders, pq.Array(arg.Projects), arg.Trait, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.ProjectID,
			pq.Array(&i.Implements),
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.AuthFlows),
			&i.Class,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProviderByID = `-- name: GetProviderByID :one
SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers WHERE id = $1
`

func (q *Queries) GetProviderByID(ctx context.Context, id uuid.UUID) (Provider, error) {
	row := q.db.QueryRowContext(ctx, getProviderByID, id)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.ProjectID,
		pq.Array(&i.Implements),
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.AuthFlows),
		&i.Class,
	)
	return i, err
}

const getProviderByIDAndProject = `-- name: GetProviderByIDAndProject :one
SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers WHERE id = $1 AND project_id = $2
`

type GetProviderByIDAndProjectParams struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (q *Queries) GetProviderByIDAndProject(ctx context.Context, arg GetProviderByIDAndProjectParams) (Provider, error) {
	row := q.db.QueryRowContext(ctx, getProviderByIDAndProject, arg.ID, arg.ProjectID)
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.ProjectID,
		pq.Array(&i.Implements),
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.AuthFlows),
		&i.Class,
	)
	return i, err
}

const getProviderByName = `-- name: GetProviderByName :one

SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers WHERE lower(name) = lower($1) AND project_id = ANY($2::uuid[])
LIMIT 1
`

type GetProviderByNameParams struct {
	Name     string      `json:"name"`
	Projects []uuid.UUID `json:"projects"`
}

// GetProviderByName allows us to get a provider by its name. This takes
// into account the project hierarchy, so it will only return the provider
// if it exists in the project or any of its ancestors. It'll return the first
// provider that matches the name.
func (q *Queries) GetProviderByName(ctx context.Context, arg GetProviderByNameParams) (Provider, error) {
	row := q.db.QueryRowContext(ctx, getProviderByName, arg.Name, pq.Array(arg.Projects))
	var i Provider
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Version,
		&i.ProjectID,
		pq.Array(&i.Implements),
		&i.Definition,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.AuthFlows),
		&i.Class,
	)
	return i, err
}

const globalListProviders = `-- name: GlobalListProviders :many
SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers
`

func (q *Queries) GlobalListProviders(ctx context.Context) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, globalListProviders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.ProjectID,
			pq.Array(&i.Implements),
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.AuthFlows),
			&i.Class,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const globalListProvidersByClass = `-- name: GlobalListProvidersByClass :many
SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers WHERE class = $1
`

func (q *Queries) GlobalListProvidersByClass(ctx context.Context, class ProviderClass) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, globalListProvidersByClass, class)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.ProjectID,
			pq.Array(&i.Implements),
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.AuthFlows),
			&i.Class,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProvidersByProjectID = `-- name: ListProvidersByProjectID :many

SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers WHERE project_id = ANY($1::uuid[])
`

// ListProvidersByProjectID allows us to list all providers
// for a given array of projects.
func (q *Queries) ListProvidersByProjectID(ctx context.Context, projects []uuid.UUID) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, listProvidersByProjectID, pq.Array(projects))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.ProjectID,
			pq.Array(&i.Implements),
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.AuthFlows),
			&i.Class,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProvidersByProjectIDPaginated = `-- name: ListProvidersByProjectIDPaginated :many

SELECT id, name, version, project_id, implements, definition, created_at, updated_at, auth_flows, class FROM providers
WHERE project_id = $1
    AND (created_at > $2 OR $2 IS NULL)
ORDER BY created_at ASC
LIMIT $3
`

type ListProvidersByProjectIDPaginatedParams struct {
	ProjectID uuid.UUID    `json:"project_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	Limit     int32        `json:"limit"`
}

// ListProvidersByProjectIDPaginated allows us to lits all providers for a given project
// with pagination taken into account. In this case, the cursor is the creation date.
func (q *Queries) ListProvidersByProjectIDPaginated(ctx context.Context, arg ListProvidersByProjectIDPaginatedParams) ([]Provider, error) {
	rows, err := q.db.QueryContext(ctx, listProvidersByProjectIDPaginated, arg.ProjectID, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provider{}
	for rows.Next() {
		var i Provider
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Version,
			&i.ProjectID,
			pq.Array(&i.Implements),
			&i.Definition,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.AuthFlows),
			&i.Class,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProvider = `-- name: UpdateProvider :exec
UPDATE providers
    SET implements = $1, definition = $2::jsonb, auth_flows = $3
    WHERE id = $4 AND project_id = $5
`

type UpdateProviderParams struct {
	Implements []ProviderType      `json:"implements"`
	Definition json.RawMessage     `json:"definition"`
	AuthFlows  []AuthorizationFlow `json:"auth_flows"`
	ID         uuid.UUID           `json:"id"`
	ProjectID  uuid.UUID           `json:"project_id"`
}

func (q *Queries) UpdateProvider(ctx context.Context, arg UpdateProviderParams) error {
	_, err := q.db.ExecContext(ctx, updateProvider,
		pq.Array(arg.Implements),
		arg.Definition,
		pq.Array(arg.AuthFlows),
		arg.ID,
		arg.ProjectID,
	)
	return err
}
